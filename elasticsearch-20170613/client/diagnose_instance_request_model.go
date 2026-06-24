// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DiagnoseInstanceRequest
	GetClientToken() *string
	SetDiagnoseItems(v []*string) *DiagnoseInstanceRequest
	GetDiagnoseItems() []*string
	SetIndices(v []*string) *DiagnoseInstanceRequest
	GetIndices() []*string
	SetType(v string) *DiagnoseInstanceRequest
	GetType() *string
	SetLang(v string) *DiagnoseInstanceRequest
	GetLang() *string
}

type DiagnoseInstanceRequest struct {
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The diagnostic items.
	DiagnoseItems []*string `json:"diagnoseItems,omitempty" xml:"diagnoseItems,omitempty" type:"Repeated"`
	// The list of indexes to diagnose.
	Indices []*string `json:"indices,omitempty" xml:"indices,omitempty" type:"Repeated"`
	// The type of the diagnostic task. Valid values:
	//
	// - ALL: Diagnoses all indexes.
	//
	// - SELECT: Diagnoses selected indexes.
	//
	// example:
	//
	// ALL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The language of the report. Default value: browser language. Valid values:
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

func (s DiagnoseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseInstanceRequest) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DiagnoseInstanceRequest) GetDiagnoseItems() []*string {
	return s.DiagnoseItems
}

func (s *DiagnoseInstanceRequest) GetIndices() []*string {
	return s.Indices
}

func (s *DiagnoseInstanceRequest) GetType() *string {
	return s.Type
}

func (s *DiagnoseInstanceRequest) GetLang() *string {
	return s.Lang
}

func (s *DiagnoseInstanceRequest) SetClientToken(v string) *DiagnoseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DiagnoseInstanceRequest) SetDiagnoseItems(v []*string) *DiagnoseInstanceRequest {
	s.DiagnoseItems = v
	return s
}

func (s *DiagnoseInstanceRequest) SetIndices(v []*string) *DiagnoseInstanceRequest {
	s.Indices = v
	return s
}

func (s *DiagnoseInstanceRequest) SetType(v string) *DiagnoseInstanceRequest {
	s.Type = &v
	return s
}

func (s *DiagnoseInstanceRequest) SetLang(v string) *DiagnoseInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DiagnoseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
