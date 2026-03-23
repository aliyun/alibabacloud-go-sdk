// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupMiscAggTimeSerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApgroupUuid(v string) *GetGroupMiscAggTimeSerRequest
	GetApgroupUuid() *string
	SetAppCode(v string) *GetGroupMiscAggTimeSerRequest
	GetAppCode() *string
	SetAppName(v string) *GetGroupMiscAggTimeSerRequest
	GetAppName() *string
	SetEnd(v int64) *GetGroupMiscAggTimeSerRequest
	GetEnd() *int64
	SetStart(v int64) *GetGroupMiscAggTimeSerRequest
	GetStart() *int64
}

type GetGroupMiscAggTimeSerRequest struct {
	// This parameter is required.
	ApgroupUuid *string `json:"ApgroupUuid,omitempty" xml:"ApgroupUuid,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// This parameter is required.
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetGroupMiscAggTimeSerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupMiscAggTimeSerRequest) GoString() string {
	return s.String()
}

func (s *GetGroupMiscAggTimeSerRequest) GetApgroupUuid() *string {
	return s.ApgroupUuid
}

func (s *GetGroupMiscAggTimeSerRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetGroupMiscAggTimeSerRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetGroupMiscAggTimeSerRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetGroupMiscAggTimeSerRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetGroupMiscAggTimeSerRequest) SetApgroupUuid(v string) *GetGroupMiscAggTimeSerRequest {
	s.ApgroupUuid = &v
	return s
}

func (s *GetGroupMiscAggTimeSerRequest) SetAppCode(v string) *GetGroupMiscAggTimeSerRequest {
	s.AppCode = &v
	return s
}

func (s *GetGroupMiscAggTimeSerRequest) SetAppName(v string) *GetGroupMiscAggTimeSerRequest {
	s.AppName = &v
	return s
}

func (s *GetGroupMiscAggTimeSerRequest) SetEnd(v int64) *GetGroupMiscAggTimeSerRequest {
	s.End = &v
	return s
}

func (s *GetGroupMiscAggTimeSerRequest) SetStart(v int64) *GetGroupMiscAggTimeSerRequest {
	s.Start = &v
	return s
}

func (s *GetGroupMiscAggTimeSerRequest) Validate() error {
	return dara.Validate(s)
}
