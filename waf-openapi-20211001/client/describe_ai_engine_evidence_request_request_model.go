// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEvidenceRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *DescribeAiEngineEvidenceRequestRequest
	GetEventId() *string
	SetEvidenceId(v string) *DescribeAiEngineEvidenceRequestRequest
	GetEvidenceId() *string
	SetInstanceId(v string) *DescribeAiEngineEvidenceRequestRequest
	GetInstanceId() *string
	SetLanguage(v string) *DescribeAiEngineEvidenceRequestRequest
	GetLanguage() *string
	SetRegionId(v string) *DescribeAiEngineEvidenceRequestRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEvidenceRequestRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeAiEngineEvidenceRequestRequest struct {
	// The event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EVT-20241025-109
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The evidence ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EVD-20241025-001
	EvidenceId *string `json:"EvidenceId,omitempty" xml:"EvidenceId,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-sssde*****
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
	// rg-demo****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestRequest) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeAiEngineEvidenceRequestRequest) GetEvidenceId() *string {
	return s.EvidenceId
}

func (s *DescribeAiEngineEvidenceRequestRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAiEngineEvidenceRequestRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAiEngineEvidenceRequestRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAiEngineEvidenceRequestRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAiEngineEvidenceRequestRequest) SetEventId(v string) *DescribeAiEngineEvidenceRequestRequest {
	s.EventId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestRequest) SetEvidenceId(v string) *DescribeAiEngineEvidenceRequestRequest {
	s.EvidenceId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestRequest) SetInstanceId(v string) *DescribeAiEngineEvidenceRequestRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestRequest) SetLanguage(v string) *DescribeAiEngineEvidenceRequestRequest {
	s.Language = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestRequest) SetRegionId(v string) *DescribeAiEngineEvidenceRequestRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestRequest) SetResourceManagerResourceGroupId(v string) *DescribeAiEngineEvidenceRequestRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestRequest) Validate() error {
	return dara.Validate(s)
}
