// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTrafficControlTaskDeployResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *QueryTrafficControlTaskDeployResultRequest
	GetEnvironment() *string
	SetInstanceId(v string) *QueryTrafficControlTaskDeployResultRequest
	GetInstanceId() *string
}

type QueryTrafficControlTaskDeployResultRequest struct {
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// pairec_123****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryTrafficControlTaskDeployResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTaskDeployResultRequest) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTaskDeployResultRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *QueryTrafficControlTaskDeployResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTrafficControlTaskDeployResultRequest) SetEnvironment(v string) *QueryTrafficControlTaskDeployResultRequest {
	s.Environment = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultRequest) SetInstanceId(v string) *QueryTrafficControlTaskDeployResultRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultRequest) Validate() error {
	return dara.Validate(s)
}
