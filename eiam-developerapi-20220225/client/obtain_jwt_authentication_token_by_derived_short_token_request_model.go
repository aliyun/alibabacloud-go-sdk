// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainJwtAuthenticationTokenByDerivedShortTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDerivedShortToken(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenRequest
	GetDerivedShortToken() *string
}

type ObtainJwtAuthenticationTokenByDerivedShortTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sk-Nx2vzxxxxxxxxxxxxxxxxx
	DerivedShortToken *string `json:"derivedShortToken,omitempty" xml:"derivedShortToken,omitempty"`
}

func (s ObtainJwtAuthenticationTokenByDerivedShortTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ObtainJwtAuthenticationTokenByDerivedShortTokenRequest) GoString() string {
	return s.String()
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenRequest) GetDerivedShortToken() *string {
	return s.DerivedShortToken
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenRequest) SetDerivedShortToken(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenRequest {
	s.DerivedShortToken = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenRequest) Validate() error {
	return dara.Validate(s)
}
