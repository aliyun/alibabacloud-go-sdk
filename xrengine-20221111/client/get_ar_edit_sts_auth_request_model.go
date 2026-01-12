// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArEditStsAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *GetArEditStsAuthRequest
	GetJwtToken() *string
	SetMapId(v int64) *GetArEditStsAuthRequest
	GetMapId() *int64
}

type GetArEditStsAuthRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	MapId    *int64  `json:"MapId,omitempty" xml:"MapId,omitempty"`
}

func (s GetArEditStsAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArEditStsAuthRequest) GoString() string {
	return s.String()
}

func (s *GetArEditStsAuthRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *GetArEditStsAuthRequest) GetMapId() *int64 {
	return s.MapId
}

func (s *GetArEditStsAuthRequest) SetJwtToken(v string) *GetArEditStsAuthRequest {
	s.JwtToken = &v
	return s
}

func (s *GetArEditStsAuthRequest) SetMapId(v int64) *GetArEditStsAuthRequest {
	s.MapId = &v
	return s
}

func (s *GetArEditStsAuthRequest) Validate() error {
	return dara.Validate(s)
}
