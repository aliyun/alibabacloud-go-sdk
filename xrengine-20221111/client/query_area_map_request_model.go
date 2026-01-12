// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAreaMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *QueryAreaMapRequest
	GetId() *int64
	SetJwtToken(v string) *QueryAreaMapRequest
	GetJwtToken() *string
}

type QueryAreaMapRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s QueryAreaMapRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAreaMapRequest) GoString() string {
	return s.String()
}

func (s *QueryAreaMapRequest) GetId() *int64 {
	return s.Id
}

func (s *QueryAreaMapRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *QueryAreaMapRequest) SetId(v int64) *QueryAreaMapRequest {
	s.Id = &v
	return s
}

func (s *QueryAreaMapRequest) SetJwtToken(v string) *QueryAreaMapRequest {
	s.JwtToken = &v
	return s
}

func (s *QueryAreaMapRequest) Validate() error {
	return dara.Validate(s)
}
