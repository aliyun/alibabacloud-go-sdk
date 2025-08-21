// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StopStreamRequest
	GetId() *string
	SetName(v string) *StopStreamRequest
	GetName() *string
	SetOwnerId(v int64) *StopStreamRequest
	GetOwnerId() *int64
	SetStartTime(v string) *StopStreamRequest
	GetStartTime() *string
}

type StopStreamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 32388487****92997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 31000000*****0000002
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-12-12T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s StopStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s StopStreamRequest) GoString() string {
	return s.String()
}

func (s *StopStreamRequest) GetId() *string {
	return s.Id
}

func (s *StopStreamRequest) GetName() *string {
	return s.Name
}

func (s *StopStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopStreamRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *StopStreamRequest) SetId(v string) *StopStreamRequest {
	s.Id = &v
	return s
}

func (s *StopStreamRequest) SetName(v string) *StopStreamRequest {
	s.Name = &v
	return s
}

func (s *StopStreamRequest) SetOwnerId(v int64) *StopStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StopStreamRequest) SetStartTime(v string) *StopStreamRequest {
	s.StartTime = &v
	return s
}

func (s *StopStreamRequest) Validate() error {
	return dara.Validate(s)
}
