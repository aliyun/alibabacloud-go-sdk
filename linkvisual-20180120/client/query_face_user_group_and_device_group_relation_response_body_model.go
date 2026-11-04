// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserGroupAndDeviceGroupRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody
	GetCode() *string
	SetData(v *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody
	GetData() *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData
	SetErrorMessage(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody
	GetSuccess() *bool
}

type QueryFaceUserGroupAndDeviceGroupRelationResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value of the API call.
	Data *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the API call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) GetData() *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	return s.Data
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetData(v *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetErrorMessage(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *QueryFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData struct {
	// The control ID for the face synchronization relationship between user groups and device groups.
	//
	// example:
	//
	// ni1iyxqry0j03plv
	ControlId *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	// The face synchronization relationship.
	//
	// example:
	//
	// SYNC_INFO_PICTURE
	ControlType *string `json:"ControlType,omitempty" xml:"ControlType,omitempty"`
	// The device group ID.
	//
	// example:
	//
	// TestDeviceGroupId
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
	// TestUserGroupId
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) GetControlId() *string {
	return s.ControlId
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) GetControlType() *string {
	return s.ControlType
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlType(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlType = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetDeviceGroupId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.DeviceGroupId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetModifiedTime(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetUserGroupId(v string) *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
