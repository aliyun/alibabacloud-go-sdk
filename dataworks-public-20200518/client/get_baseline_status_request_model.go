// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v int64) *GetBaselineStatusRequest
	GetBaselineId() *int64
	SetBizdate(v string) *GetBaselineStatusRequest
	GetBizdate() *string
	SetInGroupId(v int32) *GetBaselineStatusRequest
	GetInGroupId() *int32
}

type GetBaselineStatusRequest struct {
	// The ID of the baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The data timestamp of the baseline instance. Specify the time in the ISO 8601 standard in the yyyy-MM-dd\\"T\\"HH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-07-07T00:00:00+0800
	Bizdate *string `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The ID of the scheduling cycle of the baseline instance. For a baseline instance that is scheduled by day, the value of this parameter is 1. For a baseline instance that is scheduled by hour, the value of this parameter ranges from 1 to 24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	InGroupId *int32 `json:"InGroupId,omitempty" xml:"InGroupId,omitempty"`
}

func (s GetBaselineStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineStatusRequest) GoString() string {
	return s.String()
}

func (s *GetBaselineStatusRequest) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetBaselineStatusRequest) GetBizdate() *string {
	return s.Bizdate
}

func (s *GetBaselineStatusRequest) GetInGroupId() *int32 {
	return s.InGroupId
}

func (s *GetBaselineStatusRequest) SetBaselineId(v int64) *GetBaselineStatusRequest {
	s.BaselineId = &v
	return s
}

func (s *GetBaselineStatusRequest) SetBizdate(v string) *GetBaselineStatusRequest {
	s.Bizdate = &v
	return s
}

func (s *GetBaselineStatusRequest) SetInGroupId(v int32) *GetBaselineStatusRequest {
	s.InGroupId = &v
	return s
}

func (s *GetBaselineStatusRequest) Validate() error {
	return dara.Validate(s)
}
