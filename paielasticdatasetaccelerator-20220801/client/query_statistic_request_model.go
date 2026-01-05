// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *QueryStatisticRequest
	GetEndTime() *string
	SetFields(v string) *QueryStatisticRequest
	GetFields() *string
	SetStartTime(v string) *QueryStatisticRequest
	GetStartTime() *string
}

type QueryStatisticRequest struct {
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// InstanceCapacityEachType
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// example:
	//
	// 2020-11-08T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryStatisticRequest) GoString() string {
	return s.String()
}

func (s *QueryStatisticRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryStatisticRequest) GetFields() *string {
	return s.Fields
}

func (s *QueryStatisticRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryStatisticRequest) SetEndTime(v string) *QueryStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *QueryStatisticRequest) SetFields(v string) *QueryStatisticRequest {
	s.Fields = &v
	return s
}

func (s *QueryStatisticRequest) SetStartTime(v string) *QueryStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *QueryStatisticRequest) Validate() error {
	return dara.Validate(s)
}
