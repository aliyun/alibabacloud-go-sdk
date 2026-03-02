// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetTraceDetailRequest
	GetEndTime() *string
	SetEnv(v string) *GetTraceDetailRequest
	GetEnv() *string
	SetStartTime(v string) *GetTraceDetailRequest
	GetStartTime() *string
}

type GetTraceDetailRequest struct {
	// example:
	//
	// 2022-10-27 15:05:48
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 2022-10-28 14:36:53
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetTraceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTraceDetailRequest) GoString() string {
	return s.String()
}

func (s *GetTraceDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTraceDetailRequest) GetEnv() *string {
	return s.Env
}

func (s *GetTraceDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTraceDetailRequest) SetEndTime(v string) *GetTraceDetailRequest {
	s.EndTime = &v
	return s
}

func (s *GetTraceDetailRequest) SetEnv(v string) *GetTraceDetailRequest {
	s.Env = &v
	return s
}

func (s *GetTraceDetailRequest) SetStartTime(v string) *GetTraceDetailRequest {
	s.StartTime = &v
	return s
}

func (s *GetTraceDetailRequest) Validate() error {
	return dara.Validate(s)
}
