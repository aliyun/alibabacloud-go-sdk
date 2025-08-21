// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StartDeviceRequest
	GetId() *string
	SetOwnerId(v int64) *StartDeviceRequest
	GetOwnerId() *int64
}

type StartDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 323884****9092996
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StartDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDeviceRequest) GoString() string {
	return s.String()
}

func (s *StartDeviceRequest) GetId() *string {
	return s.Id
}

func (s *StartDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartDeviceRequest) SetId(v string) *StartDeviceRequest {
	s.Id = &v
	return s
}

func (s *StartDeviceRequest) SetOwnerId(v int64) *StartDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *StartDeviceRequest) Validate() error {
	return dara.Validate(s)
}
