// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocationTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateLocationTreeRequest
	GetId() *int64
	SetJwtToken(v string) *UpdateLocationTreeRequest
	GetJwtToken() *string
	SetTreeConfig(v string) *UpdateLocationTreeRequest
	GetTreeConfig() *string
}

type UpdateLocationTreeRequest struct {
	// This parameter is required.
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	JwtToken   *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	TreeConfig *string `json:"TreeConfig,omitempty" xml:"TreeConfig,omitempty"`
}

func (s UpdateLocationTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocationTreeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLocationTreeRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateLocationTreeRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *UpdateLocationTreeRequest) GetTreeConfig() *string {
	return s.TreeConfig
}

func (s *UpdateLocationTreeRequest) SetId(v int64) *UpdateLocationTreeRequest {
	s.Id = &v
	return s
}

func (s *UpdateLocationTreeRequest) SetJwtToken(v string) *UpdateLocationTreeRequest {
	s.JwtToken = &v
	return s
}

func (s *UpdateLocationTreeRequest) SetTreeConfig(v string) *UpdateLocationTreeRequest {
	s.TreeConfig = &v
	return s
}

func (s *UpdateLocationTreeRequest) Validate() error {
	return dara.Validate(s)
}
