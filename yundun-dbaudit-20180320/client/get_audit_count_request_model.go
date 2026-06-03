// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetAuditCountRequest
	GetBeginDate() *string
	SetDbId(v int32) *GetAuditCountRequest
	GetDbId() *int32
	SetEndDate(v string) *GetAuditCountRequest
	GetEndDate() *string
	SetInstanceId(v string) *GetAuditCountRequest
	GetInstanceId() *string
	SetLang(v string) *GetAuditCountRequest
	GetLang() *string
	SetRegionId(v string) *GetAuditCountRequest
	GetRegionId() *string
}

type GetAuditCountRequest struct {
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 0
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// 2019-06-06 23:59:59
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAuditCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuditCountRequest) GoString() string {
	return s.String()
}

func (s *GetAuditCountRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetAuditCountRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetAuditCountRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetAuditCountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAuditCountRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAuditCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAuditCountRequest) SetBeginDate(v string) *GetAuditCountRequest {
	s.BeginDate = &v
	return s
}

func (s *GetAuditCountRequest) SetDbId(v int32) *GetAuditCountRequest {
	s.DbId = &v
	return s
}

func (s *GetAuditCountRequest) SetEndDate(v string) *GetAuditCountRequest {
	s.EndDate = &v
	return s
}

func (s *GetAuditCountRequest) SetInstanceId(v string) *GetAuditCountRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAuditCountRequest) SetLang(v string) *GetAuditCountRequest {
	s.Lang = &v
	return s
}

func (s *GetAuditCountRequest) SetRegionId(v string) *GetAuditCountRequest {
	s.RegionId = &v
	return s
}

func (s *GetAuditCountRequest) Validate() error {
	return dara.Validate(s)
}
