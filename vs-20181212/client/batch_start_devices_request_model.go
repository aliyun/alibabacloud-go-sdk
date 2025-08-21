// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *BatchStartDevicesRequest
	GetId() *string
	SetOwnerId(v int64) *BatchStartDevicesRequest
	GetOwnerId() *int64
}

type BatchStartDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 32388487****92996-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchStartDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStartDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchStartDevicesRequest) GetId() *string {
	return s.Id
}

func (s *BatchStartDevicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchStartDevicesRequest) SetId(v string) *BatchStartDevicesRequest {
	s.Id = &v
	return s
}

func (s *BatchStartDevicesRequest) SetOwnerId(v int64) *BatchStartDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStartDevicesRequest) Validate() error {
	return dara.Validate(s)
}
