// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArEditCommonMaterialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *GetArEditCommonMaterialRequest
	GetJwtToken() *string
}

type GetArEditCommonMaterialRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s GetArEditCommonMaterialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArEditCommonMaterialRequest) GoString() string {
	return s.String()
}

func (s *GetArEditCommonMaterialRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *GetArEditCommonMaterialRequest) SetJwtToken(v string) *GetArEditCommonMaterialRequest {
	s.JwtToken = &v
	return s
}

func (s *GetArEditCommonMaterialRequest) Validate() error {
	return dara.Validate(s)
}
