// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetTraceRequest
	GetEndTime() *string
	SetStartTime(v string) *GetTraceRequest
	GetStartTime() *string
}

type GetTraceRequest struct {
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetTraceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTraceRequest) GoString() string {
	return s.String()
}

func (s *GetTraceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTraceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTraceRequest) SetEndTime(v string) *GetTraceRequest {
	s.EndTime = &v
	return s
}

func (s *GetTraceRequest) SetStartTime(v string) *GetTraceRequest {
	s.StartTime = &v
	return s
}

func (s *GetTraceRequest) Validate() error {
	return dara.Validate(s)
}
