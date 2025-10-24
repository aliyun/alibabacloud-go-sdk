// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceRegisterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeviceRegisterRequest
	GetAppId() *string
	SetNonce(v string) *DeviceRegisterRequest
	GetNonce() *string
	SetRequestTime(v string) *DeviceRegisterRequest
	GetRequestTime() *string
	SetSignature(v string) *DeviceRegisterRequest
	GetSignature() *string
}

type DeviceRegisterRequest struct {
	// This parameter is required.
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2a64edd96296880f55aa61987b
	Nonce *string `json:"nonce,omitempty" xml:"nonce,omitempty"`
	// This parameter is required.
	RequestTime *string `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// This parameter is required.
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s DeviceRegisterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeviceRegisterRequest) GoString() string {
	return s.String()
}

func (s *DeviceRegisterRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeviceRegisterRequest) GetNonce() *string {
	return s.Nonce
}

func (s *DeviceRegisterRequest) GetRequestTime() *string {
	return s.RequestTime
}

func (s *DeviceRegisterRequest) GetSignature() *string {
	return s.Signature
}

func (s *DeviceRegisterRequest) SetAppId(v string) *DeviceRegisterRequest {
	s.AppId = &v
	return s
}

func (s *DeviceRegisterRequest) SetNonce(v string) *DeviceRegisterRequest {
	s.Nonce = &v
	return s
}

func (s *DeviceRegisterRequest) SetRequestTime(v string) *DeviceRegisterRequest {
	s.RequestTime = &v
	return s
}

func (s *DeviceRegisterRequest) SetSignature(v string) *DeviceRegisterRequest {
	s.Signature = &v
	return s
}

func (s *DeviceRegisterRequest) Validate() error {
	return dara.Validate(s)
}
