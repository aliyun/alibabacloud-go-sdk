// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDeviceSeatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSerialNoList(v []*string) *UnbindDeviceSeatsRequest
	GetSerialNoList() []*string
}

type UnbindDeviceSeatsRequest struct {
	SerialNoList []*string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty" type:"Repeated"`
}

func (s UnbindDeviceSeatsRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceSeatsRequest) GoString() string {
	return s.String()
}

func (s *UnbindDeviceSeatsRequest) GetSerialNoList() []*string {
	return s.SerialNoList
}

func (s *UnbindDeviceSeatsRequest) SetSerialNoList(v []*string) *UnbindDeviceSeatsRequest {
	s.SerialNoList = v
	return s
}

func (s *UnbindDeviceSeatsRequest) Validate() error {
	return dara.Validate(s)
}
