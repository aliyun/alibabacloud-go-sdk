// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAppInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupModels(v []*ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) *ListPublishedAppInstanceGroupResponseBody
	GetAppInstanceGroupModels() []*ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels
	SetPageNumber(v int32) *ListPublishedAppInstanceGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPublishedAppInstanceGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListPublishedAppInstanceGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPublishedAppInstanceGroupResponseBody
	GetTotalCount() *int32
}

type ListPublishedAppInstanceGroupResponseBody struct {
	// The list of published delivery groups on the current page, sorted by creation time from newest to oldest. An empty list is returned if no results match or if the requested page exceeds the result range.
	AppInstanceGroupModels []*ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels `json:"AppInstanceGroupModels,omitempty" xml:"AppInstanceGroupModels,omitempty" type:"Repeated"`
	// The page number specified in this request.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of delivery groups per page specified in the request. Unit: delivery groups. This value does not represent the actual number of delivery groups returned on the current page. The actual number on the current page may be less than this value.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID. You can use this ID to locate and troubleshoot issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of delivery groups that match all filter conditions. Unit: delivery groups. This value is not the number of delivery groups returned on the current page. Each delivery group is counted only once. The value is `0` if no results match.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublishedAppInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInstanceGroupResponseBody) GetAppInstanceGroupModels() []*ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels {
	return s.AppInstanceGroupModels
}

func (s *ListPublishedAppInstanceGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPublishedAppInstanceGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishedAppInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublishedAppInstanceGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPublishedAppInstanceGroupResponseBody) SetAppInstanceGroupModels(v []*ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) *ListPublishedAppInstanceGroupResponseBody {
	s.AppInstanceGroupModels = v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBody) SetPageNumber(v int32) *ListPublishedAppInstanceGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBody) SetPageSize(v int32) *ListPublishedAppInstanceGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBody) SetRequestId(v string) *ListPublishedAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBody) SetTotalCount(v int32) *ListPublishedAppInstanceGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBody) Validate() error {
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

type ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels struct {
	// The application image ID.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The delivery group ID.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name.
	//
	// example:
	//
	// OfficeApps
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The delivery group specification type.
	//
	// example:
	//
	// __dynamic__
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// The list of deployed applications in the delivery group image. The `AppId` and `AppName` parameters in the request only determine whether a delivery group is included in the results. They do not trim this list to only the matched applications.
	Apps []*ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The expiration time of the delivery group. For delivery groups sold as resources, this is the resource expiration time. For other delivery groups, this is the delivery group expiration time. The value is in ISO 8601 format with milliseconds and a time zone offset. The returned time zone offset is +00:00. Format: yyyy-MM-ddTHH:mm:ss.SSS+HH:mm.
	//
	// example:
	//
	// 2022-04-27T16:00:00.000+00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The creation time of the delivery group. The value is in ISO 8601 format with milliseconds and a time zone offset. The returned time zone offset is +00:00. Format: yyyy-MM-ddTHH:mm:ss.SSS+HH:mm.
	//
	// example:
	//
	// 2022-04-26T15:06:16.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The region ID of the delivery group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The delivery group status. This operation returns delivery groups only in the following status:
	//
	// - `PUBLISHED`: Published and listed. The delivery group can appear in the query results of this operation. This status does not indicate that a specified user has been granted access permissions.
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppCenterImageId() *string {
	return s.AppCenterImageId
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) GetAppInstanceType() *string {
	return s.AppInstanceType
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) GetApps() []*ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	return s.Apps
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) GetStatus() *string {
	return s.Status
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageId(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageId = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupId(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupName(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceType(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceType = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) SetApps(v []*ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Apps = v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) SetExpiredTime(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ExpiredTime = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) SetGmtCreate(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.GmtCreate = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) SetRegionId(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.RegionId = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) SetStatus(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Status = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModels) Validate() error {
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

type ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps struct {
	// The application icon.
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
	// The application name.
	//
	// example:
	//
	// OfficeApps
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application version.
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

func (s ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppIcon() *string {
	return s.AppIcon
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppId() *string {
	return s.AppId
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppName() *string {
	return s.AppName
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GetAppVersionName() *string {
	return s.AppVersionName
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppIcon(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppIcon = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppId(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppId = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppName(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppName = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersion(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersion = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersionName(v string) *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersionName = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) Validate() error {
	return dara.Validate(s)
}
