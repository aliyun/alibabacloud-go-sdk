// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficControlTaskTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *GetTrafficControlTaskTrafficRequest
	GetEnvironment() *string
	SetInstanceId(v string) *GetTrafficControlTaskTrafficRequest
	GetInstanceId() *string
}

type GetTrafficControlTaskTrafficRequest struct {
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// pairec-test-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTrafficControlTaskTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTaskTrafficRequest) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskTrafficRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *GetTrafficControlTaskTrafficRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTrafficControlTaskTrafficRequest) SetEnvironment(v string) *GetTrafficControlTaskTrafficRequest {
	s.Environment = &v
	return s
}

func (s *GetTrafficControlTaskTrafficRequest) SetInstanceId(v string) *GetTrafficControlTaskTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTrafficControlTaskTrafficRequest) Validate() error {
	return dara.Validate(s)
}
