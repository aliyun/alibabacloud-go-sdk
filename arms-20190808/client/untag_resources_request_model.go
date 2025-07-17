// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesRequest
	GetAll() *bool
	SetResourceId(v []*string) *UntagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagResourcesRequest
	GetTagKey() []*string
	SetTags(v []*UntagResourcesRequestTags) *UntagResourcesRequest
	GetTags() []*UntagResourcesRequestTags
}

type UntagResourcesRequest struct {
	// Specifies whether to delete all tags. This parameter takes effect only when the TagKey.N parameter is not specified. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// False
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The resource IDs. You can specify a maximum of 50 resource IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the ARMS resources for which you want to modify tags. Valid values:
	//
	// 	- WEB: Browser Monitoring
	//
	// 	- APPLICATION: Application Monitoring
	//
	// 	- PROMETHEUS: Managed Service for Prometheus
	//
	// 	- SYNTHETICTASK: Synthetic Monitoring
	//
	// 	- ALERTRULE: Application Monitoring alert rule
	//
	// 	- PROMETHEUSALERTRULE: Managed Service for Prometheus alert rule
	//
	// 	- XTRACEAPP: Managed Service for OpenTelemetry
	//
	// This parameter is required.
	//
	// example:
	//
	// PROMETHEUS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys. You can specify a maximum of 20 tag keys.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	// The list of tags.
	Tags []*UntagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UntagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UntagResourcesRequest) GetTags() []*UntagResourcesRequestTags {
	return s.Tags
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) SetTags(v []*UntagResourcesRequestTags) *UntagResourcesRequest {
	s.Tags = v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type UntagResourcesRequestTags struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UntagResourcesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequestTags) GetKey() *string {
	return s.Key
}

func (s *UntagResourcesRequestTags) GetValue() *string {
	return s.Value
}

func (s *UntagResourcesRequestTags) SetKey(v string) *UntagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *UntagResourcesRequestTags) SetValue(v string) *UntagResourcesRequestTags {
	s.Value = &v
	return s
}

func (s *UntagResourcesRequestTags) Validate() error {
	return dara.Validate(s)
}
