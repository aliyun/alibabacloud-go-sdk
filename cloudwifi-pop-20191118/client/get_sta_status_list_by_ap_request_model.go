// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStaStatusListByApRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *GetStaStatusListByApRequest
	GetApMac() *string
	SetAppCode(v string) *GetStaStatusListByApRequest
	GetAppCode() *string
	SetAppName(v string) *GetStaStatusListByApRequest
	GetAppName() *string
	SetCursor(v int64) *GetStaStatusListByApRequest
	GetCursor() *int64
	SetPageSize(v int32) *GetStaStatusListByApRequest
	GetPageSize() *int32
}

type GetStaStatusListByApRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Cursor *int64 `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetStaStatusListByApRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStaStatusListByApRequest) GoString() string {
	return s.String()
}

func (s *GetStaStatusListByApRequest) GetApMac() *string {
	return s.ApMac
}

func (s *GetStaStatusListByApRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetStaStatusListByApRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetStaStatusListByApRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *GetStaStatusListByApRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetStaStatusListByApRequest) SetApMac(v string) *GetStaStatusListByApRequest {
	s.ApMac = &v
	return s
}

func (s *GetStaStatusListByApRequest) SetAppCode(v string) *GetStaStatusListByApRequest {
	s.AppCode = &v
	return s
}

func (s *GetStaStatusListByApRequest) SetAppName(v string) *GetStaStatusListByApRequest {
	s.AppName = &v
	return s
}

func (s *GetStaStatusListByApRequest) SetCursor(v int64) *GetStaStatusListByApRequest {
	s.Cursor = &v
	return s
}

func (s *GetStaStatusListByApRequest) SetPageSize(v int32) *GetStaStatusListByApRequest {
	s.PageSize = &v
	return s
}

func (s *GetStaStatusListByApRequest) Validate() error {
	return dara.Validate(s)
}
