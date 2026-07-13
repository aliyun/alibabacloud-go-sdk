// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetToolCallDistributionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetToolCallDistributionRequest
	GetEndTime() *string
	SetInstanceId(v string) *GetToolCallDistributionRequest
	GetInstanceId() *string
	SetStartTime(v string) *GetToolCallDistributionRequest
	GetStartTime() *string
}

type GetToolCallDistributionRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetToolCallDistributionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetToolCallDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetToolCallDistributionRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetToolCallDistributionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetToolCallDistributionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetToolCallDistributionRequest) SetEndTime(v string) *GetToolCallDistributionRequest {
	s.EndTime = &v
	return s
}

func (s *GetToolCallDistributionRequest) SetInstanceId(v string) *GetToolCallDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetToolCallDistributionRequest) SetStartTime(v string) *GetToolCallDistributionRequest {
	s.StartTime = &v
	return s
}

func (s *GetToolCallDistributionRequest) Validate() error {
	return dara.Validate(s)
}
