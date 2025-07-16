// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableDataByFormInstanceIdTableIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *ListTableDataByFormInstanceIdTableIdRequest
	GetAppType() *string
	SetFormInstanceId(v string) *ListTableDataByFormInstanceIdTableIdRequest
	GetFormInstanceId() *string
	SetFormUuid(v string) *ListTableDataByFormInstanceIdTableIdRequest
	GetFormUuid() *string
	SetPageNumber(v int32) *ListTableDataByFormInstanceIdTableIdRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTableDataByFormInstanceIdTableIdRequest
	GetPageSize() *int32
	SetSystemToken(v string) *ListTableDataByFormInstanceIdTableIdRequest
	GetSystemToken() *string
	SetTableFieldId(v string) *ListTableDataByFormInstanceIdTableIdRequest
	GetTableFieldId() *string
}

type ListTableDataByFormInstanceIdTableIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0xxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM_PBKT0xxx
	FormInstanceId *string `json:"FormInstanceId,omitempty" xml:"FormInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-xxxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// 20
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
	// This parameter is required.
	//
	// example:
	//
	// 1111
	TableFieldId *string `json:"TableFieldId,omitempty" xml:"TableFieldId,omitempty"`
}

func (s ListTableDataByFormInstanceIdTableIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTableDataByFormInstanceIdTableIdRequest) GoString() string {
	return s.String()
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) GetAppType() *string {
	return s.AppType
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) GetTableFieldId() *string {
	return s.TableFieldId
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetAppType(v string) *ListTableDataByFormInstanceIdTableIdRequest {
	s.AppType = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetFormInstanceId(v string) *ListTableDataByFormInstanceIdTableIdRequest {
	s.FormInstanceId = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetFormUuid(v string) *ListTableDataByFormInstanceIdTableIdRequest {
	s.FormUuid = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetPageNumber(v int32) *ListTableDataByFormInstanceIdTableIdRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetPageSize(v int32) *ListTableDataByFormInstanceIdTableIdRequest {
	s.PageSize = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetSystemToken(v string) *ListTableDataByFormInstanceIdTableIdRequest {
	s.SystemToken = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetTableFieldId(v string) *ListTableDataByFormInstanceIdTableIdRequest {
	s.TableFieldId = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) Validate() error {
	return dara.Validate(s)
}
