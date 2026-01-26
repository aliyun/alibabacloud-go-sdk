// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRetcodeAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListRetcodeAppsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListRetcodeAppsRequest
	GetResourceGroupId() *string
	SetTags(v []*ListRetcodeAppsRequestTags) *ListRetcodeAppsRequest
	GetTags() []*ListRetcodeAppsRequestTags
}

type ListRetcodeAppsRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. You can obtain the resource group ID in the **Resource Management*	- console.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that you want to add to the task.
	Tags []*ListRetcodeAppsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListRetcodeAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRetcodeAppsRequest) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRetcodeAppsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListRetcodeAppsRequest) GetTags() []*ListRetcodeAppsRequestTags {
	return s.Tags
}

func (s *ListRetcodeAppsRequest) SetRegionId(v string) *ListRetcodeAppsRequest {
	s.RegionId = &v
	return s
}

func (s *ListRetcodeAppsRequest) SetResourceGroupId(v string) *ListRetcodeAppsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRetcodeAppsRequest) SetTags(v []*ListRetcodeAppsRequestTags) *ListRetcodeAppsRequest {
	s.Tags = v
	return s
}

func (s *ListRetcodeAppsRequest) Validate() error {
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

type ListRetcodeAppsRequestTags struct {
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

func (s ListRetcodeAppsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListRetcodeAppsRequestTags) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListRetcodeAppsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListRetcodeAppsRequestTags) SetKey(v string) *ListRetcodeAppsRequestTags {
	s.Key = &v
	return s
}

func (s *ListRetcodeAppsRequestTags) SetValue(v string) *ListRetcodeAppsRequestTags {
	s.Value = &v
	return s
}

func (s *ListRetcodeAppsRequestTags) Validate() error {
	return dara.Validate(s)
}
