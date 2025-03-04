package formatters

import (
	"fmt"
	"strings"

	"github.com/owenrumney/squealer/internal/pkg/match"
)

type DefaultFormatter struct {
}

func (d DefaultFormatter) PrintTransgressions(transgressions []match.Transgression, redacted bool) (string, error) {
	builder := strings.Builder{}

	for _, t := range transgressions {
		var content = t.LineContent
		if redacted {
			content = t.RedactedContent
		}
		builder.Write([]byte(fmt.Sprintf(`
Match Description: │ %s
Content:           │ %s
Filename:          │ %s
Line No:           │ %d
Secret Hash:       │ %s
Commit:            │ %s
Committer:         │ %s (%s)
Committed:         │ %s
Exclude rule:      │ %s
`, t.MatchDescription, content, t.Filename, t.LineNo, t.Hash, t.CommitHash, t.Committer, t.CommitterEmail, t.Committed, t.ExcludeRule)))
	}
	return builder.String(), nil
}
