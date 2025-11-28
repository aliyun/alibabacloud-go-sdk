// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDeviceResponseBody
	GetCode() *string
	SetData(v *ListDeviceResponseBodyData) *ListDeviceResponseBody
	GetData() *ListDeviceResponseBodyData
	SetHttpStatusCode(v int32) *ListDeviceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDeviceResponseBody
	GetSuccess() *bool
}

type ListDeviceResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDeviceResponseBody) GetData() *ListDeviceResponseBodyData {
	return s.Data
}

func (s *ListDeviceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDeviceResponseBody) SetCode(v string) *ListDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceResponseBody) SetData(v *ListDeviceResponseBodyData) *ListDeviceResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceResponseBody) SetHttpStatusCode(v int32) *ListDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDeviceResponseBody) SetMessage(v string) *ListDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceResponseBody) SetRequestId(v string) *ListDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceResponseBody) SetSuccess(v bool) *ListDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *ListDeviceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeviceResponseBodyData struct {
	Num      *int32                                `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListDeviceResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeviceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceResponseBodyData) GetNum() *int32 {
	return s.Num
}

func (s *ListDeviceResponseBodyData) GetPageData() []*ListDeviceResponseBodyDataPageData {
	return s.PageData
}

func (s *ListDeviceResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListDeviceResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListDeviceResponseBodyData) SetNum(v int32) *ListDeviceResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListDeviceResponseBodyData) SetPageData(v []*ListDeviceResponseBodyDataPageData) *ListDeviceResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListDeviceResponseBodyData) SetSize(v int32) *ListDeviceResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListDeviceResponseBodyData) SetTotal(v int32) *ListDeviceResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListDeviceResponseBodyData) Validate() error {
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

type ListDeviceResponseBodyDataPageData struct {
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	LastOnchainTime  *string `json:"LastOnchainTime,omitempty" xml:"LastOnchainTime,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UsedOnchainCount *int64  `json:"UsedOnchainCount,omitempty" xml:"UsedOnchainCount,omitempty"`
}

func (s ListDeviceResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListDeviceResponseBodyDataPageData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ListDeviceResponseBodyDataPageData) GetLastOnchainTime() *string {
	return s.LastOnchainTime
}

func (s *ListDeviceResponseBodyDataPageData) GetName() *string {
	return s.Name
}

func (s *ListDeviceResponseBodyDataPageData) GetStatus() *string {
	return s.Status
}

func (s *ListDeviceResponseBodyDataPageData) GetUsedOnchainCount() *int64 {
	return s.UsedOnchainCount
}

func (s *ListDeviceResponseBodyDataPageData) SetDeviceId(v string) *ListDeviceResponseBodyDataPageData {
	s.DeviceId = &v
	return s
}

func (s *ListDeviceResponseBodyDataPageData) SetLastOnchainTime(v string) *ListDeviceResponseBodyDataPageData {
	s.LastOnchainTime = &v
	return s
}

func (s *ListDeviceResponseBodyDataPageData) SetName(v string) *ListDeviceResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListDeviceResponseBodyDataPageData) SetStatus(v string) *ListDeviceResponseBodyDataPageData {
	s.Status = &v
	return s
}

func (s *ListDeviceResponseBodyDataPageData) SetUsedOnchainCount(v int64) *ListDeviceResponseBodyDataPageData {
	s.UsedOnchainCount = &v
	return s
}

func (s *ListDeviceResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
