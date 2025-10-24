// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeAppsResponseBodyData) *DescribeAppsResponseBody
	GetData() []*DescribeAppsResponseBodyData
	SetNextToken(v string) *DescribeAppsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAppsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeAppsResponseBody
	GetTotalCount() *string
}

type DescribeAppsResponseBody struct {
	// The objects that are returned.
	Data []*DescribeAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CB95E410-FD1D-53C5-9F7D-93CC44D7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) GetData() []*DescribeAppsResponseBodyData {
	return s.Data
}

func (s *DescribeAppsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeAppsResponseBody) SetData(v []*DescribeAppsResponseBodyData) *DescribeAppsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAppsResponseBody) SetNextToken(v string) *DescribeAppsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetTotalCount(v string) *DescribeAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAppsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppsResponseBodyData struct {
	// The version of the application.
	//
	// example:
	//
	// 1.0.0
	AndroidAppVersion *string `json:"AndroidAppVersion,omitempty" xml:"AndroidAppVersion,omitempty"`
	// Apk size.
	//
	// example:
	//
	// 10244893
	ApkSize *string `json:"ApkSize,omitempty" xml:"ApkSize,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 10404
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// testapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// Region id.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// default description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2022-08-11 17:45:03
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the application was last modified.
	//
	// example:
	//
	// 2022-08-11 17:45:03
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The icon URL of the application.
	//
	// example:
	//
	// https://test.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
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
	// The list of instance groups where the application is installed.
	InstanceGroupList []*string `json:"InstanceGroupList,omitempty" xml:"InstanceGroupList,omitempty" type:"Repeated"`
	// The value of MD5.
	//
	// example:
	//
	// THCIEH73KEK3334
	MD5 *string `json:"MD5,omitempty" xml:"MD5,omitempty"`
	// The name of the application package.
	//
	// example:
	//
	// cn.rdstar.rdstarandroid
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
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

func (s DescribeAppsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyData) GetAndroidAppVersion() *string {
	return s.AndroidAppVersion
}

func (s *DescribeAppsResponseBodyData) GetApkSize() *string {
	return s.ApkSize
}

func (s *DescribeAppsResponseBodyData) GetAppId() *int32 {
	return s.AppId
}

func (s *DescribeAppsResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppsResponseBodyData) GetAppType() *string {
	return s.AppType
}

func (s *DescribeAppsResponseBodyData) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeAppsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppsResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeAppsResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAppsResponseBodyData) GetIconUrl() *string {
	return s.IconUrl
}

func (s *DescribeAppsResponseBodyData) GetInstallationStatus() *string {
	return s.InstallationStatus
}

func (s *DescribeAppsResponseBodyData) GetInstanceGroupList() []*string {
	return s.InstanceGroupList
}

func (s *DescribeAppsResponseBodyData) GetMD5() *string {
	return s.MD5
}

func (s *DescribeAppsResponseBodyData) GetPackageName() *string {
	return s.PackageName
}

func (s *DescribeAppsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppsResponseBodyData) SetAndroidAppVersion(v string) *DescribeAppsResponseBodyData {
	s.AndroidAppVersion = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetApkSize(v string) *DescribeAppsResponseBodyData {
	s.ApkSize = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetAppId(v int32) *DescribeAppsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetAppName(v string) *DescribeAppsResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetAppType(v string) *DescribeAppsResponseBodyData {
	s.AppType = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetBizRegionId(v string) *DescribeAppsResponseBodyData {
	s.BizRegionId = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetDescription(v string) *DescribeAppsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetGmtCreate(v string) *DescribeAppsResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetGmtModified(v string) *DescribeAppsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetIconUrl(v string) *DescribeAppsResponseBodyData {
	s.IconUrl = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetInstallationStatus(v string) *DescribeAppsResponseBodyData {
	s.InstallationStatus = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetInstanceGroupList(v []*string) *DescribeAppsResponseBodyData {
	s.InstanceGroupList = v
	return s
}

func (s *DescribeAppsResponseBodyData) SetMD5(v string) *DescribeAppsResponseBodyData {
	s.MD5 = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetPackageName(v string) *DescribeAppsResponseBodyData {
	s.PackageName = &v
	return s
}

func (s *DescribeAppsResponseBodyData) SetStatus(v string) *DescribeAppsResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeAppsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
