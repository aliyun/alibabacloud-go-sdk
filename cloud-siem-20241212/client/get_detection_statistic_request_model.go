// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectionStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetDetectionStatisticRequest
	GetLang() *string
	SetRegionId(v string) *GetDetectionStatisticRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetDetectionStatisticRequest
	GetRoleFor() *int64
}

type GetDetectionStatisticRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: China (Hangzhou). Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Singapore. Your assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s GetDetectionStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDetectionStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetDetectionStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *GetDetectionStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDetectionStatisticRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetDetectionStatisticRequest) SetLang(v string) *GetDetectionStatisticRequest {
	s.Lang = &v
	return s
}

func (s *GetDetectionStatisticRequest) SetRegionId(v string) *GetDetectionStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *GetDetectionStatisticRequest) SetRoleFor(v int64) *GetDetectionStatisticRequest {
	s.RoleFor = &v
	return s
}

func (s *GetDetectionStatisticRequest) Validate() error {
	return dara.Validate(s)
}
