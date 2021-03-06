//revive:disable:var-naming
package dynamic_classifier

//revive:enable:var-naming

import (
	"github.com/seznam/slo-exporter/pkg/event"
	"io"
)

type matcherType string

type matcher interface {
	getType() matcherType
	set(key string, classification *event.SloClassification) error
	get(key string) (*event.SloClassification, error)
	dumpCSV(w io.Writer) error
}
