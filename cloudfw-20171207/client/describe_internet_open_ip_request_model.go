// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsInstanceId(v string) *DescribeInternetOpenIpRequest
	GetAssetsInstanceId() *string
	SetAssetsInstanceName(v string) *DescribeInternetOpenIpRequest
	GetAssetsInstanceName() *string
	SetAssetsType(v string) *DescribeInternetOpenIpRequest
	GetAssetsType() *string
	SetCurrentPage(v string) *DescribeInternetOpenIpRequest
	GetCurrentPage() *string
	SetEndTime(v string) *DescribeInternetOpenIpRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInternetOpenIpRequest
	GetLang() *string
	SetPageSize(v string) *DescribeInternetOpenIpRequest
	GetPageSize() *string
	SetPort(v string) *DescribeInternetOpenIpRequest
	GetPort() *string
	SetPublicIp(v string) *DescribeInternetOpenIpRequest
	GetPublicIp() *string
	SetRegionNo(v string) *DescribeInternetOpenIpRequest
	GetRegionNo() *string
	SetRiskLevel(v string) *DescribeInternetOpenIpRequest
	GetRiskLevel() *string
	SetServiceName(v string) *DescribeInternetOpenIpRequest
	GetServiceName() *string
	SetStartTime(v string) *DescribeInternetOpenIpRequest
	GetStartTime() *string
}

type DescribeInternetOpenIpRequest struct {
	// The instance ID.
	//
	// example:
	//
	// i-uf6faknmuby7ezht****
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// launch-advisor-2023****
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// The asset type of the instance.
	//
	// example:
	//
	// EcsEIP
	AssetsType *string `json:"AssetsType,omitempty" xml:"AssetsType,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1663640336
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The port number.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The public IP address of the instance.
	//
	// example:
	//
	// 203.0.113.1
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The risk level. If you leave this parameter empty, all risk levels are queried. Valid values:
	//
	// 	- **3**: high risk
	//
	// 	- **2**: medium risk
	//
	// 	- **1**: low risk
	//
	// 	- **0**: no risk
	//
	// example:
	//
	// 2
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The application.
	//
	// example:
	//
	// SSH
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1681957629
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInternetOpenIpRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenIpRequest) GetAssetsInstanceId() *string {
	return s.AssetsInstanceId
}

func (s *DescribeInternetOpenIpRequest) GetAssetsInstanceName() *string {
	return s.AssetsInstanceName
}

func (s *DescribeInternetOpenIpRequest) GetAssetsType() *string {
	return s.AssetsType
}

func (s *DescribeInternetOpenIpRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeInternetOpenIpRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInternetOpenIpRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetOpenIpRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInternetOpenIpRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeInternetOpenIpRequest) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeInternetOpenIpRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeInternetOpenIpRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeInternetOpenIpRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeInternetOpenIpRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInternetOpenIpRequest) SetAssetsInstanceId(v string) *DescribeInternetOpenIpRequest {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetAssetsInstanceName(v string) *DescribeInternetOpenIpRequest {
	s.AssetsInstanceName = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetAssetsType(v string) *DescribeInternetOpenIpRequest {
	s.AssetsType = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetCurrentPage(v string) *DescribeInternetOpenIpRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetEndTime(v string) *DescribeInternetOpenIpRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetLang(v string) *DescribeInternetOpenIpRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetPageSize(v string) *DescribeInternetOpenIpRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetPort(v string) *DescribeInternetOpenIpRequest {
	s.Port = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetPublicIp(v string) *DescribeInternetOpenIpRequest {
	s.PublicIp = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetRegionNo(v string) *DescribeInternetOpenIpRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetRiskLevel(v string) *DescribeInternetOpenIpRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetServiceName(v string) *DescribeInternetOpenIpRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetStartTime(v string) *DescribeInternetOpenIpRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) Validate() error {
	return dara.Validate(s)
}
