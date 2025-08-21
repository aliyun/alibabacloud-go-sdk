// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnoseIndicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListDiagnoseIndicesRequest
	GetLang() *string
}

type ListDiagnoseIndicesRequest struct {
	// The language. Multiple languages are supported.
	//
	// example:
	//
	// en
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s ListDiagnoseIndicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseIndicesRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnoseIndicesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDiagnoseIndicesRequest) SetLang(v string) *ListDiagnoseIndicesRequest {
	s.Lang = &v
	return s
}

func (s *ListDiagnoseIndicesRequest) Validate() error {
	return dara.Validate(s)
}
