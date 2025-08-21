// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *BatchDeleteDevicesRequest
	GetId() *string
	SetOwnerId(v int64) *BatchDeleteDevicesRequest
	GetOwnerId() *int64
}

type BatchDeleteDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 32388****39092996
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchDeleteDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDevicesRequest) GetId() *string {
	return s.Id
}

func (s *BatchDeleteDevicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchDeleteDevicesRequest) SetId(v string) *BatchDeleteDevicesRequest {
	s.Id = &v
	return s
}

func (s *BatchDeleteDevicesRequest) SetOwnerId(v int64) *BatchDeleteDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchDeleteDevicesRequest) Validate() error {
	return dara.Validate(s)
}
