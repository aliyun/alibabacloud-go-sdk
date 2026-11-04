// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceDeviceGroupsByDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryFaceDeviceGroupsByDeviceRequest
	GetDeviceName() *string
	SetIotInstanceId(v string) *QueryFaceDeviceGroupsByDeviceRequest
	GetIotInstanceId() *string
	SetIsolationId(v string) *QueryFaceDeviceGroupsByDeviceRequest
	GetIsolationId() *string
	SetPageNo(v int32) *QueryFaceDeviceGroupsByDeviceRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryFaceDeviceGroupsByDeviceRequest
	GetPageSize() *int32
	SetProductKey(v string) *QueryFaceDeviceGroupsByDeviceRequest
	GetProductKey() *string
}

type QueryFaceDeviceGroupsByDeviceRequest struct {
	// The device name.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestDeviceName
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIotInstanceId
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolationId
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product key.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestProductKey
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryFaceDeviceGroupsByDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceDeviceGroupsByDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetDeviceName(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetIotInstanceId(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetIsolationId(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetPageNo(v int32) *QueryFaceDeviceGroupsByDeviceRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetPageSize(v int32) *QueryFaceDeviceGroupsByDeviceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) SetProductKey(v string) *QueryFaceDeviceGroupsByDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryFaceDeviceGroupsByDeviceRequest) Validate() error {
	return dara.Validate(s)
}
