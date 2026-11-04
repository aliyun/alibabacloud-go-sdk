// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllUserGroupAndDeviceGroupRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody
	GetCode() *string
	SetData(v *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody
	GetData() *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData
	SetErrorMessage(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody
	GetSuccess() *bool
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A13279A-5640-45E7-87AA-83750541AD0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) GetData() *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	return s.Data
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetData(v *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetErrorMessage(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData struct {
	// The list of face synchronization relationships between user groups and device groups.
	List []*QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) GetList() []*QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	return s.List
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetList(v []*QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.List = v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetPage(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetPageSize(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) SetTotal(v int32) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList struct {
	// The control ID of the synchronization relationship between the user group and the device group.
	//
	// example:
	//
	// ni1iyxqry0j0****
	ControlId *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	// The face synchronization relationship type, which synchronizes user photos and user information.
	//
	// example:
	//
	// SYNC_INFO_PICTURE
	ControlType *string `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	// The device group ID.
	//
	// example:
	//
	// TestDeviceGro****
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2020-05-12 17:12:37
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The user group ID.
	//
	// example:
	//
	// TestUserGro****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) GetControlId() *string {
	return s.ControlId
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) GetControlType() *string {
	return s.ControlType
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetControlId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.ControlId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetControlType(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.ControlType = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetDeviceGroupId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.DeviceGroupId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetModifiedTime(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) SetUserGroupId(v string) *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList {
	s.UserGroupId = &v
	return s
}

func (s *QueryFaceAllUserGroupAndDeviceGroupRelationResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
