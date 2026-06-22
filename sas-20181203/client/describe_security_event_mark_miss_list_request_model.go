// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventMarkMissListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSecurityEventMarkMissListRequest
	GetCurrentPage() *int32
	SetEventName(v string) *DescribeSecurityEventMarkMissListRequest
	GetEventName() *string
	SetPageSize(v int32) *DescribeSecurityEventMarkMissListRequest
	GetPageSize() *int32
	SetRemark(v string) *DescribeSecurityEventMarkMissListRequest
	GetRemark() *string
	SetResourceOwnerId(v int64) *DescribeSecurityEventMarkMissListRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeSecurityEventMarkMissListRequest
	GetSourceIp() *string
}

type DescribeSecurityEventMarkMissListRequest struct {
	// The page number of the page to return. Default value: **1**, which indicates that the first page is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The alert event name (child class).
	//
	// example:
	//
	// Login with unusual location
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The number of whitelist rules to display on each page in a paged query. Default value: **20**, which indicates that 20 whitelist rules are displayed on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The asset search keyword. You can enter the IP address, public IP address, private IP address, or asset name for fuzzy matching.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 125.210.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSecurityEventMarkMissListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventMarkMissListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventMarkMissListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSecurityEventMarkMissListRequest) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSecurityEventMarkMissListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSecurityEventMarkMissListRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeSecurityEventMarkMissListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecurityEventMarkMissListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSecurityEventMarkMissListRequest) SetCurrentPage(v int32) *DescribeSecurityEventMarkMissListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetEventName(v string) *DescribeSecurityEventMarkMissListRequest {
	s.EventName = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetPageSize(v int32) *DescribeSecurityEventMarkMissListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetRemark(v string) *DescribeSecurityEventMarkMissListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetResourceOwnerId(v int64) *DescribeSecurityEventMarkMissListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetSourceIp(v string) *DescribeSecurityEventMarkMissListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) Validate() error {
	return dara.Validate(s)
}
