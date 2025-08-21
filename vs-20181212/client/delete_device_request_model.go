// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteDeviceRequest
	GetId() *string
	SetOwnerId(v int64) *DeleteDeviceRequest
	GetOwnerId() *int64
}

type DeleteDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3238848****092996
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceRequest) GetId() *string {
	return s.Id
}

func (s *DeleteDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDeviceRequest) SetId(v string) *DeleteDeviceRequest {
	s.Id = &v
	return s
}

func (s *DeleteDeviceRequest) SetOwnerId(v int64) *DeleteDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDeviceRequest) Validate() error {
	return dara.Validate(s)
}
