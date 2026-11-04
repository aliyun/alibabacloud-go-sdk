// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGbSubDeviceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryGbSubDeviceListRequest
	GetDeviceName() *string
	SetIotId(v string) *QueryGbSubDeviceListRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryGbSubDeviceListRequest
	GetIotInstanceId() *string
	SetPageSize(v int32) *QueryGbSubDeviceListRequest
	GetPageSize() *int32
	SetPageStart(v int32) *QueryGbSubDeviceListRequest
	GetPageStart() *int32
	SetProductKey(v string) *QueryGbSubDeviceListRequest
	GetProductKey() *string
}

type QueryGbSubDeviceListRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// This parameter is required.
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageStart     *int32  `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryGbSubDeviceListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGbSubDeviceListRequest) GoString() string {
	return s.String()
}

func (s *QueryGbSubDeviceListRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryGbSubDeviceListRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryGbSubDeviceListRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryGbSubDeviceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryGbSubDeviceListRequest) GetPageStart() *int32 {
	return s.PageStart
}

func (s *QueryGbSubDeviceListRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryGbSubDeviceListRequest) SetDeviceName(v string) *QueryGbSubDeviceListRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) SetIotId(v string) *QueryGbSubDeviceListRequest {
	s.IotId = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) SetIotInstanceId(v string) *QueryGbSubDeviceListRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) SetPageSize(v int32) *QueryGbSubDeviceListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) SetPageStart(v int32) *QueryGbSubDeviceListRequest {
	s.PageStart = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) SetProductKey(v string) *QueryGbSubDeviceListRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryGbSubDeviceListRequest) Validate() error {
	return dara.Validate(s)
}
