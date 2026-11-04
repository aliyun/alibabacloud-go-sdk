// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryVisionDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceNameList(v []*string) *BatchQueryVisionDeviceInfoRequest
	GetDeviceNameList() []*string
	SetIotIdList(v []*string) *BatchQueryVisionDeviceInfoRequest
	GetIotIdList() []*string
	SetIotInstanceId(v string) *BatchQueryVisionDeviceInfoRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *BatchQueryVisionDeviceInfoRequest
	GetProductKey() *string
}

type BatchQueryVisionDeviceInfoRequest struct {
	DeviceNameList []*string `json:"DeviceNameList,omitempty" xml:"DeviceNameList,omitempty" type:"Repeated"`
	IotIdList      []*string `json:"IotIdList,omitempty" xml:"IotIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s BatchQueryVisionDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoRequest) GetDeviceNameList() []*string {
	return s.DeviceNameList
}

func (s *BatchQueryVisionDeviceInfoRequest) GetIotIdList() []*string {
	return s.IotIdList
}

func (s *BatchQueryVisionDeviceInfoRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *BatchQueryVisionDeviceInfoRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *BatchQueryVisionDeviceInfoRequest) SetDeviceNameList(v []*string) *BatchQueryVisionDeviceInfoRequest {
	s.DeviceNameList = v
	return s
}

func (s *BatchQueryVisionDeviceInfoRequest) SetIotIdList(v []*string) *BatchQueryVisionDeviceInfoRequest {
	s.IotIdList = v
	return s
}

func (s *BatchQueryVisionDeviceInfoRequest) SetIotInstanceId(v string) *BatchQueryVisionDeviceInfoRequest {
	s.IotInstanceId = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoRequest) SetProductKey(v string) *BatchQueryVisionDeviceInfoRequest {
	s.ProductKey = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
