// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopOnlineTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *StopOnlineTestRequest
	GetResourceType() *string
	SetServiceCode(v string) *StopOnlineTestRequest
	GetServiceCode() *string
	SetTaskId(v string) *StopOnlineTestRequest
	GetTaskId() *string
}

type StopOnlineTestRequest struct {
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// VideoModeration
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopOnlineTestRequest) String() string {
	return dara.Prettify(s)
}

func (s StopOnlineTestRequest) GoString() string {
	return s.String()
}

func (s *StopOnlineTestRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *StopOnlineTestRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *StopOnlineTestRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopOnlineTestRequest) SetResourceType(v string) *StopOnlineTestRequest {
	s.ResourceType = &v
	return s
}

func (s *StopOnlineTestRequest) SetServiceCode(v string) *StopOnlineTestRequest {
	s.ServiceCode = &v
	return s
}

func (s *StopOnlineTestRequest) SetTaskId(v string) *StopOnlineTestRequest {
	s.TaskId = &v
	return s
}

func (s *StopOnlineTestRequest) Validate() error {
	return dara.Validate(s)
}
