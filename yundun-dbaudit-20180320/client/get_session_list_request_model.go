// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionList(v []*string) *GetSessionListRequest
	GetActionList() []*string
	SetBeginDate(v string) *GetSessionListRequest
	GetBeginDate() *string
	SetClientIpList(v []*string) *GetSessionListRequest
	GetClientIpList() []*string
	SetClientProgramList(v []*string) *GetSessionListRequest
	GetClientProgramList() []*string
	SetDbHostList(v []*string) *GetSessionListRequest
	GetDbHostList() []*string
	SetDbId(v int32) *GetSessionListRequest
	GetDbId() *int32
	SetDbUserList(v []*string) *GetSessionListRequest
	GetDbUserList() []*string
	SetEndDate(v string) *GetSessionListRequest
	GetEndDate() *string
	SetInstanceId(v string) *GetSessionListRequest
	GetInstanceId() *string
	SetLang(v string) *GetSessionListRequest
	GetLang() *string
	SetPageNumber(v int32) *GetSessionListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetSessionListRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetSessionListRequest
	GetRegionId() *string
	SetSessionId(v string) *GetSessionListRequest
	GetSessionId() *string
}

type GetSessionListRequest struct {
	// example:
	//
	// 0
	ActionList []*string `json:"ActionList,omitempty" xml:"ActionList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	ClientIpList []*string `json:"ClientIpList,omitempty" xml:"ClientIpList,omitempty" type:"Repeated"`
	// example:
	//
	// navicat
	ClientProgramList []*string `json:"ClientProgramList,omitempty" xml:"ClientProgramList,omitempty" type:"Repeated"`
	// example:
	//
	// 192.168.XX.XX
	DbHostList []*string `json:"DbHostList,omitempty" xml:"DbHostList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// root
	DbUserList []*string `json:"DbUserList,omitempty" xml:"DbUserList,omitempty" type:"Repeated"`
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
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3011610850021000000
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetSessionListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSessionListRequest) GoString() string {
	return s.String()
}

func (s *GetSessionListRequest) GetActionList() []*string {
	return s.ActionList
}

func (s *GetSessionListRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetSessionListRequest) GetClientIpList() []*string {
	return s.ClientIpList
}

func (s *GetSessionListRequest) GetClientProgramList() []*string {
	return s.ClientProgramList
}

func (s *GetSessionListRequest) GetDbHostList() []*string {
	return s.DbHostList
}

func (s *GetSessionListRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetSessionListRequest) GetDbUserList() []*string {
	return s.DbUserList
}

func (s *GetSessionListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetSessionListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSessionListRequest) GetLang() *string {
	return s.Lang
}

func (s *GetSessionListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetSessionListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSessionListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSessionListRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetSessionListRequest) SetActionList(v []*string) *GetSessionListRequest {
	s.ActionList = v
	return s
}

func (s *GetSessionListRequest) SetBeginDate(v string) *GetSessionListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetSessionListRequest) SetClientIpList(v []*string) *GetSessionListRequest {
	s.ClientIpList = v
	return s
}

func (s *GetSessionListRequest) SetClientProgramList(v []*string) *GetSessionListRequest {
	s.ClientProgramList = v
	return s
}

func (s *GetSessionListRequest) SetDbHostList(v []*string) *GetSessionListRequest {
	s.DbHostList = v
	return s
}

func (s *GetSessionListRequest) SetDbId(v int32) *GetSessionListRequest {
	s.DbId = &v
	return s
}

func (s *GetSessionListRequest) SetDbUserList(v []*string) *GetSessionListRequest {
	s.DbUserList = v
	return s
}

func (s *GetSessionListRequest) SetEndDate(v string) *GetSessionListRequest {
	s.EndDate = &v
	return s
}

func (s *GetSessionListRequest) SetInstanceId(v string) *GetSessionListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSessionListRequest) SetLang(v string) *GetSessionListRequest {
	s.Lang = &v
	return s
}

func (s *GetSessionListRequest) SetPageNumber(v int32) *GetSessionListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSessionListRequest) SetPageSize(v int32) *GetSessionListRequest {
	s.PageSize = &v
	return s
}

func (s *GetSessionListRequest) SetRegionId(v string) *GetSessionListRequest {
	s.RegionId = &v
	return s
}

func (s *GetSessionListRequest) SetSessionId(v string) *GetSessionListRequest {
	s.SessionId = &v
	return s
}

func (s *GetSessionListRequest) Validate() error {
	return dara.Validate(s)
}
