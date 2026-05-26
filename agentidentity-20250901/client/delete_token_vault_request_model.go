// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTokenVaultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTokenVaultName(v string) *DeleteTokenVaultRequest
	GetTokenVaultName() *string
}

type DeleteTokenVaultRequest struct {
	// example:
	//
	// default
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s DeleteTokenVaultRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTokenVaultRequest) GoString() string {
	return s.String()
}

func (s *DeleteTokenVaultRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *DeleteTokenVaultRequest) SetTokenVaultName(v string) *DeleteTokenVaultRequest {
	s.TokenVaultName = &v
	return s
}

func (s *DeleteTokenVaultRequest) Validate() error {
	return dara.Validate(s)
}
