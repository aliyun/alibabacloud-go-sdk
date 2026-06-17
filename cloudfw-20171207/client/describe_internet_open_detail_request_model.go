// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsInstanceId(v string) *DescribeInternetOpenDetailRequest
	GetAssetsInstanceId() *string
	SetAssetsInstanceName(v string) *DescribeInternetOpenDetailRequest
	GetAssetsInstanceName() *string
	SetAssetsType(v string) *DescribeInternetOpenDetailRequest
	GetAssetsType() *string
	SetCurrentPage(v string) *DescribeInternetOpenDetailRequest
	GetCurrentPage() *string
	SetEndTime(v string) *DescribeInternetOpenDetailRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInternetOpenDetailRequest
	GetLang() *string
	SetPageSize(v string) *DescribeInternetOpenDetailRequest
	GetPageSize() *string
	SetPort(v string) *DescribeInternetOpenDetailRequest
	GetPort() *string
	SetPublicIp(v string) *DescribeInternetOpenDetailRequest
	GetPublicIp() *string
	SetRegionNo(v string) *DescribeInternetOpenDetailRequest
	GetRegionNo() *string
	SetRiskLevel(v string) *DescribeInternetOpenDetailRequest
	GetRiskLevel() *string
	SetServiceName(v string) *DescribeInternetOpenDetailRequest
	GetServiceName() *string
	SetServiceNameFuzzy(v string) *DescribeInternetOpenDetailRequest
	GetServiceNameFuzzy() *string
	SetSortList(v []*DescribeInternetOpenDetailRequestSortList) *DescribeInternetOpenDetailRequest
	GetSortList() []*DescribeInternetOpenDetailRequestSortList
	SetSourceIp(v string) *DescribeInternetOpenDetailRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInternetOpenDetailRequest
	GetStartTime() *string
	SetSuggestLevel(v string) *DescribeInternetOpenDetailRequest
	GetSuggestLevel() *string
}

type DescribeInternetOpenDetailRequest struct {
	// The ID of the asset. Fuzzy search is supported.
	//
	// example:
	//
	// i-uf6faknmuby7ezht****
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The name of the asset. Fuzzy search is supported.
	//
	// example:
	//
	// instance_test
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// The type of the asset for an exact match. If you leave this parameter empty, all asset types are queried.
	//
	// example:
	//
	// EcsPublicIP
	AssetsType *string `json:"AssetsType,omitempty" xml:"AssetsType,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1745251200
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The port for an exact match. The value must be an integer from 1 to 65535. If you leave this parameter empty, all ports are queried.
	//
	// example:
	//
	// 9100
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The public IP address for an exact match. If you leave this parameter empty, all public IP addresses are queried.
	//
	// example:
	//
	// 203.0.13.XX
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The risk level. If you leave this parameter empty, all risk levels are queried.
	//
	// example:
	//
	// 3
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The name of the application for an exact match. If you leave this parameter empty, all applications are queried.
	//
	// example:
	//
	// SMB
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The name of the application for a fuzzy match. If you leave this parameter empty, all applications are queried.
	//
	// example:
	//
	// SMB
	ServiceNameFuzzy *string `json:"ServiceNameFuzzy,omitempty" xml:"ServiceNameFuzzy,omitempty"`
	// The sorting conditions.
	SortList []*DescribeInternetOpenDetailRequestSortList `json:"SortList,omitempty" xml:"SortList,omitempty" type:"Repeated"`
	// The source IP address of the access request.
	//
	// example:
	//
	// 222.212.86.7XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The start of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1656837360
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The recommended policy level.
	//
	// example:
	//
	// 10
	SuggestLevel *string `json:"SuggestLevel,omitempty" xml:"SuggestLevel,omitempty"`
}

func (s DescribeInternetOpenDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenDetailRequest) GetAssetsInstanceId() *string {
	return s.AssetsInstanceId
}

func (s *DescribeInternetOpenDetailRequest) GetAssetsInstanceName() *string {
	return s.AssetsInstanceName
}

func (s *DescribeInternetOpenDetailRequest) GetAssetsType() *string {
	return s.AssetsType
}

func (s *DescribeInternetOpenDetailRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeInternetOpenDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInternetOpenDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetOpenDetailRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInternetOpenDetailRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeInternetOpenDetailRequest) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeInternetOpenDetailRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeInternetOpenDetailRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeInternetOpenDetailRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeInternetOpenDetailRequest) GetServiceNameFuzzy() *string {
	return s.ServiceNameFuzzy
}

func (s *DescribeInternetOpenDetailRequest) GetSortList() []*DescribeInternetOpenDetailRequestSortList {
	return s.SortList
}

func (s *DescribeInternetOpenDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetOpenDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInternetOpenDetailRequest) GetSuggestLevel() *string {
	return s.SuggestLevel
}

func (s *DescribeInternetOpenDetailRequest) SetAssetsInstanceId(v string) *DescribeInternetOpenDetailRequest {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetAssetsInstanceName(v string) *DescribeInternetOpenDetailRequest {
	s.AssetsInstanceName = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetAssetsType(v string) *DescribeInternetOpenDetailRequest {
	s.AssetsType = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetCurrentPage(v string) *DescribeInternetOpenDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetEndTime(v string) *DescribeInternetOpenDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetLang(v string) *DescribeInternetOpenDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetPageSize(v string) *DescribeInternetOpenDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetPort(v string) *DescribeInternetOpenDetailRequest {
	s.Port = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetPublicIp(v string) *DescribeInternetOpenDetailRequest {
	s.PublicIp = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetRegionNo(v string) *DescribeInternetOpenDetailRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetRiskLevel(v string) *DescribeInternetOpenDetailRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetServiceName(v string) *DescribeInternetOpenDetailRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetServiceNameFuzzy(v string) *DescribeInternetOpenDetailRequest {
	s.ServiceNameFuzzy = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetSortList(v []*DescribeInternetOpenDetailRequestSortList) *DescribeInternetOpenDetailRequest {
	s.SortList = v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetSourceIp(v string) *DescribeInternetOpenDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetStartTime(v string) *DescribeInternetOpenDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) SetSuggestLevel(v string) *DescribeInternetOpenDetailRequest {
	s.SuggestLevel = &v
	return s
}

func (s *DescribeInternetOpenDetailRequest) Validate() error {
	if s.SortList != nil {
		for _, item := range s.SortList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInternetOpenDetailRequestSortList struct {
	// The sort order.
	//
	// example:
	//
	// asc
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The sorting key.
	//
	// example:
	//
	// ServiceName
	SortKey *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
}

func (s DescribeInternetOpenDetailRequestSortList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenDetailRequestSortList) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenDetailRequestSortList) GetDir() *string {
	return s.Dir
}

func (s *DescribeInternetOpenDetailRequestSortList) GetSortKey() *string {
	return s.SortKey
}

func (s *DescribeInternetOpenDetailRequestSortList) SetDir(v string) *DescribeInternetOpenDetailRequestSortList {
	s.Dir = &v
	return s
}

func (s *DescribeInternetOpenDetailRequestSortList) SetSortKey(v string) *DescribeInternetOpenDetailRequestSortList {
	s.SortKey = &v
	return s
}

func (s *DescribeInternetOpenDetailRequestSortList) Validate() error {
	return dara.Validate(s)
}
