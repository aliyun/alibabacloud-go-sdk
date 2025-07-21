// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v string) *DeleteDevicesRequest
	GetForce() *string
	SetSerialNos(v string) *DeleteDevicesRequest
	GetSerialNos() *string
	SetUuids(v string) *DeleteDevicesRequest
	GetUuids() *string
}

type DeleteDevicesRequest struct {
	// This parameter is required.
	Force     *string `json:"Force,omitempty" xml:"Force,omitempty"`
	SerialNos *string `json:"SerialNos,omitempty" xml:"SerialNos,omitempty"`
	Uuids     *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DeleteDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDevicesRequest) GetForce() *string {
	return s.Force
}

func (s *DeleteDevicesRequest) GetSerialNos() *string {
	return s.SerialNos
}

func (s *DeleteDevicesRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DeleteDevicesRequest) SetForce(v string) *DeleteDevicesRequest {
	s.Force = &v
	return s
}

func (s *DeleteDevicesRequest) SetSerialNos(v string) *DeleteDevicesRequest {
	s.SerialNos = &v
	return s
}

func (s *DeleteDevicesRequest) SetUuids(v string) *DeleteDevicesRequest {
	s.Uuids = &v
	return s
}

func (s *DeleteDevicesRequest) Validate() error {
	return dara.Validate(s)
}
