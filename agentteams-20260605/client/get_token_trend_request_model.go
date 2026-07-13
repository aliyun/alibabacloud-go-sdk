// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetTokenTrendRequest
	GetEndTime() *string
	SetGroupBy(v string) *GetTokenTrendRequest
	GetGroupBy() *string
	SetInstanceId(v string) *GetTokenTrendRequest
	GetInstanceId() *string
	SetStartTime(v string) *GetTokenTrendRequest
	GetStartTime() *string
}

type GetTokenTrendRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupBy    *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetTokenTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenTrendRequest) GoString() string {
	return s.String()
}

func (s *GetTokenTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTokenTrendRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *GetTokenTrendRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTokenTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTokenTrendRequest) SetEndTime(v string) *GetTokenTrendRequest {
	s.EndTime = &v
	return s
}

func (s *GetTokenTrendRequest) SetGroupBy(v string) *GetTokenTrendRequest {
	s.GroupBy = &v
	return s
}

func (s *GetTokenTrendRequest) SetInstanceId(v string) *GetTokenTrendRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTokenTrendRequest) SetStartTime(v string) *GetTokenTrendRequest {
	s.StartTime = &v
	return s
}

func (s *GetTokenTrendRequest) Validate() error {
	return dara.Validate(s)
}
