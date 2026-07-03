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
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
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
	// The region where the data management center of the threat analysis feature is located. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets reside outside China.
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
