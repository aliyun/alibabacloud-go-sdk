// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineKeyPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v int64) *GetBaselineKeyPathRequest
	GetBaselineId() *int64
	SetBizdate(v string) *GetBaselineKeyPathRequest
	GetBizdate() *string
	SetInGroupId(v int32) *GetBaselineKeyPathRequest
	GetInGroupId() *int32
}

type GetBaselineKeyPathRequest struct {
	// The ID of the baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The business date in the yyyy-MM-dd\\"T\\"HH:mm:ssZ UTC format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-07-07T00:00:00+0800
	Bizdate *string `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The cycle number of the baseline instance. The value is 1 for a daily baseline. Valid values for an hourly baseline: [1,24\\].
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	InGroupId *int32 `json:"InGroupId,omitempty" xml:"InGroupId,omitempty"`
}

func (s GetBaselineKeyPathRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineKeyPathRequest) GoString() string {
	return s.String()
}

func (s *GetBaselineKeyPathRequest) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetBaselineKeyPathRequest) GetBizdate() *string {
	return s.Bizdate
}

func (s *GetBaselineKeyPathRequest) GetInGroupId() *int32 {
	return s.InGroupId
}

func (s *GetBaselineKeyPathRequest) SetBaselineId(v int64) *GetBaselineKeyPathRequest {
	s.BaselineId = &v
	return s
}

func (s *GetBaselineKeyPathRequest) SetBizdate(v string) *GetBaselineKeyPathRequest {
	s.Bizdate = &v
	return s
}

func (s *GetBaselineKeyPathRequest) SetInGroupId(v int32) *GetBaselineKeyPathRequest {
	s.InGroupId = &v
	return s
}

func (s *GetBaselineKeyPathRequest) Validate() error {
	return dara.Validate(s)
}
