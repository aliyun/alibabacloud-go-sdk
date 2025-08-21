// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *BatchStopDevicesRequest
	GetId() *string
	SetOwnerId(v int64) *BatchStopDevicesRequest
	GetOwnerId() *int64
	SetStartTime(v string) *BatchStopDevicesRequest
	GetStartTime() *string
}

type BatchStopDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 32388487****92996
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-10-14T23:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s BatchStopDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStopDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchStopDevicesRequest) GetId() *string {
	return s.Id
}

func (s *BatchStopDevicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchStopDevicesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchStopDevicesRequest) SetId(v string) *BatchStopDevicesRequest {
	s.Id = &v
	return s
}

func (s *BatchStopDevicesRequest) SetOwnerId(v int64) *BatchStopDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopDevicesRequest) SetStartTime(v string) *BatchStopDevicesRequest {
	s.StartTime = &v
	return s
}

func (s *BatchStopDevicesRequest) Validate() error {
	return dara.Validate(s)
}
