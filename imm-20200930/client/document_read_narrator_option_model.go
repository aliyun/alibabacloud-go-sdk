// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentReadNarratorOption interface {
	dara.Model
	String() string
	GoString() string
	SetNarrate(v bool) *DocumentReadNarratorOption
	GetNarrate() *bool
}

type DocumentReadNarratorOption struct {
	// Specifies whether to fetch.
	//
	// example:
	//
	// true
	Narrate *bool `json:"Narrate,omitempty" xml:"Narrate,omitempty"`
}

func (s DocumentReadNarratorOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentReadNarratorOption) GoString() string {
	return s.String()
}

func (s *DocumentReadNarratorOption) GetNarrate() *bool {
	return s.Narrate
}

func (s *DocumentReadNarratorOption) SetNarrate(v bool) *DocumentReadNarratorOption {
	s.Narrate = &v
	return s
}

func (s *DocumentReadNarratorOption) Validate() error {
	return dara.Validate(s)
}
