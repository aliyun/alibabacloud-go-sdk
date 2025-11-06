// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedStatisticsDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeExposedStatisticsDetailRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *DescribeExposedStatisticsDetailRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeExposedStatisticsDetailRequest
	GetPageSize() *int32
	SetResourceDirectoryAccountId(v int64) *DescribeExposedStatisticsDetailRequest
	GetResourceDirectoryAccountId() *int64
	SetStatisticsType(v string) *DescribeExposedStatisticsDetailRequest
	GetStatisticsType() *string
	SetStatisticsTypeGatewayType(v string) *DescribeExposedStatisticsDetailRequest
	GetStatisticsTypeGatewayType() *string
	SetStatisticsTypeInstanceValue(v string) *DescribeExposedStatisticsDetailRequest
	GetStatisticsTypeInstanceValue() *string
	SetUuid(v string) *DescribeExposedStatisticsDetailRequest
	GetUuid() *string
}

type DescribeExposedStatisticsDetailRequest struct {
	// example:
	//
	// {}
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the ID.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The type of the exposed asset. Valid values:
	//
	// 	- **exposureType**: gateway assets
	//
	// 	- **exposurePort**: ports
	//
	// 	- **exposureComponent**: system components
	//
	// 	- **exposureIp**: IP addresses
	//
	// This parameter is required.
	//
	// example:
	//
	// exposureType
	StatisticsType *string `json:"StatisticsType,omitempty" xml:"StatisticsType,omitempty"`
	// The type of the gateway asset. This parameter is required when the **StatisticsType*	- parameter is set to **exposureType**. Valid values:
	//
	// 	- **SLB**: the public IP address of a Server Load Balancer (SLB) instance
	//
	// 	- **DNAT**: the NAT gateway that connects to the Internet by using the DNAT feature
	//
	// example:
	//
	// SLB
	StatisticsTypeGatewayType *string `json:"StatisticsTypeGatewayType,omitempty" xml:"StatisticsTypeGatewayType,omitempty"`
	// The ID of the gateway asset. This parameter is required when the **StatisticsType*	- parameter is set to **exposureType**.
	//
	// example:
	//
	// lb-2ze4rso39h4nczcqs****
	StatisticsTypeInstanceValue *string `json:"StatisticsTypeInstanceValue,omitempty" xml:"StatisticsTypeInstanceValue,omitempty"`
	// example:
	//
	// c9107c04-942f-40c1-981a-f1c1***
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExposedStatisticsDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedStatisticsDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeExposedStatisticsDetailRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeExposedStatisticsDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeExposedStatisticsDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExposedStatisticsDetailRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeExposedStatisticsDetailRequest) GetStatisticsType() *string {
	return s.StatisticsType
}

func (s *DescribeExposedStatisticsDetailRequest) GetStatisticsTypeGatewayType() *string {
	return s.StatisticsTypeGatewayType
}

func (s *DescribeExposedStatisticsDetailRequest) GetStatisticsTypeInstanceValue() *string {
	return s.StatisticsTypeInstanceValue
}

func (s *DescribeExposedStatisticsDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeExposedStatisticsDetailRequest) SetCriteria(v string) *DescribeExposedStatisticsDetailRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeExposedStatisticsDetailRequest) SetCurrentPage(v int32) *DescribeExposedStatisticsDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeExposedStatisticsDetailRequest) SetPageSize(v int32) *DescribeExposedStatisticsDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeExposedStatisticsDetailRequest) SetResourceDirectoryAccountId(v int64) *DescribeExposedStatisticsDetailRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeExposedStatisticsDetailRequest) SetStatisticsType(v string) *DescribeExposedStatisticsDetailRequest {
	s.StatisticsType = &v
	return s
}

func (s *DescribeExposedStatisticsDetailRequest) SetStatisticsTypeGatewayType(v string) *DescribeExposedStatisticsDetailRequest {
	s.StatisticsTypeGatewayType = &v
	return s
}

func (s *DescribeExposedStatisticsDetailRequest) SetStatisticsTypeInstanceValue(v string) *DescribeExposedStatisticsDetailRequest {
	s.StatisticsTypeInstanceValue = &v
	return s
}

func (s *DescribeExposedStatisticsDetailRequest) SetUuid(v string) *DescribeExposedStatisticsDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedStatisticsDetailRequest) Validate() error {
	return dara.Validate(s)
}
