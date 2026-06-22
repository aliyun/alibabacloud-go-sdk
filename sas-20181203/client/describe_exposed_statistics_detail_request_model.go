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
	SetExposureIp(v string) *DescribeExposedStatisticsDetailRequest
	GetExposureIp() *string
	SetInstanceId(v string) *DescribeExposedStatisticsDetailRequest
	GetInstanceId() *string
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
	// The query condition.
	//
	// example:
	//
	// {}
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The public IP address that is exposed on the Internet for the asset to query.
	//
	// example:
	//
	// 116.12.XX.XX
	ExposureIp *string `json:"ExposureIp,omitempty" xml:"ExposureIp,omitempty"`
	// The instance ID of the asset to query.
	//
	// example:
	//
	// s-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries per page in a paged query. Default value: 20. If you leave this parameter empty, 20 entries are returned per page.
	//
	// > Set PageSize to a non-empty value.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud account ID of the member accounts in the resource folder.
	//
	// > Invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The type of statistics to query. Valid values:
	//
	// - **exposureType**: gateway assets exposed on the Internet.
	//
	// - **exposurePort**: ports exposed on the Internet.
	//
	// - **exposureComponent**: system components exposed on the Internet.
	//
	// - **exposureIp**: IP addresses exposed on the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// exposureType
	StatisticsType *string `json:"StatisticsType,omitempty" xml:"StatisticsType,omitempty"`
	// The Asset Type of the gateway to query. This parameter takes effect only when **StatisticsType*	- is set to **exposureType**. Valid values:
	//
	// - **SLB**: public IP address of a load balancing SLB instance.
	//
	// - **DNAT**: NAT gateway that uses the DNAT feature to connect to the Internet.
	//
	// example:
	//
	// SLB
	StatisticsTypeGatewayType *string `json:"StatisticsTypeGatewayType,omitempty" xml:"StatisticsTypeGatewayType,omitempty"`
	// The instance ID of the gateway to query. This parameter takes effect only when **StatisticsType*	- is set to **exposureType**.
	//
	// example:
	//
	// lb-2ze4rso39h4nczcqs****
	StatisticsTypeInstanceValue *string `json:"StatisticsTypeInstanceValue,omitempty" xml:"StatisticsTypeInstanceValue,omitempty"`
	// The UUID of the server to query.
	//
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

func (s *DescribeExposedStatisticsDetailRequest) GetExposureIp() *string {
	return s.ExposureIp
}

func (s *DescribeExposedStatisticsDetailRequest) GetInstanceId() *string {
	return s.InstanceId
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

func (s *DescribeExposedStatisticsDetailRequest) SetExposureIp(v string) *DescribeExposedStatisticsDetailRequest {
	s.ExposureIp = &v
	return s
}

func (s *DescribeExposedStatisticsDetailRequest) SetInstanceId(v string) *DescribeExposedStatisticsDetailRequest {
	s.InstanceId = &v
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
