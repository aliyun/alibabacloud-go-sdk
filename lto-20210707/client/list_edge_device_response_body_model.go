// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEdgeDeviceResponseBody
	GetCode() *string
	SetData(v *ListEdgeDeviceResponseBodyData) *ListEdgeDeviceResponseBody
	GetData() *ListEdgeDeviceResponseBodyData
	SetHttpStatusCode(v int32) *ListEdgeDeviceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListEdgeDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEdgeDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEdgeDeviceResponseBody
	GetSuccess() *bool
}

type ListEdgeDeviceResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListEdgeDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEdgeDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEdgeDeviceResponseBody) GetData() *ListEdgeDeviceResponseBodyData {
	return s.Data
}

func (s *ListEdgeDeviceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListEdgeDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEdgeDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEdgeDeviceResponseBody) SetCode(v string) *ListEdgeDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ListEdgeDeviceResponseBody) SetData(v *ListEdgeDeviceResponseBodyData) *ListEdgeDeviceResponseBody {
	s.Data = v
	return s
}

func (s *ListEdgeDeviceResponseBody) SetHttpStatusCode(v int32) *ListEdgeDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEdgeDeviceResponseBody) SetMessage(v string) *ListEdgeDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ListEdgeDeviceResponseBody) SetRequestId(v string) *ListEdgeDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeDeviceResponseBody) SetSuccess(v bool) *ListEdgeDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *ListEdgeDeviceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEdgeDeviceResponseBodyData struct {
	Num      *int32                                    `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListEdgeDeviceResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                    `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEdgeDeviceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEdgeDeviceResponseBodyData) GetNum() *int32 {
	return s.Num
}

func (s *ListEdgeDeviceResponseBodyData) GetPageData() []*ListEdgeDeviceResponseBodyDataPageData {
	return s.PageData
}

func (s *ListEdgeDeviceResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListEdgeDeviceResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListEdgeDeviceResponseBodyData) SetNum(v int32) *ListEdgeDeviceResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListEdgeDeviceResponseBodyData) SetPageData(v []*ListEdgeDeviceResponseBodyDataPageData) *ListEdgeDeviceResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListEdgeDeviceResponseBodyData) SetSize(v int32) *ListEdgeDeviceResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListEdgeDeviceResponseBodyData) SetTotal(v int32) *ListEdgeDeviceResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListEdgeDeviceResponseBodyData) Validate() error {
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

type ListEdgeDeviceResponseBodyDataPageData struct {
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	LastOnchainTime  *string `json:"LastOnchainTime,omitempty" xml:"LastOnchainTime,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UsedOnchainCount *int64  `json:"UsedOnchainCount,omitempty" xml:"UsedOnchainCount,omitempty"`
}

func (s ListEdgeDeviceResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeDeviceResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListEdgeDeviceResponseBodyDataPageData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ListEdgeDeviceResponseBodyDataPageData) GetLastOnchainTime() *string {
	return s.LastOnchainTime
}

func (s *ListEdgeDeviceResponseBodyDataPageData) GetName() *string {
	return s.Name
}

func (s *ListEdgeDeviceResponseBodyDataPageData) GetStatus() *string {
	return s.Status
}

func (s *ListEdgeDeviceResponseBodyDataPageData) GetUsedOnchainCount() *int64 {
	return s.UsedOnchainCount
}

func (s *ListEdgeDeviceResponseBodyDataPageData) SetDeviceId(v string) *ListEdgeDeviceResponseBodyDataPageData {
	s.DeviceId = &v
	return s
}

func (s *ListEdgeDeviceResponseBodyDataPageData) SetLastOnchainTime(v string) *ListEdgeDeviceResponseBodyDataPageData {
	s.LastOnchainTime = &v
	return s
}

func (s *ListEdgeDeviceResponseBodyDataPageData) SetName(v string) *ListEdgeDeviceResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListEdgeDeviceResponseBodyDataPageData) SetStatus(v string) *ListEdgeDeviceResponseBodyDataPageData {
	s.Status = &v
	return s
}

func (s *ListEdgeDeviceResponseBodyDataPageData) SetUsedOnchainCount(v int64) *ListEdgeDeviceResponseBodyDataPageData {
	s.UsedOnchainCount = &v
	return s
}

func (s *ListEdgeDeviceResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
