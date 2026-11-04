// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGbSubDeviceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryGbSubDeviceListResponseBody
	GetCode() *string
	SetData(v *QueryGbSubDeviceListResponseBodyData) *QueryGbSubDeviceListResponseBody
	GetData() *QueryGbSubDeviceListResponseBodyData
	SetErrorMessage(v string) *QueryGbSubDeviceListResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryGbSubDeviceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryGbSubDeviceListResponseBody
	GetSuccess() *bool
}

type QueryGbSubDeviceListResponseBody struct {
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryGbSubDeviceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryGbSubDeviceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGbSubDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGbSubDeviceListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryGbSubDeviceListResponseBody) GetData() *QueryGbSubDeviceListResponseBodyData {
	return s.Data
}

func (s *QueryGbSubDeviceListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryGbSubDeviceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGbSubDeviceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryGbSubDeviceListResponseBody) SetCode(v string) *QueryGbSubDeviceListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBody) SetData(v *QueryGbSubDeviceListResponseBodyData) *QueryGbSubDeviceListResponseBody {
	s.Data = v
	return s
}

func (s *QueryGbSubDeviceListResponseBody) SetErrorMessage(v string) *QueryGbSubDeviceListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBody) SetRequestId(v string) *QueryGbSubDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBody) SetSuccess(v bool) *QueryGbSubDeviceListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryGbSubDeviceListResponseBodyData struct {
	GbSubDeviceList []*QueryGbSubDeviceListResponseBodyDataGbSubDeviceList `json:"GbSubDeviceList,omitempty" xml:"GbSubDeviceList,omitempty" type:"Repeated"`
	Total           *int32                                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryGbSubDeviceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryGbSubDeviceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGbSubDeviceListResponseBodyData) GetGbSubDeviceList() []*QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	return s.GbSubDeviceList
}

func (s *QueryGbSubDeviceListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryGbSubDeviceListResponseBodyData) SetGbSubDeviceList(v []*QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) *QueryGbSubDeviceListResponseBodyData {
	s.GbSubDeviceList = v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyData) SetTotal(v int32) *QueryGbSubDeviceListResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyData) Validate() error {
	if s.GbSubDeviceList != nil {
		for _, item := range s.GbSubDeviceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryGbSubDeviceListResponseBodyDataGbSubDeviceList struct {
	DeviceEnable *int32  `json:"DeviceEnable,omitempty" xml:"DeviceEnable,omitempty"`
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceName   *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId        *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	ProductKey   *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) String() string {
	return dara.Prettify(s)
}

func (s QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) GoString() string {
	return s.String()
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) GetDeviceEnable() *int32 {
	return s.DeviceEnable
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) GetDeviceId() *string {
	return s.DeviceId
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) GetIotId() *string {
	return s.IotId
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) SetDeviceEnable(v int32) *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	s.DeviceEnable = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) SetDeviceId(v string) *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	s.DeviceId = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) SetDeviceName(v string) *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	s.DeviceName = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) SetIotId(v string) *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	s.IotId = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) SetProductKey(v string) *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList {
	s.ProductKey = &v
	return s
}

func (s *QueryGbSubDeviceListResponseBodyDataGbSubDeviceList) Validate() error {
	return dara.Validate(s)
}
