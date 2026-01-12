// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArEditUgcMaterialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *GetArEditUgcMaterialRequest
	GetJwtToken() *string
	SetMapId(v int64) *GetArEditUgcMaterialRequest
	GetMapId() *int64
}

type GetArEditUgcMaterialRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	MapId    *int64  `json:"MapId,omitempty" xml:"MapId,omitempty"`
}

func (s GetArEditUgcMaterialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArEditUgcMaterialRequest) GoString() string {
	return s.String()
}

func (s *GetArEditUgcMaterialRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *GetArEditUgcMaterialRequest) GetMapId() *int64 {
	return s.MapId
}

func (s *GetArEditUgcMaterialRequest) SetJwtToken(v string) *GetArEditUgcMaterialRequest {
	s.JwtToken = &v
	return s
}

func (s *GetArEditUgcMaterialRequest) SetMapId(v int64) *GetArEditUgcMaterialRequest {
	s.MapId = &v
	return s
}

func (s *GetArEditUgcMaterialRequest) Validate() error {
	return dara.Validate(s)
}
