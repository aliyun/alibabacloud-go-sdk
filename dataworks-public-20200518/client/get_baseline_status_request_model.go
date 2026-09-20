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
	// The business date in UTC format (yyyy-MM-dd\\"T\\"HH:mm:ssZ).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-07-07T00:00:00+0800
	Bizdate *string `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The cycle number of the baseline instance. The value is 1 for daily baselines. The value ranges from [1,24\\] for hourly baselines.
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
