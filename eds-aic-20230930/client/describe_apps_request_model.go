// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIdList(v []*string) *DescribeAppsRequest
	GetAppIdList() []*string
	SetAppName(v string) *DescribeAppsRequest
	GetAppName() *string
	SetAppType(v string) *DescribeAppsRequest
	GetAppType() *string
	SetBizRegionId(v string) *DescribeAppsRequest
	GetBizRegionId() *string
	SetInstallationStatus(v string) *DescribeAppsRequest
	GetInstallationStatus() *string
	SetMD5(v string) *DescribeAppsRequest
	GetMD5() *string
	SetMaxResults(v int32) *DescribeAppsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAppsRequest
	GetNextToken() *string
	SetStatus(v string) *DescribeAppsRequest
	GetStatus() *string
}

type DescribeAppsRequest struct {
	// The IDs of the applications.
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	// The name of the application.
	//
	// example:
	//
	// defaultAppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// Region id.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The installation/uninstallation status of the application.
	//
	// Valid values:
	//
	// 	- INSTALLFAILED: The application failed to be installed.
	//
	// 	- UNINSTALLING: The application is being uninstalled.
	//
	// 	- INSTALLING: The application is being installed.
	//
	// 	- UNINSTALLED: The application is uninstalled.
	//
	// 	- INSTALLED: The application is installed.
	//
	// 	- UNINSTALLFAILED: The application failed to be uninstalled.
	//
	// example:
	//
	// INSTALLING
	InstallationStatus *string `json:"InstallationStatus,omitempty" xml:"InstallationStatus,omitempty"`
	// The value of MD5.
	//
	// example:
	//
	// THCIEH73KEK3334
	MD5 *string `json:"MD5,omitempty" xml:"MD5,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The status of the application.
	//
	// Valid values:
	//
	// 	- FAILED: The application failed to be created.
	//
	// 	- NORMAL: The application is available.
	//
	// 	- CREATING: The application is being created.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsRequest) GetAppIdList() []*string {
	return s.AppIdList
}

func (s *DescribeAppsRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppsRequest) GetAppType() *string {
	return s.AppType
}

func (s *DescribeAppsRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeAppsRequest) GetInstallationStatus() *string {
	return s.InstallationStatus
}

func (s *DescribeAppsRequest) GetMD5() *string {
	return s.MD5
}

func (s *DescribeAppsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAppsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAppsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppsRequest) SetAppIdList(v []*string) *DescribeAppsRequest {
	s.AppIdList = v
	return s
}

func (s *DescribeAppsRequest) SetAppName(v string) *DescribeAppsRequest {
	s.AppName = &v
	return s
}

func (s *DescribeAppsRequest) SetAppType(v string) *DescribeAppsRequest {
	s.AppType = &v
	return s
}

func (s *DescribeAppsRequest) SetBizRegionId(v string) *DescribeAppsRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeAppsRequest) SetInstallationStatus(v string) *DescribeAppsRequest {
	s.InstallationStatus = &v
	return s
}

func (s *DescribeAppsRequest) SetMD5(v string) *DescribeAppsRequest {
	s.MD5 = &v
	return s
}

func (s *DescribeAppsRequest) SetMaxResults(v int32) *DescribeAppsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAppsRequest) SetNextToken(v string) *DescribeAppsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAppsRequest) SetStatus(v string) *DescribeAppsRequest {
	s.Status = &v
	return s
}

func (s *DescribeAppsRequest) Validate() error {
	return dara.Validate(s)
}
