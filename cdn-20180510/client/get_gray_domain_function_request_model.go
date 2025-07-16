// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGrayDomainFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *GetGrayDomainFunctionRequest
	GetDomainName() *string
	SetFunctionNames(v string) *GetGrayDomainFunctionRequest
	GetFunctionNames() *string
}

type GetGrayDomainFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// domain_status,https_option
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
}

func (s GetGrayDomainFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGrayDomainFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetGrayDomainFunctionRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *GetGrayDomainFunctionRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *GetGrayDomainFunctionRequest) SetDomainName(v string) *GetGrayDomainFunctionRequest {
	s.DomainName = &v
	return s
}

func (s *GetGrayDomainFunctionRequest) SetFunctionNames(v string) *GetGrayDomainFunctionRequest {
	s.FunctionNames = &v
	return s
}

func (s *GetGrayDomainFunctionRequest) Validate() error {
	return dara.Validate(s)
}
