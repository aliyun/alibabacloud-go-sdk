// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackEventDashboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetAttackEventDashboardRequest
	GetEndTime() *int64
	SetLang(v string) *GetAttackEventDashboardRequest
	GetLang() *string
	SetStartTime(v int64) *GetAttackEventDashboardRequest
	GetStartTime() *int64
}

type GetAttackEventDashboardRequest struct {
	// Timestamp of the end time.
	//
	// example:
	//
	// 1753153137284
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Language type for request and response messages. Default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Timestamp of the start time.
	//
	// example:
	//
	// 1752548337284
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAttackEventDashboardRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttackEventDashboardRequest) GoString() string {
	return s.String()
}

func (s *GetAttackEventDashboardRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetAttackEventDashboardRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAttackEventDashboardRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetAttackEventDashboardRequest) SetEndTime(v int64) *GetAttackEventDashboardRequest {
	s.EndTime = &v
	return s
}

func (s *GetAttackEventDashboardRequest) SetLang(v string) *GetAttackEventDashboardRequest {
	s.Lang = &v
	return s
}

func (s *GetAttackEventDashboardRequest) SetStartTime(v int64) *GetAttackEventDashboardRequest {
	s.StartTime = &v
	return s
}

func (s *GetAttackEventDashboardRequest) Validate() error {
	return dara.Validate(s)
}
