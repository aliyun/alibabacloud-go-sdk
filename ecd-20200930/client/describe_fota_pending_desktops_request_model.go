// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFotaPendingDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *DescribeFotaPendingDesktopsRequest
	GetDesktopId() *string
	SetDesktopName(v string) *DescribeFotaPendingDesktopsRequest
	GetDesktopName() *string
	SetMaxResults(v int32) *DescribeFotaPendingDesktopsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeFotaPendingDesktopsRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeFotaPendingDesktopsRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DescribeFotaPendingDesktopsRequest
	GetRegionId() *string
	SetTaskUid(v string) *DescribeFotaPendingDesktopsRequest
	GetTaskUid() *string
}

type DescribeFotaPendingDesktopsRequest struct {
	// The ID of the cloud computer.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud computer.
	//
	// example:
	//
	// demo-computer
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The maximum number of entries per page for a paged query.
	//
	// - Maximum value: 100
	//
	// - Default value: 20
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set the value to the `NextToken` value returned from the previous call to this operation. You do not need to specify this parameter for the first call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the office network. You can call the [DescribeOfficeSites](https://help.aliyun.com/document_detail/216071.html) operation to obtain this parameter.
	//
	// example:
	//
	// cn-shanghai+dir-428216****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the list of regions supported by cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about the image upgrade task. You can call the [DescribeFotaTasks](https://help.aliyun.com/document_detail/437001.html) operation to obtain this parameter.
	//
	// example:
	//
	// aot-c4khwrp9ocml4****
	TaskUid *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s DescribeFotaPendingDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFotaPendingDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFotaPendingDesktopsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeFotaPendingDesktopsRequest) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeFotaPendingDesktopsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeFotaPendingDesktopsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFotaPendingDesktopsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeFotaPendingDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFotaPendingDesktopsRequest) GetTaskUid() *string {
	return s.TaskUid
}

func (s *DescribeFotaPendingDesktopsRequest) SetDesktopId(v string) *DescribeFotaPendingDesktopsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeFotaPendingDesktopsRequest) SetDesktopName(v string) *DescribeFotaPendingDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeFotaPendingDesktopsRequest) SetMaxResults(v int32) *DescribeFotaPendingDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFotaPendingDesktopsRequest) SetNextToken(v string) *DescribeFotaPendingDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeFotaPendingDesktopsRequest) SetOfficeSiteId(v string) *DescribeFotaPendingDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeFotaPendingDesktopsRequest) SetRegionId(v string) *DescribeFotaPendingDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFotaPendingDesktopsRequest) SetTaskUid(v string) *DescribeFotaPendingDesktopsRequest {
	s.TaskUid = &v
	return s
}

func (s *DescribeFotaPendingDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
