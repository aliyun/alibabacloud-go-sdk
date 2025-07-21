// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFbIssueLabelsByLCRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaller(v string) *ListFbIssueLabelsByLCRequest
	GetCaller() *string
	SetLanguageType(v string) *ListFbIssueLabelsByLCRequest
	GetLanguageType() *string
}

type ListFbIssueLabelsByLCRequest struct {
	Caller       *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	LanguageType *string `json:"LanguageType,omitempty" xml:"LanguageType,omitempty"`
}

func (s ListFbIssueLabelsByLCRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFbIssueLabelsByLCRequest) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsByLCRequest) GetCaller() *string {
	return s.Caller
}

func (s *ListFbIssueLabelsByLCRequest) GetLanguageType() *string {
	return s.LanguageType
}

func (s *ListFbIssueLabelsByLCRequest) SetCaller(v string) *ListFbIssueLabelsByLCRequest {
	s.Caller = &v
	return s
}

func (s *ListFbIssueLabelsByLCRequest) SetLanguageType(v string) *ListFbIssueLabelsByLCRequest {
	s.LanguageType = &v
	return s
}

func (s *ListFbIssueLabelsByLCRequest) Validate() error {
	return dara.Validate(s)
}
