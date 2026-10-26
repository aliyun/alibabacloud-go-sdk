// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostShareKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListHostShareKeysRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListHostShareKeysRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostShareKeysRequest
	GetPageSize() *string
	SetProjectId(v int64) *ListHostShareKeysRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListHostShareKeysRequest
	GetRegionId() *string
}

type ListHostShareKeysRequest struct {
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Specifies the Settings for the number of host shared keys to return on each page during a paged query. Default value: **20**, which indicates that 20 host shared keys are returned per page during paging.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListHostShareKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostShareKeysRequest) GoString() string {
	return s.String()
}

func (s *ListHostShareKeysRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostShareKeysRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostShareKeysRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostShareKeysRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListHostShareKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostShareKeysRequest) SetInstanceId(v string) *ListHostShareKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostShareKeysRequest) SetPageNumber(v string) *ListHostShareKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostShareKeysRequest) SetPageSize(v string) *ListHostShareKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostShareKeysRequest) SetProjectId(v int64) *ListHostShareKeysRequest {
	s.ProjectId = &v
	return s
}

func (s *ListHostShareKeysRequest) SetRegionId(v string) *ListHostShareKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostShareKeysRequest) Validate() error {
	return dara.Validate(s)
}
