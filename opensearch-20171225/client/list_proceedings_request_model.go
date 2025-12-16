// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProceedingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterFinished(v bool) *ListProceedingsRequest
	GetFilterFinished() *bool
}

type ListProceedingsRequest struct {
	// Specifies whether the filtering is complete.
	//
	// example:
	//
	// true
	FilterFinished *bool `json:"filterFinished,omitempty" xml:"filterFinished,omitempty"`
}

func (s ListProceedingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProceedingsRequest) GoString() string {
	return s.String()
}

func (s *ListProceedingsRequest) GetFilterFinished() *bool {
	return s.FilterFinished
}

func (s *ListProceedingsRequest) SetFilterFinished(v bool) *ListProceedingsRequest {
	s.FilterFinished = &v
	return s
}

func (s *ListProceedingsRequest) Validate() error {
	return dara.Validate(s)
}
