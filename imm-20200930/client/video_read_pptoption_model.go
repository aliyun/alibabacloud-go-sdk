// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoReadPPTOption interface {
	dara.Model
	String() string
	GoString() string
	SetExtract(v bool) *VideoReadPPTOption
	GetExtract() *bool
}

type VideoReadPPTOption struct {
	// Specifies whether to fetch.
	//
	// example:
	//
	// true
	Extract *bool `json:"Extract,omitempty" xml:"Extract,omitempty"`
}

func (s VideoReadPPTOption) String() string {
	return dara.Prettify(s)
}

func (s VideoReadPPTOption) GoString() string {
	return s.String()
}

func (s *VideoReadPPTOption) GetExtract() *bool {
	return s.Extract
}

func (s *VideoReadPPTOption) SetExtract(v bool) *VideoReadPPTOption {
	s.Extract = &v
	return s
}

func (s *VideoReadPPTOption) Validate() error {
	return dara.Validate(s)
}
