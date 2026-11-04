// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserGroupAndDeviceGroupRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody
	GetCode() *string
	SetData(v *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) *AddFaceUserGroupAndDeviceGroupRelationResponseBody
	GetData() *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData
	SetErrorMessage(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddFaceUserGroupAndDeviceGroupRelationResponseBody
	GetSuccess() *bool
}

type AddFaceUserGroupAndDeviceGroupRelationResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The face synchronization relationship information.
	Data *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) GetData() *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	return s.Data
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetCode(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetData(v *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Data = v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetErrorMessage(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetRequestId(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) SetSuccess(v bool) *AddFaceUserGroupAndDeviceGroupRelationResponseBody {
	s.Success = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFaceUserGroupAndDeviceGroupRelationResponseBodyData struct {
	// The associated control ID.
	//
	// example:
	//
	// qcb2yea4ha4dxmly
	ControlId *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2020-11-20 18:59:01
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) GetControlId() *string {
	return s.ControlId
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetControlId(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ControlId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) SetModifiedTime(v string) *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
