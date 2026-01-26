// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTraceAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *ListTraceAppsRequest
	GetAppType() *string
	SetRegion(v string) *ListTraceAppsRequest
	GetRegion() *string
	SetRegionId(v string) *ListTraceAppsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTraceAppsRequest
	GetResourceGroupId() *string
	SetTags(v []*ListTraceAppsRequestTags) *ListTraceAppsRequest
	GetTags() []*ListTraceAppsRequestTags
}

type ListTraceAppsRequest struct {
	// The type of the application that is associated with the alert rule. Valid values:
	//
	// - TRACE: Application Monitoring
	//
	// - EBPF: Application Monitoring eBPF Edition
	//
	// example:
	//
	// TRACE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListTraceAppsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTraceAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTraceAppsRequest) GoString() string {
	return s.String()
}

func (s *ListTraceAppsRequest) GetAppType() *string {
	return s.AppType
}

func (s *ListTraceAppsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListTraceAppsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTraceAppsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTraceAppsRequest) GetTags() []*ListTraceAppsRequestTags {
	return s.Tags
}

func (s *ListTraceAppsRequest) SetAppType(v string) *ListTraceAppsRequest {
	s.AppType = &v
	return s
}

func (s *ListTraceAppsRequest) SetRegion(v string) *ListTraceAppsRequest {
	s.Region = &v
	return s
}

func (s *ListTraceAppsRequest) SetRegionId(v string) *ListTraceAppsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTraceAppsRequest) SetResourceGroupId(v string) *ListTraceAppsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTraceAppsRequest) SetTags(v []*ListTraceAppsRequestTags) *ListTraceAppsRequest {
	s.Tags = v
	return s
}

func (s *ListTraceAppsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTraceAppsRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTraceAppsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListTraceAppsRequestTags) GoString() string {
	return s.String()
}

func (s *ListTraceAppsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListTraceAppsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListTraceAppsRequestTags) SetKey(v string) *ListTraceAppsRequestTags {
	s.Key = &v
	return s
}

func (s *ListTraceAppsRequestTags) SetValue(v string) *ListTraceAppsRequestTags {
	s.Value = &v
	return s
}

func (s *ListTraceAppsRequestTags) Validate() error {
	return dara.Validate(s)
}
