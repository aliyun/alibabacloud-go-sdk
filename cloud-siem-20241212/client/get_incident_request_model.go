// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIncidentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncidentUuid(v string) *GetIncidentRequest
	GetIncidentUuid() *string
	SetLang(v string) *GetIncidentRequest
	GetLang() *string
	SetRegionId(v string) *GetIncidentRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetIncidentRequest
	GetRoleFor() *int64
}

type GetIncidentRequest struct {
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
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

func (s GetIncidentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIncidentRequest) GoString() string {
	return s.String()
}

func (s *GetIncidentRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *GetIncidentRequest) GetLang() *string {
	return s.Lang
}

func (s *GetIncidentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIncidentRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetIncidentRequest) SetIncidentUuid(v string) *GetIncidentRequest {
	s.IncidentUuid = &v
	return s
}

func (s *GetIncidentRequest) SetLang(v string) *GetIncidentRequest {
	s.Lang = &v
	return s
}

func (s *GetIncidentRequest) SetRegionId(v string) *GetIncidentRequest {
	s.RegionId = &v
	return s
}

func (s *GetIncidentRequest) SetRoleFor(v int64) *GetIncidentRequest {
	s.RoleFor = &v
	return s
}

func (s *GetIncidentRequest) Validate() error {
	return dara.Validate(s)
}
