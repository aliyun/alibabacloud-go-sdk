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
	// The timestamp when the diagnostic report was generated.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken   *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DiagnoseItems []*string `json:"diagnoseItems,omitempty" xml:"diagnoseItems,omitempty" type:"Repeated"`
	Indices       []*string `json:"indices,omitempty" xml:"indices,omitempty" type:"Repeated"`
	// example:
	//
	// ALL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The returned data.
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
