// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApgroupDescendantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApgroupId(v int64) *ListApgroupDescendantRequest
	GetApgroupId() *int64
	SetApgroupUuid(v string) *ListApgroupDescendantRequest
	GetApgroupUuid() *string
	SetAppCode(v string) *ListApgroupDescendantRequest
	GetAppCode() *string
	SetAppName(v string) *ListApgroupDescendantRequest
	GetAppName() *string
	SetCursor(v int64) *ListApgroupDescendantRequest
	GetCursor() *int64
	SetLevel(v int64) *ListApgroupDescendantRequest
	GetLevel() *int64
	SetPageSize(v int32) *ListApgroupDescendantRequest
	GetPageSize() *int32
}

type ListApgroupDescendantRequest struct {
	// This parameter is required.
	ApgroupId   *int64  `json:"ApgroupId,omitempty" xml:"ApgroupId,omitempty"`
	ApgroupUuid *string `json:"ApgroupUuid,omitempty" xml:"ApgroupUuid,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Cursor   *int64  `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	Level    *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApgroupDescendantRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApgroupDescendantRequest) GoString() string {
	return s.String()
}

func (s *ListApgroupDescendantRequest) GetApgroupId() *int64 {
	return s.ApgroupId
}

func (s *ListApgroupDescendantRequest) GetApgroupUuid() *string {
	return s.ApgroupUuid
}

func (s *ListApgroupDescendantRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *ListApgroupDescendantRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListApgroupDescendantRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *ListApgroupDescendantRequest) GetLevel() *int64 {
	return s.Level
}

func (s *ListApgroupDescendantRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApgroupDescendantRequest) SetApgroupId(v int64) *ListApgroupDescendantRequest {
	s.ApgroupId = &v
	return s
}

func (s *ListApgroupDescendantRequest) SetApgroupUuid(v string) *ListApgroupDescendantRequest {
	s.ApgroupUuid = &v
	return s
}

func (s *ListApgroupDescendantRequest) SetAppCode(v string) *ListApgroupDescendantRequest {
	s.AppCode = &v
	return s
}

func (s *ListApgroupDescendantRequest) SetAppName(v string) *ListApgroupDescendantRequest {
	s.AppName = &v
	return s
}

func (s *ListApgroupDescendantRequest) SetCursor(v int64) *ListApgroupDescendantRequest {
	s.Cursor = &v
	return s
}

func (s *ListApgroupDescendantRequest) SetLevel(v int64) *ListApgroupDescendantRequest {
	s.Level = &v
	return s
}

func (s *ListApgroupDescendantRequest) SetPageSize(v int32) *ListApgroupDescendantRequest {
	s.PageSize = &v
	return s
}

func (s *ListApgroupDescendantRequest) Validate() error {
	return dara.Validate(s)
}
