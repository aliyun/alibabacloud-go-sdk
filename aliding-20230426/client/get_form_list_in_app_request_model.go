// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormListInAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetFormListInAppRequest
	GetAppType() *string
	SetFormTypes(v string) *GetFormListInAppRequest
	GetFormTypes() *string
	SetPageNumber(v int32) *GetFormListInAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetFormListInAppRequest
	GetPageSize() *int32
	SetSystemToken(v string) *GetFormListInAppRequest
	GetSystemToken() *string
}

type GetFormListInAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0xxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// receipt
	FormTypes *string `json:"FormTypes,omitempty" xml:"FormTypes,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s GetFormListInAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFormListInAppRequest) GoString() string {
	return s.String()
}

func (s *GetFormListInAppRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetFormListInAppRequest) GetFormTypes() *string {
	return s.FormTypes
}

func (s *GetFormListInAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetFormListInAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetFormListInAppRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetFormListInAppRequest) SetAppType(v string) *GetFormListInAppRequest {
	s.AppType = &v
	return s
}

func (s *GetFormListInAppRequest) SetFormTypes(v string) *GetFormListInAppRequest {
	s.FormTypes = &v
	return s
}

func (s *GetFormListInAppRequest) SetPageNumber(v int32) *GetFormListInAppRequest {
	s.PageNumber = &v
	return s
}

func (s *GetFormListInAppRequest) SetPageSize(v int32) *GetFormListInAppRequest {
	s.PageSize = &v
	return s
}

func (s *GetFormListInAppRequest) SetSystemToken(v string) *GetFormListInAppRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormListInAppRequest) Validate() error {
	return dara.Validate(s)
}
