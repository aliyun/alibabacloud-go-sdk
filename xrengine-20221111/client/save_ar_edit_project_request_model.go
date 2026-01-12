// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveArEditProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *SaveArEditProjectRequest
	GetJwtToken() *string
	SetMapId(v int64) *SaveArEditProjectRequest
	GetMapId() *int64
}

type SaveArEditProjectRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	MapId    *int64  `json:"MapId,omitempty" xml:"MapId,omitempty"`
}

func (s SaveArEditProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveArEditProjectRequest) GoString() string {
	return s.String()
}

func (s *SaveArEditProjectRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *SaveArEditProjectRequest) GetMapId() *int64 {
	return s.MapId
}

func (s *SaveArEditProjectRequest) SetJwtToken(v string) *SaveArEditProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *SaveArEditProjectRequest) SetMapId(v int64) *SaveArEditProjectRequest {
	s.MapId = &v
	return s
}

func (s *SaveArEditProjectRequest) Validate() error {
	return dara.Validate(s)
}
