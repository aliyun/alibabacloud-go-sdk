// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopStreamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *BatchStopStreamsRequest
	GetId() *string
	SetOwnerId(v int64) *BatchStopStreamsRequest
	GetOwnerId() *int64
	SetStartTime(v string) *BatchStopStreamsRequest
	GetStartTime() *string
}

type BatchStopStreamsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 323*****997-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-12-10T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s BatchStopStreamsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStopStreamsRequest) GoString() string {
	return s.String()
}

func (s *BatchStopStreamsRequest) GetId() *string {
	return s.Id
}

func (s *BatchStopStreamsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchStopStreamsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchStopStreamsRequest) SetId(v string) *BatchStopStreamsRequest {
	s.Id = &v
	return s
}

func (s *BatchStopStreamsRequest) SetOwnerId(v int64) *BatchStopStreamsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopStreamsRequest) SetStartTime(v string) *BatchStopStreamsRequest {
	s.StartTime = &v
	return s
}

func (s *BatchStopStreamsRequest) Validate() error {
	return dara.Validate(s)
}
