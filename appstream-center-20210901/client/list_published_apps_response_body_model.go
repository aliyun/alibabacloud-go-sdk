// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApps(v []*ListPublishedAppsResponseBodyApps) *ListPublishedAppsResponseBody
	GetApps() []*ListPublishedAppsResponseBodyApps
	SetPageNumber(v int32) *ListPublishedAppsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPublishedAppsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListPublishedAppsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPublishedAppsResponseBody
	GetTotalCount() *int32
}

type ListPublishedAppsResponseBody struct {
	// The list of application records on the current page. Each record corresponds to a deployed application in a published delivery group, sorted in descending order of the creation time of the associated delivery group. If no matching results are found or the requested page exceeds the result range, an empty list is returned.
	Apps []*ListPublishedAppsResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The page number specified in the current request.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page specified in the current request. This value does not represent the actual number of records returned on the current page. The actual number on the current page may be less than this value.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the current request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that meet all filter conditions. The count is based on the "delivery group + application" combination. If the same application appears in multiple published delivery groups, it is counted separately. This value is not the number of records returned on the current page. If no matching results are found, the value is `0`.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublishedAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedAppsResponseBody) GetApps() []*ListPublishedAppsResponseBodyApps {
	return s.Apps
}

func (s *ListPublishedAppsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPublishedAppsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishedAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublishedAppsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPublishedAppsResponseBody) SetApps(v []*ListPublishedAppsResponseBodyApps) *ListPublishedAppsResponseBody {
	s.Apps = v
	return s
}

func (s *ListPublishedAppsResponseBody) SetPageNumber(v int32) *ListPublishedAppsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPublishedAppsResponseBody) SetPageSize(v int32) *ListPublishedAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPublishedAppsResponseBody) SetRequestId(v string) *ListPublishedAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishedAppsResponseBody) SetTotalCount(v int32) *ListPublishedAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublishedAppsResponseBody) Validate() error {
	if s.Apps != nil {
		for _, item := range s.Apps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPublishedAppsResponseBodyApps struct {
	// The access URL of the application icon.
	//
	// example:
	//
	// https://app-center-icon-****.png
	AppIcon *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	// The application ID. When you invoke the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation to grant authorization to a user for this application, pass this value together with `AppInstanceGroupId`.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the published delivery group to which the application belongs.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the published delivery group to which the application belongs.
	//
	// example:
	//
	// Office applications
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The application name.
	//
	// example:
	//
	// Office applications
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application version number.
	//
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The application version name, which is the name set for this version when the application is published.
	//
	// example:
	//
	// Initial version
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	// The number of users currently authorized for this application on a per-application basis within the associated delivery group. Only users authorized on a per-application basis through [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) are counted. If no users are authorized on a per-application basis, `0` is returned.
	//
	// example:
	//
	// 5
	AuthorizedUserCount *int32 `json:"AuthorizedUserCount,omitempty" xml:"AuthorizedUserCount,omitempty"`
}

func (s ListPublishedAppsResponseBodyApps) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppsResponseBodyApps) GoString() string {
	return s.String()
}

func (s *ListPublishedAppsResponseBodyApps) GetAppIcon() *string {
	return s.AppIcon
}

func (s *ListPublishedAppsResponseBodyApps) GetAppId() *string {
	return s.AppId
}

func (s *ListPublishedAppsResponseBodyApps) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListPublishedAppsResponseBodyApps) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ListPublishedAppsResponseBodyApps) GetAppName() *string {
	return s.AppName
}

func (s *ListPublishedAppsResponseBodyApps) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListPublishedAppsResponseBodyApps) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *ListPublishedAppsResponseBodyApps) GetAuthorizedUserCount() *int32 {
	return s.AuthorizedUserCount
}

func (s *ListPublishedAppsResponseBodyApps) SetAppIcon(v string) *ListPublishedAppsResponseBodyApps {
	s.AppIcon = &v
	return s
}

func (s *ListPublishedAppsResponseBodyApps) SetAppId(v string) *ListPublishedAppsResponseBodyApps {
	s.AppId = &v
	return s
}

func (s *ListPublishedAppsResponseBodyApps) SetAppInstanceGroupId(v string) *ListPublishedAppsResponseBodyApps {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListPublishedAppsResponseBodyApps) SetAppInstanceGroupName(v string) *ListPublishedAppsResponseBodyApps {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListPublishedAppsResponseBodyApps) SetAppName(v string) *ListPublishedAppsResponseBodyApps {
	s.AppName = &v
	return s
}

func (s *ListPublishedAppsResponseBodyApps) SetAppVersion(v string) *ListPublishedAppsResponseBodyApps {
	s.AppVersion = &v
	return s
}

func (s *ListPublishedAppsResponseBodyApps) SetAppVersionName(v string) *ListPublishedAppsResponseBodyApps {
	s.AppVersionName = &v
	return s
}

func (s *ListPublishedAppsResponseBodyApps) SetAuthorizedUserCount(v int32) *ListPublishedAppsResponseBodyApps {
	s.AuthorizedUserCount = &v
	return s
}

func (s *ListPublishedAppsResponseBodyApps) Validate() error {
	return dara.Validate(s)
}
