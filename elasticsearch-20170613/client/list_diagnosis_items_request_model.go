// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosisItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDiagnosisItemsRequest
	GetInstanceId() *string
	SetLang(v string) *ListDiagnosisItemsRequest
	GetLang() *string
}

type ListDiagnosisItemsRequest struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The language of the request. Default value: the browser language. Valid values:
	//
	// - en: English
	//
	// - zh: Simplified Chinese
	//
	// - zt: Traditional Chinese
	//
	// - es: Spanish
	//
	// - fr: French.
	//
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

func (s *ListDiagnosisItemsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDiagnosisItemsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDiagnosisItemsRequest) SetInstanceId(v string) *ListDiagnosisItemsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDiagnosisItemsRequest) SetLang(v string) *ListDiagnosisItemsRequest {
	s.Lang = &v
	return s
}

func (s *ListDiagnosisItemsRequest) Validate() error {
	return dara.Validate(s)
}
