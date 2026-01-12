// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplyStatusForTryOnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *GetApplyStatusForTryOnRequest
	GetJwtToken() *string
}

type GetApplyStatusForTryOnRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s GetApplyStatusForTryOnRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplyStatusForTryOnRequest) GoString() string {
	return s.String()
}

func (s *GetApplyStatusForTryOnRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *GetApplyStatusForTryOnRequest) SetJwtToken(v string) *GetApplyStatusForTryOnRequest {
	s.JwtToken = &v
	return s
}

func (s *GetApplyStatusForTryOnRequest) Validate() error {
	return dara.Validate(s)
}
