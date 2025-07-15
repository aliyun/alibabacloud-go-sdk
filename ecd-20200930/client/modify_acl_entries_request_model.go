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
	// The Internet access control policy.
	//
	// Valid values:
	//
	// 	- allow: allows access to the Internet.
	//
	// 	- disable: forbids access to the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// allow
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance IDs (office network IDs or cloud computer IDs) to which the Internet access control policy is applicable.
	//
	// This parameter is required.
	SourceId []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
	// The granularity to which the Internet access control policy is applicable.
	//
	// Valid values:
	//
	// 	- desktop: cloud computer granularity.
	//
	// 	- vpc: office network granularity.
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
