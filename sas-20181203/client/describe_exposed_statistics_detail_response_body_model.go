// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedStatisticsDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeExposedStatisticsDetailResponseBodyPageInfo) *DescribeExposedStatisticsDetailResponseBody
	GetPageInfo() *DescribeExposedStatisticsDetailResponseBodyPageInfo
	SetRequestId(v string) *DescribeExposedStatisticsDetailResponseBody
	GetRequestId() *string
	SetStatisticsDetails(v []*DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) *DescribeExposedStatisticsDetailResponseBody
	GetStatisticsDetails() []*DescribeExposedStatisticsDetailResponseBodyStatisticsDetails
}

type DescribeExposedStatisticsDetailResponseBody struct {
	// The pagination information.
	PageInfo *DescribeExposedStatisticsDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7CBAFB3F-1ED7-4A23-986A-6F67F0466BD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of the gateway assets, ports, system components, or public IP addresses that are exposed on the Internet and are returned.
	StatisticsDetails []*DescribeExposedStatisticsDetailResponseBodyStatisticsDetails `json:"StatisticsDetails,omitempty" xml:"StatisticsDetails,omitempty" type:"Repeated"`
}

func (s DescribeExposedStatisticsDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedStatisticsDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposedStatisticsDetailResponseBody) GetPageInfo() *DescribeExposedStatisticsDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeExposedStatisticsDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExposedStatisticsDetailResponseBody) GetStatisticsDetails() []*DescribeExposedStatisticsDetailResponseBodyStatisticsDetails {
	return s.StatisticsDetails
}

func (s *DescribeExposedStatisticsDetailResponseBody) SetPageInfo(v *DescribeExposedStatisticsDetailResponseBodyPageInfo) *DescribeExposedStatisticsDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBody) SetRequestId(v string) *DescribeExposedStatisticsDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBody) SetStatisticsDetails(v []*DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) *DescribeExposedStatisticsDetailResponseBody {
	s.StatisticsDetails = v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExposedStatisticsDetailResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExposedStatisticsDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedStatisticsDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeExposedStatisticsDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeExposedStatisticsDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeExposedStatisticsDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExposedStatisticsDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeExposedStatisticsDetailResponseBodyPageInfo) SetCount(v int32) *DescribeExposedStatisticsDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeExposedStatisticsDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribeExposedStatisticsDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribeExposedStatisticsDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeExposedStatisticsDetailResponseBodyStatisticsDetails struct {
	// The total number of system vulnerabilities that are detected on your server and are exposed on the Internet.
	//
	// example:
	//
	// 0
	ExposedCount *int32 `json:"ExposedCount,omitempty" xml:"ExposedCount,omitempty"`
	// The system component that is exposed on the Internet.
	//
	// example:
	//
	// tomcat
	ExposureComponent *string `json:"ExposureComponent,omitempty" xml:"ExposureComponent,omitempty"`
	// The public IP address that is exposed on the Internet.
	//
	// example:
	//
	// 123.57.XX.XX
	ExposureIp *string `json:"ExposureIp,omitempty" xml:"ExposureIp,omitempty"`
	// The port that is exposed on the Internet.
	//
	// example:
	//
	// 22
	ExposurePort *string `json:"ExposurePort,omitempty" xml:"ExposurePort,omitempty"`
	// The resource from which the asset is exposed. Valid values:
	//
	// 	- **INTERNET_IP**: the IP address of the Elastic Compute Service (ECS) instance
	//
	// 	- **SLB**: the public IP address of the SLB instance
	//
	// 	- **EIP**: the elastic IP address (EIP)
	//
	// 	- **DNAT**: the NAT gateway that connects to the Internet by using the DNAT feature
	//
	// example:
	//
	// SLB
	ExposureType *string `json:"ExposureType,omitempty" xml:"ExposureType,omitempty"`
	// The ID of the instance to which the resource belongs. The valid values of this parameter vary based on the value of the ExposureType parameter.
	//
	// 	- If the value of the **ExposureType*	- parameter is **INTERNET_IP**, the value of this parameter is an empty string.
	//
	// 	- If the value of the **ExposureType*	- parameter is **SLB**, the value of this parameter is the ID of the Internet-facing SLB instance.
	//
	// 	- If the value of the **ExposureType*	- parameter is **EIP**, the value of this parameter is the ID of the EIP.
	//
	// 	- If the value of the **ExposureType*	- parameter is **DNAT**, the value of this parameter is the ID of the NAT gateway.
	//
	// example:
	//
	// lb-2ze4rso39h4nczcqs****
	ExposureTypeId *string `json:"ExposureTypeId,omitempty" xml:"ExposureTypeId,omitempty"`
	// The name of the gateway asset that is exposed on the Internet.
	//
	// example:
	//
	// ngw-bp1vkbju8f3w87c9v****
	ExposureTypeInstanceName *string `json:"ExposureTypeInstanceName,omitempty" xml:"ExposureTypeInstanceName,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) GoString() string {
	return s.String()
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) GetExposedCount() *int32 {
	return s.ExposedCount
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) GetExposureComponent() *string {
	return s.ExposureComponent
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) GetExposureIp() *string {
	return s.ExposureIp
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) GetExposurePort() *string {
	return s.ExposurePort
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) GetExposureType() *string {
	return s.ExposureType
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) GetExposureTypeId() *string {
	return s.ExposureTypeId
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) GetExposureTypeInstanceName() *string {
	return s.ExposureTypeInstanceName
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) SetExposedCount(v int32) *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails {
	s.ExposedCount = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) SetExposureComponent(v string) *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails {
	s.ExposureComponent = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) SetExposureIp(v string) *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails {
	s.ExposureIp = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) SetExposurePort(v string) *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails {
	s.ExposurePort = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) SetExposureType(v string) *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails {
	s.ExposureType = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) SetExposureTypeId(v string) *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails {
	s.ExposureTypeId = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) SetExposureTypeInstanceName(v string) *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails {
	s.ExposureTypeInstanceName = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) SetRegionId(v string) *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails {
	s.RegionId = &v
	return s
}

func (s *DescribeExposedStatisticsDetailResponseBodyStatisticsDetails) Validate() error {
	return dara.Validate(s)
}
