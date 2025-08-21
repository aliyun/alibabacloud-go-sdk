// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosisItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListDiagnosisItemsRequest
	GetLang() *string
}

type ListDiagnosisItemsRequest struct {
	// example:
	//
	// en
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s ListDiagnosisItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisItemsRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosisItemsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDiagnosisItemsRequest) SetLang(v string) *ListDiagnosisItemsRequest {
	s.Lang = &v
	return s
}

func (s *ListDiagnosisItemsRequest) Validate() error {
	return dara.Validate(s)
}
