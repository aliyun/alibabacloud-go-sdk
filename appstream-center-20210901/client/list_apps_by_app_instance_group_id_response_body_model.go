// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppsByAppInstanceGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApps(v []*ListAppsByAppInstanceGroupIdResponseBodyApps) *ListAppsByAppInstanceGroupIdResponseBody
	GetApps() []*ListAppsByAppInstanceGroupIdResponseBodyApps
	SetPageNumber(v int32) *ListAppsByAppInstanceGroupIdResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppsByAppInstanceGroupIdResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAppsByAppInstanceGroupIdResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAppsByAppInstanceGroupIdResponseBody
	GetTotalCount() *int32
}

type ListAppsByAppInstanceGroupIdResponseBody struct {
	// The list of application information on the current page. This is an empty list if no deployed applications exist in the delivery group image.
	Apps []*ListAppsByAppInstanceGroupIdResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The current page number, which is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of applications returned per page, which is the same as the PageSize request parameter.
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
	// The total number of deployed applications in the delivery group.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppsByAppInstanceGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppsByAppInstanceGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) GetApps() []*ListAppsByAppInstanceGroupIdResponseBodyApps {
	return s.Apps
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) SetApps(v []*ListAppsByAppInstanceGroupIdResponseBodyApps) *ListAppsByAppInstanceGroupIdResponseBody {
	s.Apps = v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) SetPageNumber(v int32) *ListAppsByAppInstanceGroupIdResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) SetPageSize(v int32) *ListAppsByAppInstanceGroupIdResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) SetRequestId(v string) *ListAppsByAppInstanceGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) SetTotalCount(v int32) *ListAppsByAppInstanceGroupIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBody) Validate() error {
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

type ListAppsByAppInstanceGroupIdResponseBodyApps struct {
	// The URL of the application icon.
	//
	// example:
	//
	// https://app-center-icon-****.png
	AppIcon *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	// The application ID. Pass in this value when you call the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation to authorize users for this application.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// OfficeApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application version number.
	//
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The application version name.
	//
	// example:
	//
	// InitialVersion
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	// The number of users currently authorized by application for this application in the delivery group. The value 0 is returned if no users are authorized by application.
	//
	// example:
	//
	// 5
	AuthorizedUserCount *int32 `json:"AuthorizedUserCount,omitempty" xml:"AuthorizedUserCount,omitempty"`
}

func (s ListAppsByAppInstanceGroupIdResponseBodyApps) String() string {
	return dara.Prettify(s)
}

func (s ListAppsByAppInstanceGroupIdResponseBodyApps) GoString() string {
	return s.String()
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) GetAppIcon() *string {
	return s.AppIcon
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) GetAppId() *string {
	return s.AppId
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) GetAppName() *string {
	return s.AppName
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) GetAuthorizedUserCount() *int32 {
	return s.AuthorizedUserCount
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) SetAppIcon(v string) *ListAppsByAppInstanceGroupIdResponseBodyApps {
	s.AppIcon = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) SetAppId(v string) *ListAppsByAppInstanceGroupIdResponseBodyApps {
	s.AppId = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) SetAppName(v string) *ListAppsByAppInstanceGroupIdResponseBodyApps {
	s.AppName = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) SetAppVersion(v string) *ListAppsByAppInstanceGroupIdResponseBodyApps {
	s.AppVersion = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) SetAppVersionName(v string) *ListAppsByAppInstanceGroupIdResponseBodyApps {
	s.AppVersionName = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) SetAuthorizedUserCount(v int32) *ListAppsByAppInstanceGroupIdResponseBodyApps {
	s.AuthorizedUserCount = &v
	return s
}

func (s *ListAppsByAppInstanceGroupIdResponseBodyApps) Validate() error {
	return dara.Validate(s)
}
