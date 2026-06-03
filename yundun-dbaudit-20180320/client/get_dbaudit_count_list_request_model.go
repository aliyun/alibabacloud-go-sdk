// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBAuditCountListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetDBAuditCountListRequest
	GetBeginDate() *string
	SetDbId(v int32) *GetDBAuditCountListRequest
	GetDbId() *int32
	SetEndDate(v string) *GetDBAuditCountListRequest
	GetEndDate() *string
	SetInstanceId(v string) *GetDBAuditCountListRequest
	GetInstanceId() *string
	SetLang(v string) *GetDBAuditCountListRequest
	GetLang() *string
	SetRegionId(v string) *GetDBAuditCountListRequest
	GetRegionId() *string
}

type GetDBAuditCountListRequest struct {
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

func (s GetDBAuditCountListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDBAuditCountListRequest) GoString() string {
	return s.String()
}

func (s *GetDBAuditCountListRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetDBAuditCountListRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetDBAuditCountListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetDBAuditCountListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDBAuditCountListRequest) GetLang() *string {
	return s.Lang
}

func (s *GetDBAuditCountListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDBAuditCountListRequest) SetBeginDate(v string) *GetDBAuditCountListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetDBAuditCountListRequest) SetDbId(v int32) *GetDBAuditCountListRequest {
	s.DbId = &v
	return s
}

func (s *GetDBAuditCountListRequest) SetEndDate(v string) *GetDBAuditCountListRequest {
	s.EndDate = &v
	return s
}

func (s *GetDBAuditCountListRequest) SetInstanceId(v string) *GetDBAuditCountListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDBAuditCountListRequest) SetLang(v string) *GetDBAuditCountListRequest {
	s.Lang = &v
	return s
}

func (s *GetDBAuditCountListRequest) SetRegionId(v string) *GetDBAuditCountListRequest {
	s.RegionId = &v
	return s
}

func (s *GetDBAuditCountListRequest) Validate() error {
	return dara.Validate(s)
}
