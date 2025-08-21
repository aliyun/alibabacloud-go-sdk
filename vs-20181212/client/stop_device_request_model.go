// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StopDeviceRequest
	GetId() *string
	SetOwnerId(v int64) *StopDeviceRequest
	GetOwnerId() *int64
	SetStartTime(v string) *StopDeviceRequest
	GetStartTime() *string
}

type StopDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-12-10T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s StopDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDeviceRequest) GoString() string {
	return s.String()
}

func (s *StopDeviceRequest) GetId() *string {
	return s.Id
}

func (s *StopDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopDeviceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *StopDeviceRequest) SetId(v string) *StopDeviceRequest {
	s.Id = &v
	return s
}

func (s *StopDeviceRequest) SetOwnerId(v int64) *StopDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *StopDeviceRequest) SetStartTime(v string) *StopDeviceRequest {
	s.StartTime = &v
	return s
}

func (s *StopDeviceRequest) Validate() error {
	return dara.Validate(s)
}
