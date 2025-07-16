// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFormRemarksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *ListFormRemarksShrinkRequest
	GetAppType() *string
	SetFormInstanceIdListShrink(v string) *ListFormRemarksShrinkRequest
	GetFormInstanceIdListShrink() *string
	SetFormUuid(v string) *ListFormRemarksShrinkRequest
	GetFormUuid() *string
	SetSystemToken(v string) *ListFormRemarksShrinkRequest
	GetSystemToken() *string
}

type ListFormRemarksShrinkRequest struct {
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
	FormInstanceIdListShrink *string `json:"FormInstanceIdList,omitempty" xml:"FormInstanceIdList,omitempty"`
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

func (s ListFormRemarksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFormRemarksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFormRemarksShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *ListFormRemarksShrinkRequest) GetFormInstanceIdListShrink() *string {
	return s.FormInstanceIdListShrink
}

func (s *ListFormRemarksShrinkRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *ListFormRemarksShrinkRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *ListFormRemarksShrinkRequest) SetAppType(v string) *ListFormRemarksShrinkRequest {
	s.AppType = &v
	return s
}

func (s *ListFormRemarksShrinkRequest) SetFormInstanceIdListShrink(v string) *ListFormRemarksShrinkRequest {
	s.FormInstanceIdListShrink = &v
	return s
}

func (s *ListFormRemarksShrinkRequest) SetFormUuid(v string) *ListFormRemarksShrinkRequest {
	s.FormUuid = &v
	return s
}

func (s *ListFormRemarksShrinkRequest) SetSystemToken(v string) *ListFormRemarksShrinkRequest {
	s.SystemToken = &v
	return s
}

func (s *ListFormRemarksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
