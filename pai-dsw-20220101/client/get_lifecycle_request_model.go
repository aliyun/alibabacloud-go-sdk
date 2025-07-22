// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLifecycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetLifecycleRequest
	GetEndTime() *string
	SetLimit(v int32) *GetLifecycleRequest
	GetLimit() *int32
	SetOrder(v string) *GetLifecycleRequest
	GetOrder() *string
	SetSessionNumber(v int32) *GetLifecycleRequest
	GetSessionNumber() *int32
	SetStartTime(v string) *GetLifecycleRequest
	GetStartTime() *string
	SetToken(v string) *GetLifecycleRequest
	GetToken() *string
}

type GetLifecycleRequest struct {
	// The end of the time range to query.
	//
	// example:
	//
	// 2020-11-08T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of sessions to query.
	//
	// example:
	//
	// 1
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The sorting order of the results. Valid values:
	//
	// 	- ASC: sorted by time in ascending order.
	//
	// 	- DESC: sorted by time in descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// A session refers to the process of an instance from startup to failure or shutdown. The sessionNumber indicates the offset value for the instance\\"s session sequence.
	//
	// example:
	//
	// 1
	SessionNumber *int32 `json:"SessionNumber,omitempty" xml:"SessionNumber,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2020-11-08T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The token used to share the URL.
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetLifecycleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLifecycleRequest) GoString() string {
	return s.String()
}

func (s *GetLifecycleRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetLifecycleRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *GetLifecycleRequest) GetOrder() *string {
	return s.Order
}

func (s *GetLifecycleRequest) GetSessionNumber() *int32 {
	return s.SessionNumber
}

func (s *GetLifecycleRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetLifecycleRequest) GetToken() *string {
	return s.Token
}

func (s *GetLifecycleRequest) SetEndTime(v string) *GetLifecycleRequest {
	s.EndTime = &v
	return s
}

func (s *GetLifecycleRequest) SetLimit(v int32) *GetLifecycleRequest {
	s.Limit = &v
	return s
}

func (s *GetLifecycleRequest) SetOrder(v string) *GetLifecycleRequest {
	s.Order = &v
	return s
}

func (s *GetLifecycleRequest) SetSessionNumber(v int32) *GetLifecycleRequest {
	s.SessionNumber = &v
	return s
}

func (s *GetLifecycleRequest) SetStartTime(v string) *GetLifecycleRequest {
	s.StartTime = &v
	return s
}

func (s *GetLifecycleRequest) SetToken(v string) *GetLifecycleRequest {
	s.Token = &v
	return s
}

func (s *GetLifecycleRequest) Validate() error {
	return dara.Validate(s)
}
