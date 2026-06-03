// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetLogDetailRequest
	GetBeginDate() *string
	SetEndDate(v string) *GetLogDetailRequest
	GetEndDate() *string
	SetInstanceId(v string) *GetLogDetailRequest
	GetInstanceId() *string
	SetLang(v string) *GetLogDetailRequest
	GetLang() *string
	SetRegionId(v string) *GetLogDetailRequest
	GetRegionId() *string
	SetSqlId(v string) *GetLogDetailRequest
	GetSqlId() *string
}

type GetLogDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// This parameter is required.
	//
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
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1907181552270011186
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
}

func (s GetLogDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogDetailRequest) GoString() string {
	return s.String()
}

func (s *GetLogDetailRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetLogDetailRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetLogDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLogDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *GetLogDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLogDetailRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetLogDetailRequest) SetBeginDate(v string) *GetLogDetailRequest {
	s.BeginDate = &v
	return s
}

func (s *GetLogDetailRequest) SetEndDate(v string) *GetLogDetailRequest {
	s.EndDate = &v
	return s
}

func (s *GetLogDetailRequest) SetInstanceId(v string) *GetLogDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogDetailRequest) SetLang(v string) *GetLogDetailRequest {
	s.Lang = &v
	return s
}

func (s *GetLogDetailRequest) SetRegionId(v string) *GetLogDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogDetailRequest) SetSqlId(v string) *GetLogDetailRequest {
	s.SqlId = &v
	return s
}

func (s *GetLogDetailRequest) Validate() error {
	return dara.Validate(s)
}
