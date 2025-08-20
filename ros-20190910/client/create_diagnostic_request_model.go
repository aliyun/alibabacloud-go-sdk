// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosticKey(v string) *CreateDiagnosticRequest
	GetDiagnosticKey() *string
	SetDiagnosticType(v string) *CreateDiagnosticRequest
	GetDiagnosticType() *string
	SetLang(v string) *CreateDiagnosticRequest
	GetLang() *string
	SetProduct(v string) *CreateDiagnosticRequest
	GetProduct() *string
}

type CreateDiagnosticRequest struct {
	// The keyword in the diagnosis.
	//
	// You can specify the ID of the stack that you want to diagnose.
	//
	// This parameter is required.
	//
	// example:
	//
	// 37A5679B-8488-5A5D-8D5C-90E66A577A5D
	DiagnosticKey *string `json:"DiagnosticKey,omitempty" xml:"DiagnosticKey,omitempty"`
	// The type of the item that is diagnosed. Set the value to Stack, which specifies that the stack is diagnosed.
	//
	// example:
	//
	// Stack
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// The language of the diagnostic report to be generated. Only Chinese and English are supported.
	//
	// Valid values:
	//
	// 	- zh-cn
	//
	// 	- en
	//
	// example:
	//
	// zh-cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the product that is diagonosed.
	//
	// example:
	//
	// ros
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s CreateDiagnosticRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticRequest) GetDiagnosticKey() *string {
	return s.DiagnosticKey
}

func (s *CreateDiagnosticRequest) GetDiagnosticType() *string {
	return s.DiagnosticType
}

func (s *CreateDiagnosticRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDiagnosticRequest) GetProduct() *string {
	return s.Product
}

func (s *CreateDiagnosticRequest) SetDiagnosticKey(v string) *CreateDiagnosticRequest {
	s.DiagnosticKey = &v
	return s
}

func (s *CreateDiagnosticRequest) SetDiagnosticType(v string) *CreateDiagnosticRequest {
	s.DiagnosticType = &v
	return s
}

func (s *CreateDiagnosticRequest) SetLang(v string) *CreateDiagnosticRequest {
	s.Lang = &v
	return s
}

func (s *CreateDiagnosticRequest) SetProduct(v string) *CreateDiagnosticRequest {
	s.Product = &v
	return s
}

func (s *CreateDiagnosticRequest) Validate() error {
	return dara.Validate(s)
}
