// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateDeviceSeatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *AddOrUpdateDeviceSeatsRequest
	GetFileName() *string
	SetUserCustomId(v string) *AddOrUpdateDeviceSeatsRequest
	GetUserCustomId() *string
	SetZoneId(v string) *AddOrUpdateDeviceSeatsRequest
	GetZoneId() *string
}

type AddOrUpdateDeviceSeatsRequest struct {
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	UserCustomId *string `json:"UserCustomId,omitempty" xml:"UserCustomId,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddOrUpdateDeviceSeatsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateDeviceSeatsRequest) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDeviceSeatsRequest) GetFileName() *string {
	return s.FileName
}

func (s *AddOrUpdateDeviceSeatsRequest) GetUserCustomId() *string {
	return s.UserCustomId
}

func (s *AddOrUpdateDeviceSeatsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddOrUpdateDeviceSeatsRequest) SetFileName(v string) *AddOrUpdateDeviceSeatsRequest {
	s.FileName = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsRequest) SetUserCustomId(v string) *AddOrUpdateDeviceSeatsRequest {
	s.UserCustomId = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsRequest) SetZoneId(v string) *AddOrUpdateDeviceSeatsRequest {
	s.ZoneId = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsRequest) Validate() error {
	return dara.Validate(s)
}
