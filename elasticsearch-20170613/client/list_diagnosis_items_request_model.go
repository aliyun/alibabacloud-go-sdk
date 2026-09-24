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
	SetLevel(v string) *ListDiagnosisItemsRequest
	GetLevel() *string
}

type ListDiagnosisItemsRequest struct {
	// The instance ID.
	//
	// example:
	//
	// es-cn-v0h14zdee000mimee
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The supported request language. Default value: the browser request language. Valid values:
	//
	// - en: English
	//
	// - zh: Simplified Chinese
	//
	// - zt: Traditional Chinese
	//
	// - es: Spanish
	//
	// - fr: French
	//
	// example:
	//
	// en
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// The diagnostic item level. Valid values:
	//
	// - BASIC: basic inspection item (free).
	//
	// - ADVANCED: advanced inspection item (consumes billable tokens).
	//
	// If this parameter is not specified, diagnostic items of all levels are returned.
	//
	// example:
	//
	// BASIC
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
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

func (s *ListDiagnosisItemsRequest) GetLevel() *string {
	return s.Level
}

func (s *ListDiagnosisItemsRequest) SetInstanceId(v string) *ListDiagnosisItemsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDiagnosisItemsRequest) SetLang(v string) *ListDiagnosisItemsRequest {
	s.Lang = &v
	return s
}

func (s *ListDiagnosisItemsRequest) SetLevel(v string) *ListDiagnosisItemsRequest {
	s.Level = &v
	return s
}

func (s *ListDiagnosisItemsRequest) Validate() error {
	return dara.Validate(s)
}
