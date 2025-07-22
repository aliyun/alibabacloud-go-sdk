// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetInstanceEventsRequest
	GetEndTime() *string
	SetMaxEventsNum(v int32) *GetInstanceEventsRequest
	GetMaxEventsNum() *int32
	SetStartTime(v string) *GetInstanceEventsRequest
	GetStartTime() *string
	SetToken(v string) *GetInstanceEventsRequest
	GetToken() *string
}

type GetInstanceEventsRequest struct {
	// The end of the time range to query.
	//
	// example:
	//
	// 2020-11-08T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of events. Default value: 2000.
	//
	// example:
	//
	// 2000
	MaxEventsNum *int32 `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2020-11-08T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The token used to share the URL.
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetInstanceEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceEventsRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceEventsRequest) GetMaxEventsNum() *int32 {
	return s.MaxEventsNum
}

func (s *GetInstanceEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceEventsRequest) GetToken() *string {
	return s.Token
}

func (s *GetInstanceEventsRequest) SetEndTime(v string) *GetInstanceEventsRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceEventsRequest) SetMaxEventsNum(v int32) *GetInstanceEventsRequest {
	s.MaxEventsNum = &v
	return s
}

func (s *GetInstanceEventsRequest) SetStartTime(v string) *GetInstanceEventsRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceEventsRequest) SetToken(v string) *GetInstanceEventsRequest {
	s.Token = &v
	return s
}

func (s *GetInstanceEventsRequest) Validate() error {
	return dara.Validate(s)
}
