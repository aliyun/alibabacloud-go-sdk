// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAppsByUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApps(v []*ListAuthorizedAppsByUserResponseBodyApps) *ListAuthorizedAppsByUserResponseBody
	GetApps() []*ListAuthorizedAppsByUserResponseBodyApps
	SetPageNumber(v int32) *ListAuthorizedAppsByUserResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedAppsByUserResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAuthorizedAppsByUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAuthorizedAppsByUserResponseBody
	GetTotalCount() *int32
}

type ListAuthorizedAppsByUserResponseBody struct {
	// The list of applications authorized to the user. Each record corresponds to one application within one delivery group.
	Apps []*ListAuthorizedAppsByUserResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The page number of the returned results.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of authorization records that match the query conditions.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthorizedAppsByUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAppsByUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAppsByUserResponseBody) GetApps() []*ListAuthorizedAppsByUserResponseBodyApps {
	return s.Apps
}

func (s *ListAuthorizedAppsByUserResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedAppsByUserResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedAppsByUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedAppsByUserResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAuthorizedAppsByUserResponseBody) SetApps(v []*ListAuthorizedAppsByUserResponseBodyApps) *ListAuthorizedAppsByUserResponseBody {
	s.Apps = v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBody) SetPageNumber(v int32) *ListAuthorizedAppsByUserResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBody) SetPageSize(v int32) *ListAuthorizedAppsByUserResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBody) SetRequestId(v string) *ListAuthorizedAppsByUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBody) SetTotalCount(v int32) *ListAuthorizedAppsByUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBody) Validate() error {
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

type ListAuthorizedAppsByUserResponseBodyApps struct {
	// The URL of the application icon. This value is empty if the deployment details of the application cannot be obtained.
	//
	// example:
	//
	// https://app-center-icon-****.png
	AppIcon *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	// The application ID.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the delivery group to which the application belongs.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery group to which the application belongs. If the delivery group information cannot be obtained, the value is the same as AppInstanceGroupId.
	//
	// example:
	//
	// OfficeApp
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The application name. If the deployment details of the application cannot be obtained, the value is the same as AppId.
	//
	// example:
	//
	// OfficeApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application version. This value is empty if the deployment details of the application cannot be obtained.
	//
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The application version name. This value is empty if the deployment details of the application cannot be obtained.
	//
	// example:
	//
	// InitialVersion
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	// The number of authorized users for the application within its delivery group. This value is empty if the deployment details of the application cannot be obtained.
	//
	// example:
	//
	// 3
	AuthorizedUserCount *int32 `json:"AuthorizedUserCount,omitempty" xml:"AuthorizedUserCount,omitempty"`
}

func (s ListAuthorizedAppsByUserResponseBodyApps) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAppsByUserResponseBodyApps) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) GetAppIcon() *string {
	return s.AppIcon
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) GetAppId() *string {
	return s.AppId
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) GetAppName() *string {
	return s.AppName
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) GetAuthorizedUserCount() *int32 {
	return s.AuthorizedUserCount
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) SetAppIcon(v string) *ListAuthorizedAppsByUserResponseBodyApps {
	s.AppIcon = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) SetAppId(v string) *ListAuthorizedAppsByUserResponseBodyApps {
	s.AppId = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) SetAppInstanceGroupId(v string) *ListAuthorizedAppsByUserResponseBodyApps {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) SetAppInstanceGroupName(v string) *ListAuthorizedAppsByUserResponseBodyApps {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) SetAppName(v string) *ListAuthorizedAppsByUserResponseBodyApps {
	s.AppName = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) SetAppVersion(v string) *ListAuthorizedAppsByUserResponseBodyApps {
	s.AppVersion = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) SetAppVersionName(v string) *ListAuthorizedAppsByUserResponseBodyApps {
	s.AppVersionName = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) SetAuthorizedUserCount(v int32) *ListAuthorizedAppsByUserResponseBodyApps {
	s.AuthorizedUserCount = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponseBodyApps) Validate() error {
	return dara.Validate(s)
}
