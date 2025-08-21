// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *StartStreamRequest
	GetEndTime() *int64
	SetId(v string) *StartStreamRequest
	GetId() *string
	SetOwnerId(v int64) *StartStreamRequest
	GetOwnerId() *int64
	SetStartTime(v int64) *StartStreamRequest
	GetStartTime() *int64
}

type StartStreamRequest struct {
	// example:
	//
	// 1599336385
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 323*****997-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1589336385
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s StartStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s StartStreamRequest) GoString() string {
	return s.String()
}

func (s *StartStreamRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *StartStreamRequest) GetId() *string {
	return s.Id
}

func (s *StartStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartStreamRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *StartStreamRequest) SetEndTime(v int64) *StartStreamRequest {
	s.EndTime = &v
	return s
}

func (s *StartStreamRequest) SetId(v string) *StartStreamRequest {
	s.Id = &v
	return s
}

func (s *StartStreamRequest) SetOwnerId(v int64) *StartStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StartStreamRequest) SetStartTime(v int64) *StartStreamRequest {
	s.StartTime = &v
	return s
}

func (s *StartStreamRequest) Validate() error {
	return dara.Validate(s)
}
