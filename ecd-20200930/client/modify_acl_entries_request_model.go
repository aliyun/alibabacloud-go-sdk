// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAclEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v string) *ModifyAclEntriesRequest
	GetPolicy() *string
	SetRegionId(v string) *ModifyAclEntriesRequest
	GetRegionId() *string
	SetSourceId(v []*string) *ModifyAclEntriesRequest
	GetSourceId() []*string
	SetSourceType(v string) *ModifyAclEntriesRequest
	GetSourceType() *string
}

type ModifyAclEntriesRequest struct {
	// The public network access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of instance IDs for public network access control, which are office network IDs or cloud computer IDs.
	//
	// This parameter is required.
	SourceId []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
	// The granularity of the public network access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// desktop
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ModifyAclEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAclEntriesRequest) GoString() string {
	return s.String()
}

func (s *ModifyAclEntriesRequest) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyAclEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAclEntriesRequest) GetSourceId() []*string {
	return s.SourceId
}

func (s *ModifyAclEntriesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ModifyAclEntriesRequest) SetPolicy(v string) *ModifyAclEntriesRequest {
	s.Policy = &v
	return s
}

func (s *ModifyAclEntriesRequest) SetRegionId(v string) *ModifyAclEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAclEntriesRequest) SetSourceId(v []*string) *ModifyAclEntriesRequest {
	s.SourceId = v
	return s
}

func (s *ModifyAclEntriesRequest) SetSourceType(v string) *ModifyAclEntriesRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyAclEntriesRequest) Validate() error {
	return dara.Validate(s)
}
