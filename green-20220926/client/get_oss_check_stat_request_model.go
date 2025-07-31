// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetByMonth(v bool) *GetOssCheckStatRequest
	GetByMonth() *bool
	SetEndDate(v string) *GetOssCheckStatRequest
	GetEndDate() *string
	SetParentTaskId(v string) *GetOssCheckStatRequest
	GetParentTaskId() *string
	SetRegionId(v string) *GetOssCheckStatRequest
	GetRegionId() *string
	SetStartDate(v string) *GetOssCheckStatRequest
	GetStartDate() *string
}

type GetOssCheckStatRequest struct {
	// example:
	//
	// true
	ByMonth *bool `json:"ByMonth,omitempty" xml:"ByMonth,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// P_UNHBH
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetOssCheckStatRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckStatRequest) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatRequest) GetByMonth() *bool {
	return s.ByMonth
}

func (s *GetOssCheckStatRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetOssCheckStatRequest) GetParentTaskId() *string {
	return s.ParentTaskId
}

func (s *GetOssCheckStatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOssCheckStatRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetOssCheckStatRequest) SetByMonth(v bool) *GetOssCheckStatRequest {
	s.ByMonth = &v
	return s
}

func (s *GetOssCheckStatRequest) SetEndDate(v string) *GetOssCheckStatRequest {
	s.EndDate = &v
	return s
}

func (s *GetOssCheckStatRequest) SetParentTaskId(v string) *GetOssCheckStatRequest {
	s.ParentTaskId = &v
	return s
}

func (s *GetOssCheckStatRequest) SetRegionId(v string) *GetOssCheckStatRequest {
	s.RegionId = &v
	return s
}

func (s *GetOssCheckStatRequest) SetStartDate(v string) *GetOssCheckStatRequest {
	s.StartDate = &v
	return s
}

func (s *GetOssCheckStatRequest) Validate() error {
	return dara.Validate(s)
}
