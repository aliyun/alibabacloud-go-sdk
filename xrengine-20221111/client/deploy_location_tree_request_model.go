// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployLocationTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceUpdate(v bool) *DeployLocationTreeRequest
	GetForceUpdate() *bool
	SetJwtToken(v string) *DeployLocationTreeRequest
	GetJwtToken() *string
	SetSvcId(v int64) *DeployLocationTreeRequest
	GetSvcId() *int64
}

type DeployLocationTreeRequest struct {
	ForceUpdate *bool   `json:"ForceUpdate,omitempty" xml:"ForceUpdate,omitempty"`
	JwtToken    *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	SvcId       *int64  `json:"SvcId,omitempty" xml:"SvcId,omitempty"`
}

func (s DeployLocationTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployLocationTreeRequest) GoString() string {
	return s.String()
}

func (s *DeployLocationTreeRequest) GetForceUpdate() *bool {
	return s.ForceUpdate
}

func (s *DeployLocationTreeRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *DeployLocationTreeRequest) GetSvcId() *int64 {
	return s.SvcId
}

func (s *DeployLocationTreeRequest) SetForceUpdate(v bool) *DeployLocationTreeRequest {
	s.ForceUpdate = &v
	return s
}

func (s *DeployLocationTreeRequest) SetJwtToken(v string) *DeployLocationTreeRequest {
	s.JwtToken = &v
	return s
}

func (s *DeployLocationTreeRequest) SetSvcId(v int64) *DeployLocationTreeRequest {
	s.SvcId = &v
	return s
}

func (s *DeployLocationTreeRequest) Validate() error {
	return dara.Validate(s)
}
