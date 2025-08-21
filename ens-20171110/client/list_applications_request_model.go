// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersions(v string) *ListApplicationsRequest
	GetAppVersions() *string
	SetClusterNames(v string) *ListApplicationsRequest
	GetClusterNames() *string
	SetLevel(v string) *ListApplicationsRequest
	GetLevel() *string
	SetMaxDate(v string) *ListApplicationsRequest
	GetMaxDate() *string
	SetMinDate(v string) *ListApplicationsRequest
	GetMinDate() *string
	SetOutAppInfoParams(v string) *ListApplicationsRequest
	GetOutAppInfoParams() *string
	SetPageNumber(v int32) *ListApplicationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationsRequest
	GetPageSize() *int32
}

type ListApplicationsRequest struct {
	// The version number of the application. Separate multiple version numbers with commas (,). If you want to query data of all versions of applications, specify All for this parameter. By default, only data of applications in the stable versions are queried.
	//
	// example:
	//
	// v1,v2
	AppVersions *string `json:"AppVersions,omitempty" xml:"AppVersions,omitempty"`
	// The name of the application cluster. Separate multiple names with commas (,). If you want to query applications of all clusters in your account, specify All for this parameter. Default value: All.
	//
	// example:
	//
	// poc,pre
	ClusterNames *string `json:"ClusterNames,omitempty" xml:"ClusterNames,omitempty"`
	// The region level by which edge resources of the application are collected. The value is of the enumeration type. Valid values:
	//
	// 	- National: Chinese mainland
	//
	// 	- Big: area
	//
	// 	- Middle: province
	//
	// 	- Small: city
	//
	// 	- RegionId: edge node
	//
	// Default value: National.
	//
	// example:
	//
	// National
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The end of the time range to query. Specify the time in the 2006-01-02 format. By default, the time range to query is not restricted.
	//
	// example:
	//
	// 2022-02-20
	MaxDate *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	// The beginning of the time range to query. Specify the time in the 2006-01-02 format. By default, the time range to query is not restricted.
	//
	// example:
	//
	// 2022-02-15
	MinDate *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	// Specifies whether to return other information about the application, such as statistics on resource instances and pods. The value must be a JSON string. By default, all information is returned.
	//
	// example:
	//
	// {\\"appInfo\\":true,\\"detailStat\\": true, \\"appVersionStat\\": true, \\"districtStat\\":true ,\\"instanceStat\\": true, \\"podCountStat\\": true}
	OutAppInfoParams *string `json:"OutAppInfoParams,omitempty" xml:"OutAppInfoParams,omitempty"`
	// The page number. Pages start from page 1. This parameter is optional if you want to return the push status of all data files.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. This parameter is optional if you want to return all information about the applications.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) GetAppVersions() *string {
	return s.AppVersions
}

func (s *ListApplicationsRequest) GetClusterNames() *string {
	return s.ClusterNames
}

func (s *ListApplicationsRequest) GetLevel() *string {
	return s.Level
}

func (s *ListApplicationsRequest) GetMaxDate() *string {
	return s.MaxDate
}

func (s *ListApplicationsRequest) GetMinDate() *string {
	return s.MinDate
}

func (s *ListApplicationsRequest) GetOutAppInfoParams() *string {
	return s.OutAppInfoParams
}

func (s *ListApplicationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationsRequest) SetAppVersions(v string) *ListApplicationsRequest {
	s.AppVersions = &v
	return s
}

func (s *ListApplicationsRequest) SetClusterNames(v string) *ListApplicationsRequest {
	s.ClusterNames = &v
	return s
}

func (s *ListApplicationsRequest) SetLevel(v string) *ListApplicationsRequest {
	s.Level = &v
	return s
}

func (s *ListApplicationsRequest) SetMaxDate(v string) *ListApplicationsRequest {
	s.MaxDate = &v
	return s
}

func (s *ListApplicationsRequest) SetMinDate(v string) *ListApplicationsRequest {
	s.MinDate = &v
	return s
}

func (s *ListApplicationsRequest) SetOutAppInfoParams(v string) *ListApplicationsRequest {
	s.OutAppInfoParams = &v
	return s
}

func (s *ListApplicationsRequest) SetPageNumber(v int32) *ListApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsRequest) SetPageSize(v int32) *ListApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
