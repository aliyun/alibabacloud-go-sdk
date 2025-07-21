// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceFromSNRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *AddDeviceFromSNRequest
	GetAlias() *string
	SetCustomProperty(v string) *AddDeviceFromSNRequest
	GetCustomProperty() *string
	SetGroupId(v string) *AddDeviceFromSNRequest
	GetGroupId() *string
	SetLabelContents(v string) *AddDeviceFromSNRequest
	GetLabelContents() *string
	SetSecureNetworkType(v string) *AddDeviceFromSNRequest
	GetSecureNetworkType() *string
	SetSerialNo(v string) *AddDeviceFromSNRequest
	GetSerialNo() *string
}

type AddDeviceFromSNRequest struct {
	Alias             *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	CustomProperty    *string `json:"CustomProperty,omitempty" xml:"CustomProperty,omitempty"`
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	LabelContents     *string `json:"LabelContents,omitempty" xml:"LabelContents,omitempty"`
	SecureNetworkType *string `json:"SecureNetworkType,omitempty" xml:"SecureNetworkType,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s AddDeviceFromSNRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceFromSNRequest) GoString() string {
	return s.String()
}

func (s *AddDeviceFromSNRequest) GetAlias() *string {
	return s.Alias
}

func (s *AddDeviceFromSNRequest) GetCustomProperty() *string {
	return s.CustomProperty
}

func (s *AddDeviceFromSNRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddDeviceFromSNRequest) GetLabelContents() *string {
	return s.LabelContents
}

func (s *AddDeviceFromSNRequest) GetSecureNetworkType() *string {
	return s.SecureNetworkType
}

func (s *AddDeviceFromSNRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *AddDeviceFromSNRequest) SetAlias(v string) *AddDeviceFromSNRequest {
	s.Alias = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetCustomProperty(v string) *AddDeviceFromSNRequest {
	s.CustomProperty = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetGroupId(v string) *AddDeviceFromSNRequest {
	s.GroupId = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetLabelContents(v string) *AddDeviceFromSNRequest {
	s.LabelContents = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetSecureNetworkType(v string) *AddDeviceFromSNRequest {
	s.SecureNetworkType = &v
	return s
}

func (s *AddDeviceFromSNRequest) SetSerialNo(v string) *AddDeviceFromSNRequest {
	s.SerialNo = &v
	return s
}

func (s *AddDeviceFromSNRequest) Validate() error {
	return dara.Validate(s)
}
