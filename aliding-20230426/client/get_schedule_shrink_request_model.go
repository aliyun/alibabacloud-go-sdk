// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetScheduleShrinkRequest
	GetEndTime() *string
	SetStartTime(v string) *GetScheduleShrinkRequest
	GetStartTime() *string
	SetTenantContextShrink(v string) *GetScheduleShrinkRequest
	GetTenantContextShrink() *string
	SetUserIdsShrink(v string) *GetScheduleShrinkRequest
	GetUserIdsShrink() *string
}

type GetScheduleShrinkRequest struct {
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	UserIdsShrink       *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s GetScheduleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetScheduleShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetScheduleShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetScheduleShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *GetScheduleShrinkRequest) SetEndTime(v string) *GetScheduleShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetScheduleShrinkRequest) SetStartTime(v string) *GetScheduleShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetScheduleShrinkRequest) SetTenantContextShrink(v string) *GetScheduleShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetScheduleShrinkRequest) SetUserIdsShrink(v string) *GetScheduleShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *GetScheduleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
