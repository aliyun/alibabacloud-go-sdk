// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceDeviceGroupsByDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceDeviceGroupsByDeviceResponseBody
	GetCode() *string
	SetData(v *QueryFaceDeviceGroupsByDeviceResponseBodyData) *QueryFaceDeviceGroupsByDeviceResponseBody
	GetData() *QueryFaceDeviceGroupsByDeviceResponseBodyData
	SetErrorMessage(v string) *QueryFaceDeviceGroupsByDeviceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryFaceDeviceGroupsByDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFaceDeviceGroupsByDeviceResponseBody
	GetSuccess() *bool
}

type QueryFaceDeviceGroupsByDeviceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryFaceDeviceGroupsByDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) GetData() *QueryFaceDeviceGroupsByDeviceResponseBodyData {
	return s.Data
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetCode(v string) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetData(v *QueryFaceDeviceGroupsByDeviceResponseBodyData) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetErrorMessage(v string) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetRequestId(v string) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) SetSuccess(v bool) *QueryFaceDeviceGroupsByDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceDeviceGroupsByDeviceResponseBodyData struct {
	// The list of device groups.
	DeviceGroupList []*QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList `json:"DeviceGroupList,omitempty" xml:"DeviceGroupList,omitempty" type:"Repeated"`
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
	// 5
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) GetDeviceGroupList() []*QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList {
	return s.DeviceGroupList
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) SetDeviceGroupList(v []*QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) *QueryFaceDeviceGroupsByDeviceResponseBodyData {
	s.DeviceGroupList = v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) SetPageNo(v int32) *QueryFaceDeviceGroupsByDeviceResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) SetPageSize(v int32) *QueryFaceDeviceGroupsByDeviceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) SetTotal(v int32) *QueryFaceDeviceGroupsByDeviceResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyData) Validate() error {
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

type QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList struct {
	// The device group ID.
	//
	// example:
	//
	// TestDeviceGroupId
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// The device group name.
	//
	// example:
	//
	// TestDeviceGroupName
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	// The time when the device group was modified.
	//
	// example:
	//
	// 2020-05-12 17:12:37
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) GetDeviceGroupName() *string {
	return s.DeviceGroupName
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) SetDeviceGroupId(v string) *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList {
	s.DeviceGroupId = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) SetDeviceGroupName(v string) *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList {
	s.DeviceGroupName = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) SetModifiedTime(v string) *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList {
	s.ModifiedTime = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceResponseBodyDataDeviceGroupList) Validate() error {
	return dara.Validate(s)
}
