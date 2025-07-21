// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDevicesSecureNetworkTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllDevices(v int64) *ModifyDevicesSecureNetworkTypeRequest
	GetAllDevices() *int64
	SetSecureNetworkType(v string) *ModifyDevicesSecureNetworkTypeRequest
	GetSecureNetworkType() *string
	SetSerialNos(v string) *ModifyDevicesSecureNetworkTypeRequest
	GetSerialNos() *string
}

type ModifyDevicesSecureNetworkTypeRequest struct {
	AllDevices *int64 `json:"AllDevices,omitempty" xml:"AllDevices,omitempty"`
	// This parameter is required.
	SecureNetworkType *string `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	SerialNos         *string `json:"SerialNos,omitempty" xml:"SerialNos,omitempty"`
}

func (s ModifyDevicesSecureNetworkTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDevicesSecureNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDevicesSecureNetworkTypeRequest) GetAllDevices() *int64 {
	return s.AllDevices
}

func (s *ModifyDevicesSecureNetworkTypeRequest) GetSecureNetworkType() *string {
	return s.SecureNetworkType
}

func (s *ModifyDevicesSecureNetworkTypeRequest) GetSerialNos() *string {
	return s.SerialNos
}

func (s *ModifyDevicesSecureNetworkTypeRequest) SetAllDevices(v int64) *ModifyDevicesSecureNetworkTypeRequest {
	s.AllDevices = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeRequest) SetSecureNetworkType(v string) *ModifyDevicesSecureNetworkTypeRequest {
	s.SecureNetworkType = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeRequest) SetSerialNos(v string) *ModifyDevicesSecureNetworkTypeRequest {
	s.SerialNos = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeRequest) Validate() error {
	return dara.Validate(s)
}
