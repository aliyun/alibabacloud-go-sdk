// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishArEditProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *PublishArEditProjectRequest
	GetJwtToken() *string
	SetMapId(v int64) *PublishArEditProjectRequest
	GetMapId() *int64
}

type PublishArEditProjectRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	MapId    *int64  `json:"MapId,omitempty" xml:"MapId,omitempty"`
}

func (s PublishArEditProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishArEditProjectRequest) GoString() string {
	return s.String()
}

func (s *PublishArEditProjectRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *PublishArEditProjectRequest) GetMapId() *int64 {
	return s.MapId
}

func (s *PublishArEditProjectRequest) SetJwtToken(v string) *PublishArEditProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *PublishArEditProjectRequest) SetMapId(v int64) *PublishArEditProjectRequest {
	s.MapId = &v
	return s
}

func (s *PublishArEditProjectRequest) Validate() error {
	return dara.Validate(s)
}
