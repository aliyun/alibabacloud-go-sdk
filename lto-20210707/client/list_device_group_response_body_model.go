// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDeviceGroupResponseBody
	GetCode() *string
	SetData(v *ListDeviceGroupResponseBodyData) *ListDeviceGroupResponseBody
	GetData() *ListDeviceGroupResponseBodyData
	SetHttpStatusCode(v int32) *ListDeviceGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDeviceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDeviceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDeviceGroupResponseBody
	GetSuccess() *bool
}

type ListDeviceGroupResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListDeviceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDeviceGroupResponseBody) GetData() *ListDeviceGroupResponseBodyData {
	return s.Data
}

func (s *ListDeviceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDeviceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDeviceGroupResponseBody) SetCode(v string) *ListDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceGroupResponseBody) SetData(v *ListDeviceGroupResponseBodyData) *ListDeviceGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceGroupResponseBody) SetHttpStatusCode(v int32) *ListDeviceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDeviceGroupResponseBody) SetMessage(v string) *ListDeviceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceGroupResponseBody) SetRequestId(v string) *ListDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceGroupResponseBody) SetSuccess(v bool) *ListDeviceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListDeviceGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeviceGroupResponseBodyData struct {
	Num      *int32                                     `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListDeviceGroupResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                     `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeviceGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupResponseBodyData) GetNum() *int32 {
	return s.Num
}

func (s *ListDeviceGroupResponseBodyData) GetPageData() []*ListDeviceGroupResponseBodyDataPageData {
	return s.PageData
}

func (s *ListDeviceGroupResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListDeviceGroupResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListDeviceGroupResponseBodyData) SetNum(v int32) *ListDeviceGroupResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListDeviceGroupResponseBodyData) SetPageData(v []*ListDeviceGroupResponseBodyDataPageData) *ListDeviceGroupResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListDeviceGroupResponseBodyData) SetSize(v int32) *ListDeviceGroupResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListDeviceGroupResponseBodyData) SetTotal(v int32) *ListDeviceGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListDeviceGroupResponseBodyData) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeviceGroupResponseBodyDataPageData struct {
	AuthorizedCount *int32  `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	CurrentUser     *bool   `json:"CurrentUser,omitempty" xml:"CurrentUser,omitempty"`
	DeviceCount     *int64  `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	DeviceGroupId   *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	MemberName      *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductKey      *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeviceGroupResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupResponseBodyDataPageData) GetAuthorizedCount() *int32 {
	return s.AuthorizedCount
}

func (s *ListDeviceGroupResponseBodyDataPageData) GetCurrentUser() *bool {
	return s.CurrentUser
}

func (s *ListDeviceGroupResponseBodyDataPageData) GetDeviceCount() *int64 {
	return s.DeviceCount
}

func (s *ListDeviceGroupResponseBodyDataPageData) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *ListDeviceGroupResponseBodyDataPageData) GetMemberName() *string {
	return s.MemberName
}

func (s *ListDeviceGroupResponseBodyDataPageData) GetName() *string {
	return s.Name
}

func (s *ListDeviceGroupResponseBodyDataPageData) GetProductKey() *string {
	return s.ProductKey
}

func (s *ListDeviceGroupResponseBodyDataPageData) GetRemark() *string {
	return s.Remark
}

func (s *ListDeviceGroupResponseBodyDataPageData) GetStatus() *string {
	return s.Status
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetAuthorizedCount(v int32) *ListDeviceGroupResponseBodyDataPageData {
	s.AuthorizedCount = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetCurrentUser(v bool) *ListDeviceGroupResponseBodyDataPageData {
	s.CurrentUser = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetDeviceCount(v int64) *ListDeviceGroupResponseBodyDataPageData {
	s.DeviceCount = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetDeviceGroupId(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.DeviceGroupId = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetMemberName(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.MemberName = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetName(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetProductKey(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.ProductKey = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetRemark(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) SetStatus(v string) *ListDeviceGroupResponseBodyDataPageData {
	s.Status = &v
	return s
}

func (s *ListDeviceGroupResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
