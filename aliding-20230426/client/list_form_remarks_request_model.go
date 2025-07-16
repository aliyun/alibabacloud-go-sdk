// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFormRemarksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *ListFormRemarksRequest
	GetAppType() *string
	SetFormInstanceIdList(v []*string) *ListFormRemarksRequest
	GetFormInstanceIdList() []*string
	SetFormUuid(v string) *ListFormRemarksRequest
	GetFormUuid() *string
	SetSystemToken(v string) *ListFormRemarksRequest
	GetSystemToken() *string
}

type ListFormRemarksRequest struct {
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
	// FORM-xxxxx
	FormInstanceIdList []*string `json:"FormInstanceIdList,omitempty" xml:"FormInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-xxxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s ListFormRemarksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFormRemarksRequest) GoString() string {
	return s.String()
}

func (s *ListFormRemarksRequest) GetAppType() *string {
	return s.AppType
}

func (s *ListFormRemarksRequest) GetFormInstanceIdList() []*string {
	return s.FormInstanceIdList
}

func (s *ListFormRemarksRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *ListFormRemarksRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *ListFormRemarksRequest) SetAppType(v string) *ListFormRemarksRequest {
	s.AppType = &v
	return s
}

func (s *ListFormRemarksRequest) SetFormInstanceIdList(v []*string) *ListFormRemarksRequest {
	s.FormInstanceIdList = v
	return s
}

func (s *ListFormRemarksRequest) SetFormUuid(v string) *ListFormRemarksRequest {
	s.FormUuid = &v
	return s
}

func (s *ListFormRemarksRequest) SetSystemToken(v string) *ListFormRemarksRequest {
	s.SystemToken = &v
	return s
}

func (s *ListFormRemarksRequest) Validate() error {
	return dara.Validate(s)
}
