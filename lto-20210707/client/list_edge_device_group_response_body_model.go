// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEdgeDeviceGroupResponseBody
	GetCode() *string
	SetData(v *ListEdgeDeviceGroupResponseBodyData) *ListEdgeDeviceGroupResponseBody
	GetData() *ListEdgeDeviceGroupResponseBodyData
	SetHttpStatusCode(v int32) *ListEdgeDeviceGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListEdgeDeviceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEdgeDeviceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEdgeDeviceGroupResponseBody
	GetSuccess() *bool
}

type ListEdgeDeviceGroupResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListEdgeDeviceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEdgeDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeDeviceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEdgeDeviceGroupResponseBody) GetData() *ListEdgeDeviceGroupResponseBodyData {
	return s.Data
}

func (s *ListEdgeDeviceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListEdgeDeviceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEdgeDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeDeviceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEdgeDeviceGroupResponseBody) SetCode(v string) *ListEdgeDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBody) SetData(v *ListEdgeDeviceGroupResponseBodyData) *ListEdgeDeviceGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListEdgeDeviceGroupResponseBody) SetHttpStatusCode(v int32) *ListEdgeDeviceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBody) SetMessage(v string) *ListEdgeDeviceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBody) SetRequestId(v string) *ListEdgeDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBody) SetSuccess(v bool) *ListEdgeDeviceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEdgeDeviceGroupResponseBodyData struct {
	Num      *int32                                         `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListEdgeDeviceGroupResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                         `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEdgeDeviceGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeDeviceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEdgeDeviceGroupResponseBodyData) GetNum() *int32 {
	return s.Num
}

func (s *ListEdgeDeviceGroupResponseBodyData) GetPageData() []*ListEdgeDeviceGroupResponseBodyDataPageData {
	return s.PageData
}

func (s *ListEdgeDeviceGroupResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListEdgeDeviceGroupResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListEdgeDeviceGroupResponseBodyData) SetNum(v int32) *ListEdgeDeviceGroupResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyData) SetPageData(v []*ListEdgeDeviceGroupResponseBodyDataPageData) *ListEdgeDeviceGroupResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyData) SetSize(v int32) *ListEdgeDeviceGroupResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyData) SetTotal(v int32) *ListEdgeDeviceGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyData) Validate() error {
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

type ListEdgeDeviceGroupResponseBodyDataPageData struct {
	AuthorizedCount *int32  `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	CurrentUser     *bool   `json:"CurrentUser,omitempty" xml:"CurrentUser,omitempty"`
	DeviceCount     *int64  `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	DeviceGroupId   *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	EdgeName        *string `json:"EdgeName,omitempty" xml:"EdgeName,omitempty"`
	MemberName      *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductKey      *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListEdgeDeviceGroupResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeDeviceGroupResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) GetAuthorizedCount() *int32 {
	return s.AuthorizedCount
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) GetCurrentUser() *bool {
	return s.CurrentUser
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) GetDeviceCount() *int64 {
	return s.DeviceCount
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) GetEdgeName() *string {
	return s.EdgeName
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) GetMemberName() *string {
	return s.MemberName
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) GetName() *string {
	return s.Name
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) GetProductKey() *string {
	return s.ProductKey
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) GetRemark() *string {
	return s.Remark
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) GetStatus() *string {
	return s.Status
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) SetAuthorizedCount(v int32) *ListEdgeDeviceGroupResponseBodyDataPageData {
	s.AuthorizedCount = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) SetCurrentUser(v bool) *ListEdgeDeviceGroupResponseBodyDataPageData {
	s.CurrentUser = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) SetDeviceCount(v int64) *ListEdgeDeviceGroupResponseBodyDataPageData {
	s.DeviceCount = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) SetDeviceGroupId(v string) *ListEdgeDeviceGroupResponseBodyDataPageData {
	s.DeviceGroupId = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) SetEdgeName(v string) *ListEdgeDeviceGroupResponseBodyDataPageData {
	s.EdgeName = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) SetMemberName(v string) *ListEdgeDeviceGroupResponseBodyDataPageData {
	s.MemberName = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) SetName(v string) *ListEdgeDeviceGroupResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) SetProductKey(v string) *ListEdgeDeviceGroupResponseBodyDataPageData {
	s.ProductKey = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) SetRemark(v string) *ListEdgeDeviceGroupResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) SetStatus(v string) *ListEdgeDeviceGroupResponseBodyDataPageData {
	s.Status = &v
	return s
}

func (s *ListEdgeDeviceGroupResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
