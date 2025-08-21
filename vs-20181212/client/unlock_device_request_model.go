// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UnlockDeviceRequest
	GetId() *string
	SetOwnerId(v int64) *UnlockDeviceRequest
	GetOwnerId() *int64
}

type UnlockDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 323884****9092996
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UnlockDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnlockDeviceRequest) GetId() *string {
	return s.Id
}

func (s *UnlockDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnlockDeviceRequest) SetId(v string) *UnlockDeviceRequest {
	s.Id = &v
	return s
}

func (s *UnlockDeviceRequest) SetOwnerId(v int64) *UnlockDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *UnlockDeviceRequest) Validate() error {
	return dara.Validate(s)
}
