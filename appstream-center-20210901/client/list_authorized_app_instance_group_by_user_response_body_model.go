// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAppInstanceGroupByUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupModels(v []*ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) *ListAuthorizedAppInstanceGroupByUserResponseBody
	GetAppInstanceGroupModels() []*ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels
	SetPageNumber(v int32) *ListAuthorizedAppInstanceGroupByUserResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedAppInstanceGroupByUserResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAuthorizedAppInstanceGroupByUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAuthorizedAppInstanceGroupByUserResponseBody
	GetTotalCount() *int32
}

type ListAuthorizedAppInstanceGroupByUserResponseBody struct {
	// The list of authorized delivery groups on the current page. This is an empty list if the user has no authorized delivery groups that match the conditions.
	AppInstanceGroupModels []*ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels `json:"AppInstanceGroupModels,omitempty" xml:"AppInstanceGroupModels,omitempty" type:"Repeated"`
	// The current page number. This value is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of delivery groups returned per page. This value is the same as the PageSize request parameter.
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
	// The total number of authorized delivery groups that match the filter conditions.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthorizedAppInstanceGroupByUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAppInstanceGroupByUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) GetAppInstanceGroupModels() []*ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels {
	return s.AppInstanceGroupModels
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) SetAppInstanceGroupModels(v []*ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) *ListAuthorizedAppInstanceGroupByUserResponseBody {
	s.AppInstanceGroupModels = v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) SetPageNumber(v int32) *ListAuthorizedAppInstanceGroupByUserResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) SetPageSize(v int32) *ListAuthorizedAppInstanceGroupByUserResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) SetRequestId(v string) *ListAuthorizedAppInstanceGroupByUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) SetTotalCount(v int32) *ListAuthorizedAppInstanceGroupByUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBody) Validate() error {
	if s.AppInstanceGroupModels != nil {
		for _, item := range s.AppInstanceGroupModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels struct {
	// The ID of the application image used by the delivery group.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The delivery group ID. You can pass this value to the [GetConnectionTicket](~~GetConnectionTicket~~) operation to specify the delivery group to connect to.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name.
	//
	// example:
	//
	// Office App
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The specification type of the delivery group.
	//
	// example:
	//
	// __dynamic__
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// The list of applications deployed in the delivery group. This list includes all deployed applications in the delivery group image and is not affected by the AppId or AppName request parameters.
	Apps []*ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The expiration time of the delivery group. The value is in the ISO 8601 datetime format with milliseconds and a time zone offset. The returned time zone offset is +00:00. Format: yyyy-MM-ddTHH:mm:ss.SSS+HH:mm.
	//
	// example:
	//
	// 2022-04-27T16:00:00.000+00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The creation time of the delivery group. The value is in the ISO 8601 datetime format with milliseconds and a time zone offset. The returned time zone offset is +00:00. Format: yyyy-MM-ddTHH:mm:ss.SSS+HH:mm.
	//
	// example:
	//
	// 2022-04-26T15:06:16.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The region ID of the delivery group. For more information about supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The delivery group status. Valid values:
	//
	// - DEPLOYING: The delivery group is being created or starting resources and is not yet connectable.
	//
	// - DEPLOYED: The delivery group is deployed. Resources are ready but the delivery group has not been listed for service.
	//
	// - PUBLISHED: The delivery group is published and listed. Users can connect to and use the delivery group.
	//
	// - STOPPING: The delivery group is being delisted and is stopping service.
	//
	// - STOPPED: The delivery group is delisted and has stopped service. Users cannot connect.
	//
	// - MAINTAINING: The delivery group is being maintained or updated.
	//
	// - FAILED: The delivery group failed to publish. Resource initialization failed.
	//
	// - MAINTAIN_FAILED: The update failed. Maintenance or changes were not successful.
	//
	// - DELETING: The delivery group is being deleted.
	//
	// > Deleted delivery groups are not returned.
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) GetAppCenterImageId() *string {
	return s.AppCenterImageId
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) GetAppInstanceType() *string {
	return s.AppInstanceType
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) GetApps() []*ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps {
	return s.Apps
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) GetStatus() *string {
	return s.Status
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) SetAppCenterImageId(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels {
	s.AppCenterImageId = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) SetAppInstanceGroupId(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) SetAppInstanceGroupName(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) SetAppInstanceType(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels {
	s.AppInstanceType = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) SetApps(v []*ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels {
	s.Apps = v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) SetExpiredTime(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels {
	s.ExpiredTime = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) SetGmtCreate(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels {
	s.GmtCreate = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) SetRegionId(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels {
	s.RegionId = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) SetStatus(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels {
	s.Status = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModels) Validate() error {
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

type ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps struct {
	// The URL of the application icon.
	//
	// example:
	//
	// https://app-center-icon-****.png
	AppIcon *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	// The application ID. Pass this value to the [GetConnectionTicket](~~GetConnectionTicket~~) operation to obtain a connection ticket for the application.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// Office App
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
	// Initial version
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
}

func (s ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) GetAppIcon() *string {
	return s.AppIcon
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) GetAppId() *string {
	return s.AppId
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) GetAppName() *string {
	return s.AppName
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) SetAppIcon(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps {
	s.AppIcon = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) SetAppId(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps {
	s.AppId = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) SetAppName(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps {
	s.AppName = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) SetAppVersion(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps {
	s.AppVersion = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) SetAppVersionName(v string) *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps {
	s.AppVersionName = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserResponseBodyAppInstanceGroupModelsApps) Validate() error {
	return dara.Validate(s)
}
