// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceAllDeviceGroupResponseBody
	GetCode() *string
	SetData(v *QueryFaceAllDeviceGroupResponseBodyData) *QueryFaceAllDeviceGroupResponseBody
	GetData() *QueryFaceAllDeviceGroupResponseBodyData
	SetErrorMessage(v string) *QueryFaceAllDeviceGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryFaceAllDeviceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFaceAllDeviceGroupResponseBody
	GetSuccess() *bool
}

type QueryFaceAllDeviceGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned by the API operation.
	Data *QueryFaceAllDeviceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s QueryFaceAllDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceAllDeviceGroupResponseBody) GetData() *QueryFaceAllDeviceGroupResponseBodyData {
	return s.Data
}

func (s *QueryFaceAllDeviceGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceAllDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceAllDeviceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetCode(v string) *QueryFaceAllDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetData(v *QueryFaceAllDeviceGroupResponseBodyData) *QueryFaceAllDeviceGroupResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetErrorMessage(v string) *QueryFaceAllDeviceGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetRequestId(v string) *QueryFaceAllDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBody) SetSuccess(v bool) *QueryFaceAllDeviceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceAllDeviceGroupResponseBodyData struct {
	// The list of device groups.
	DeviceGroupList []*QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of device groups.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceAllDeviceGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllDeviceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) GetDeviceGroupList() []*QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList {
	return s.DeviceGroupList
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) SetDeviceGroupList(v []*QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) *QueryFaceAllDeviceGroupResponseBodyData {
	s.DeviceGroupList = v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) SetPageNo(v int32) *QueryFaceAllDeviceGroupResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) SetPageSize(v int32) *QueryFaceAllDeviceGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) SetTotal(v int32) *QueryFaceAllDeviceGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyData) Validate() error {
	if s.DeviceGroupList != nil {
		for _, item := range s.DeviceGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList struct {
	// The ID of the device group.
	//
	// example:
	//
	// wy250q0m
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// The name of the device group.
	//
	// example:
	//
	// TestDeviceGroupName
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	// The time when the device group was modified.
	//
	// example:
	//
	// 2020-05-12 17:19:29
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) GetDeviceGroupName() *string {
	return s.DeviceGroupName
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) SetDeviceGroupId(v string) *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList {
	s.DeviceGroupId = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) SetDeviceGroupName(v string) *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList {
	s.DeviceGroupName = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) SetModifiedTime(v string) *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponseBodyDataDeviceGroupList) Validate() error {
	return dara.Validate(s)
}
