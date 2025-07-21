// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *GetDeviceConfigsRequest
	GetDeviceId() *string
	SetNetworkType(v string) *GetDeviceConfigsRequest
	GetNetworkType() *string
	SetRegion(v string) *GetDeviceConfigsRequest
	GetRegion() *string
	SetSerialNo(v string) *GetDeviceConfigsRequest
	GetSerialNo() *string
	SetUrclVersion(v string) *GetDeviceConfigsRequest
	GetUrclVersion() *string
	SetUserCustomId(v string) *GetDeviceConfigsRequest
	GetUserCustomId() *string
}

type GetDeviceConfigsRequest struct {
	DeviceId     *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	UrclVersion  *string `json:"UrclVersion,omitempty" xml:"UrclVersion,omitempty"`
	UserCustomId *string `json:"UserCustomId,omitempty" xml:"UserCustomId,omitempty"`
}

func (s GetDeviceConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceConfigsRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDeviceConfigsRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetDeviceConfigsRequest) GetRegion() *string {
	return s.Region
}

func (s *GetDeviceConfigsRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *GetDeviceConfigsRequest) GetUrclVersion() *string {
	return s.UrclVersion
}

func (s *GetDeviceConfigsRequest) GetUserCustomId() *string {
	return s.UserCustomId
}

func (s *GetDeviceConfigsRequest) SetDeviceId(v string) *GetDeviceConfigsRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceConfigsRequest) SetNetworkType(v string) *GetDeviceConfigsRequest {
	s.NetworkType = &v
	return s
}

func (s *GetDeviceConfigsRequest) SetRegion(v string) *GetDeviceConfigsRequest {
	s.Region = &v
	return s
}

func (s *GetDeviceConfigsRequest) SetSerialNo(v string) *GetDeviceConfigsRequest {
	s.SerialNo = &v
	return s
}

func (s *GetDeviceConfigsRequest) SetUrclVersion(v string) *GetDeviceConfigsRequest {
	s.UrclVersion = &v
	return s
}

func (s *GetDeviceConfigsRequest) SetUserCustomId(v string) *GetDeviceConfigsRequest {
	s.UserCustomId = &v
	return s
}

func (s *GetDeviceConfigsRequest) Validate() error {
	return dara.Validate(s)
}
