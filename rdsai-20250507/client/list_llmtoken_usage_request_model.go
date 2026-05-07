// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLLMTokenUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListLLMTokenUsageRequest
	GetEndTime() *string
	SetInstanceName(v string) *ListLLMTokenUsageRequest
	GetInstanceName() *string
	SetModel(v string) *ListLLMTokenUsageRequest
	GetModel() *string
	SetRegionId(v string) *ListLLMTokenUsageRequest
	GetRegionId() *string
	SetStartTime(v string) *ListLLMTokenUsageRequest
	GetStartTime() *string
}

type ListLLMTokenUsageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2026-02-03T02:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// qwen-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-03-31T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListLLMTokenUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLLMTokenUsageRequest) GoString() string {
	return s.String()
}

func (s *ListLLMTokenUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListLLMTokenUsageRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListLLMTokenUsageRequest) GetModel() *string {
	return s.Model
}

func (s *ListLLMTokenUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLLMTokenUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListLLMTokenUsageRequest) SetEndTime(v string) *ListLLMTokenUsageRequest {
	s.EndTime = &v
	return s
}

func (s *ListLLMTokenUsageRequest) SetInstanceName(v string) *ListLLMTokenUsageRequest {
	s.InstanceName = &v
	return s
}

func (s *ListLLMTokenUsageRequest) SetModel(v string) *ListLLMTokenUsageRequest {
	s.Model = &v
	return s
}

func (s *ListLLMTokenUsageRequest) SetRegionId(v string) *ListLLMTokenUsageRequest {
	s.RegionId = &v
	return s
}

func (s *ListLLMTokenUsageRequest) SetStartTime(v string) *ListLLMTokenUsageRequest {
	s.StartTime = &v
	return s
}

func (s *ListLLMTokenUsageRequest) Validate() error {
	return dara.Validate(s)
}
