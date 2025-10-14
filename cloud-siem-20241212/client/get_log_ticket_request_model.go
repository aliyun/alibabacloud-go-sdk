// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetLogTicketRequest
	GetLang() *string
	SetLogUserId(v int64) *GetLogTicketRequest
	GetLogUserId() *int64
	SetRegionId(v string) *GetLogTicketRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetLogTicketRequest
	GetRoleFor() *int64
}

type GetLogTicketRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 173326*******。
	LogUserId *int64 `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s GetLogTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogTicketRequest) GoString() string {
	return s.String()
}

func (s *GetLogTicketRequest) GetLang() *string {
	return s.Lang
}

func (s *GetLogTicketRequest) GetLogUserId() *int64 {
	return s.LogUserId
}

func (s *GetLogTicketRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLogTicketRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetLogTicketRequest) SetLang(v string) *GetLogTicketRequest {
	s.Lang = &v
	return s
}

func (s *GetLogTicketRequest) SetLogUserId(v int64) *GetLogTicketRequest {
	s.LogUserId = &v
	return s
}

func (s *GetLogTicketRequest) SetRegionId(v string) *GetLogTicketRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogTicketRequest) SetRoleFor(v int64) *GetLogTicketRequest {
	s.RoleFor = &v
	return s
}

func (s *GetLogTicketRequest) Validate() error {
	return dara.Validate(s)
}
