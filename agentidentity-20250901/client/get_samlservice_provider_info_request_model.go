// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSAMLServiceProviderInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPoolName(v string) *GetSAMLServiceProviderInfoRequest
	GetUserPoolName() *string
}

type GetSAMLServiceProviderInfoRequest struct {
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetSAMLServiceProviderInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLServiceProviderInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSAMLServiceProviderInfoRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetSAMLServiceProviderInfoRequest) SetUserPoolName(v string) *GetSAMLServiceProviderInfoRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetSAMLServiceProviderInfoRequest) Validate() error {
	return dara.Validate(s)
}
