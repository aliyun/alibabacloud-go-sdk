// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCrossAccountVerifyTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *CrossAccountVerifyTokenRequest
	GetToken() *string
}

type CrossAccountVerifyTokenRequest struct {
	// example:
	//
	// C19D103FEA2D50A584410267CE9FBA56
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CrossAccountVerifyTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CrossAccountVerifyTokenRequest) GoString() string {
	return s.String()
}

func (s *CrossAccountVerifyTokenRequest) GetToken() *string {
	return s.Token
}

func (s *CrossAccountVerifyTokenRequest) SetToken(v string) *CrossAccountVerifyTokenRequest {
	s.Token = &v
	return s
}

func (s *CrossAccountVerifyTokenRequest) Validate() error {
	return dara.Validate(s)
}
