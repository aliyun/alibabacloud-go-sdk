// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentParseNarratorOption interface {
	dara.Model
	String() string
	GoString() string
	SetNarrate(v bool) *DocumentParseNarratorOption
	GetNarrate() *bool
}

type DocumentParseNarratorOption struct {
	Narrate *bool `json:"Narrate,omitempty" xml:"Narrate,omitempty"`
}

func (s DocumentParseNarratorOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentParseNarratorOption) GoString() string {
	return s.String()
}

func (s *DocumentParseNarratorOption) GetNarrate() *bool {
	return s.Narrate
}

func (s *DocumentParseNarratorOption) SetNarrate(v bool) *DocumentParseNarratorOption {
	s.Narrate = &v
	return s
}

func (s *DocumentParseNarratorOption) Validate() error {
	return dara.Validate(s)
}
