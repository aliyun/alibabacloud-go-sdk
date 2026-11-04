// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFaceUserDoExistOnDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *CheckFaceUserDoExistOnDeviceRequest
	GetDeviceName() *string
	SetIotInstanceId(v string) *CheckFaceUserDoExistOnDeviceRequest
	GetIotInstanceId() *string
	SetIsolationId(v string) *CheckFaceUserDoExistOnDeviceRequest
	GetIsolationId() *string
	SetProductKey(v string) *CheckFaceUserDoExistOnDeviceRequest
	GetProductKey() *string
	SetUserId(v string) *CheckFaceUserDoExistOnDeviceRequest
	GetUserId() *string
}

type CheckFaceUserDoExistOnDeviceRequest struct {
	// The device name.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestDeviceName
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The isolation ID.
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
	// a103lZ86ovStDMnU
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The ProductKey of the IoT Platform product.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1654Sr4Y3f
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tibce3fsgqelezk8
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CheckFaceUserDoExistOnDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckFaceUserDoExistOnDeviceRequest) GoString() string {
	return s.String()
}

func (s *CheckFaceUserDoExistOnDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CheckFaceUserDoExistOnDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *CheckFaceUserDoExistOnDeviceRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *CheckFaceUserDoExistOnDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *CheckFaceUserDoExistOnDeviceRequest) GetUserId() *string {
	return s.UserId
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetDeviceName(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetIotInstanceId(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetIsolationId(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.IsolationId = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetProductKey(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) SetUserId(v string) *CheckFaceUserDoExistOnDeviceRequest {
	s.UserId = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceRequest) Validate() error {
	return dara.Validate(s)
}
