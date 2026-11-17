// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKyuubiServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *ListKyuubiServicesRequest
	GetToken() *string
}

type ListKyuubiServicesRequest struct {
	// The token of the Kyuubi Gateway.
	//
	// example:
	//
	// 6w3s2e7y7t9fxnvtai9sv1uebw8b7bvc
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ListKyuubiServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiServicesRequest) GoString() string {
	return s.String()
}

func (s *ListKyuubiServicesRequest) GetToken() *string {
	return s.Token
}

func (s *ListKyuubiServicesRequest) SetToken(v string) *ListKyuubiServicesRequest {
	s.Token = &v
	return s
}

func (s *ListKyuubiServicesRequest) Validate() error {
	return dara.Validate(s)
}
