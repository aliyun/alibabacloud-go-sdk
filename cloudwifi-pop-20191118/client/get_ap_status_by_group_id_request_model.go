// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApStatusByGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApgroupId(v string) *GetApStatusByGroupIdRequest
	GetApgroupId() *string
	SetAppCode(v string) *GetApStatusByGroupIdRequest
	GetAppCode() *string
	SetAppName(v string) *GetApStatusByGroupIdRequest
	GetAppName() *string
	SetCursor(v int64) *GetApStatusByGroupIdRequest
	GetCursor() *int64
	SetPageSize(v int32) *GetApStatusByGroupIdRequest
	GetPageSize() *int32
}

type GetApStatusByGroupIdRequest struct {
	// This parameter is required.
	ApgroupId *string `json:"ApgroupId,omitempty" xml:"ApgroupId,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Cursor *int64 `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetApStatusByGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApStatusByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *GetApStatusByGroupIdRequest) GetApgroupId() *string {
	return s.ApgroupId
}

func (s *GetApStatusByGroupIdRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApStatusByGroupIdRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApStatusByGroupIdRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *GetApStatusByGroupIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetApStatusByGroupIdRequest) SetApgroupId(v string) *GetApStatusByGroupIdRequest {
	s.ApgroupId = &v
	return s
}

func (s *GetApStatusByGroupIdRequest) SetAppCode(v string) *GetApStatusByGroupIdRequest {
	s.AppCode = &v
	return s
}

func (s *GetApStatusByGroupIdRequest) SetAppName(v string) *GetApStatusByGroupIdRequest {
	s.AppName = &v
	return s
}

func (s *GetApStatusByGroupIdRequest) SetCursor(v int64) *GetApStatusByGroupIdRequest {
	s.Cursor = &v
	return s
}

func (s *GetApStatusByGroupIdRequest) SetPageSize(v int32) *GetApStatusByGroupIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetApStatusByGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
