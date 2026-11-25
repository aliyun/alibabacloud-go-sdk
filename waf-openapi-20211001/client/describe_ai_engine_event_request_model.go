// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *DescribeAiEngineEventRequest
	GetEventId() *string
	SetInstanceId(v string) *DescribeAiEngineEventRequest
	GetInstanceId() *string
	SetLanguage(v string) *DescribeAiEngineEventRequest
	GetLanguage() *string
	SetRegionId(v string) *DescribeAiEngineEventRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeAiEngineEventRequest struct {
	// The event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EVT-0bda6****6b5cwwz
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_dddsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The response language.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The region ID of the WAF instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-****mo
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeAiEngineEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeAiEngineEventRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAiEngineEventRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAiEngineEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAiEngineEventRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAiEngineEventRequest) SetEventId(v string) *DescribeAiEngineEventRequest {
	s.EventId = &v
	return s
}

func (s *DescribeAiEngineEventRequest) SetInstanceId(v string) *DescribeAiEngineEventRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAiEngineEventRequest) SetLanguage(v string) *DescribeAiEngineEventRequest {
	s.Language = &v
	return s
}

func (s *DescribeAiEngineEventRequest) SetRegionId(v string) *DescribeAiEngineEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAiEngineEventRequest) SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEventRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAiEngineEventRequest) Validate() error {
	return dara.Validate(s)
}
