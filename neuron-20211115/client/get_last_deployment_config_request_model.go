// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLastDeploymentConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceGroupId(v int64) *GetLastDeploymentConfigRequest
	GetServiceGroupId() *int64
}

type GetLastDeploymentConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 101
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s GetLastDeploymentConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLastDeploymentConfigRequest) GoString() string {
	return s.String()
}

func (s *GetLastDeploymentConfigRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *GetLastDeploymentConfigRequest) SetServiceGroupId(v int64) *GetLastDeploymentConfigRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *GetLastDeploymentConfigRequest) Validate() error {
	return dara.Validate(s)
}
