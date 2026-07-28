// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpsecServerLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int32) *ListIpsecServerLogsRequest
	GetFrom() *int32
	SetIpsecServerId(v string) *ListIpsecServerLogsRequest
	GetIpsecServerId() *string
	SetMinutePeriod(v int32) *ListIpsecServerLogsRequest
	GetMinutePeriod() *int32
	SetPageNumber(v int32) *ListIpsecServerLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIpsecServerLogsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListIpsecServerLogsRequest
	GetRegionId() *string
	SetTo(v int32) *ListIpsecServerLogsRequest
	GetTo() *int32
}

type ListIpsecServerLogsRequest struct {
	// The start time of the log. Only UNIX timestamps in seconds are supported. For example, 1671003744 represents 2022-12-14 15:42:24.
	//
	// > If you specify **From**, you must also specify **To*	- or **MinutePeriod**.
	//
	// example:
	//
	// 1671003744
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the IPsec server.
	//
	// This parameter is required.
	//
	// example:
	//
	// iss-2zei2n5q5zhirfh73****
	IpsecServerId *string `json:"IpsecServerId,omitempty" xml:"IpsecServerId,omitempty"`
	// The log period. Valid values: **1*	- to **10**. Unit: minutes.
	//
	// > If you do not specify **From*	- or **To**, you must specify **MinutePeriod**.
	//
	// example:
	//
	// 10
	MinutePeriod *int32 `json:"MinutePeriod,omitempty" xml:"MinutePeriod,omitempty"`
	// The page number of the list. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the IPsec server.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The end time of the log. Only UNIX timestamps in seconds are supported. For example, 1671004344 represents 2022-12-14 15:52:24.
	//
	// > If you specify **To**, you must also specify **From*	- or **MinutePeriod**.
	//
	// example:
	//
	// 1671004344
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s ListIpsecServerLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpsecServerLogsRequest) GoString() string {
	return s.String()
}

func (s *ListIpsecServerLogsRequest) GetFrom() *int32 {
	return s.From
}

func (s *ListIpsecServerLogsRequest) GetIpsecServerId() *string {
	return s.IpsecServerId
}

func (s *ListIpsecServerLogsRequest) GetMinutePeriod() *int32 {
	return s.MinutePeriod
}

func (s *ListIpsecServerLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIpsecServerLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIpsecServerLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpsecServerLogsRequest) GetTo() *int32 {
	return s.To
}

func (s *ListIpsecServerLogsRequest) SetFrom(v int32) *ListIpsecServerLogsRequest {
	s.From = &v
	return s
}

func (s *ListIpsecServerLogsRequest) SetIpsecServerId(v string) *ListIpsecServerLogsRequest {
	s.IpsecServerId = &v
	return s
}

func (s *ListIpsecServerLogsRequest) SetMinutePeriod(v int32) *ListIpsecServerLogsRequest {
	s.MinutePeriod = &v
	return s
}

func (s *ListIpsecServerLogsRequest) SetPageNumber(v int32) *ListIpsecServerLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIpsecServerLogsRequest) SetPageSize(v int32) *ListIpsecServerLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIpsecServerLogsRequest) SetRegionId(v string) *ListIpsecServerLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpsecServerLogsRequest) SetTo(v int32) *ListIpsecServerLogsRequest {
	s.To = &v
	return s
}

func (s *ListIpsecServerLogsRequest) Validate() error {
	return dara.Validate(s)
}
