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
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
