// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenVaultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTokenVaultName(v string) *GetTokenVaultRequest
	GetTokenVaultName() *string
}

type GetTokenVaultRequest struct {
	// example:
	//
	// default
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s GetTokenVaultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenVaultRequest) GoString() string {
	return s.String()
}

func (s *GetTokenVaultRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *GetTokenVaultRequest) SetTokenVaultName(v string) *GetTokenVaultRequest {
	s.TokenVaultName = &v
	return s
}

func (s *GetTokenVaultRequest) Validate() error {
	return dara.Validate(s)
}
