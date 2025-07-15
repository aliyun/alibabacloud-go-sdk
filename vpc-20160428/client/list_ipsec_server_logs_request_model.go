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
	// The beginning of the time range to query. The value must be a UNIX timestamp. For example, 1671003744 specifies 15:42:24 (UTC+8) on December 14, 2022.
	//
	// >  If you specify **From**, you must also specify **To*	- or **MinutePeriod**.
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
	// The interval at which log data is queried. Valid values: **1*	- to **10**. Unit: minutes.
	//
	// >  If both **From*	- and **To*	- are not specified, you must specify **MinutePeriod**.
	//
	// example:
	//
	// 10
	MinutePeriod *int32 `json:"MinutePeriod,omitempty" xml:"MinutePeriod,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the IPsec server is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The end of the time range to query. The value must be a unix timestamp. For example, 1671004344 specifies 15:52:24 (UTC+8) on December 14, 2022.
	//
	// >  If you specify **To**, you must also specify **From*	- or **MinutePeriod**.
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
