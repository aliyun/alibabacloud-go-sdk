// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFormDataByIdListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *BatchGetFormDataByIdListShrinkRequest
	GetAppType() *string
	SetFormInstanceIdListShrink(v string) *BatchGetFormDataByIdListShrinkRequest
	GetFormInstanceIdListShrink() *string
	SetFormUuid(v string) *BatchGetFormDataByIdListShrinkRequest
	GetFormUuid() *string
	SetNeedFormInstanceValue(v bool) *BatchGetFormDataByIdListShrinkRequest
	GetNeedFormInstanceValue() *bool
	SetSystemToken(v string) *BatchGetFormDataByIdListShrinkRequest
	GetSystemToken() *string
}

type BatchGetFormDataByIdListShrinkRequest struct {
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
	// example:
	//
	// true
	NeedFormInstanceValue *bool `json:"NeedFormInstanceValue,omitempty" xml:"NeedFormInstanceValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s BatchGetFormDataByIdListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *BatchGetFormDataByIdListShrinkRequest) GetFormInstanceIdListShrink() *string {
	return s.FormInstanceIdListShrink
}

func (s *BatchGetFormDataByIdListShrinkRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchGetFormDataByIdListShrinkRequest) GetNeedFormInstanceValue() *bool {
	return s.NeedFormInstanceValue
}

func (s *BatchGetFormDataByIdListShrinkRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *BatchGetFormDataByIdListShrinkRequest) SetAppType(v string) *BatchGetFormDataByIdListShrinkRequest {
	s.AppType = &v
	return s
}

func (s *BatchGetFormDataByIdListShrinkRequest) SetFormInstanceIdListShrink(v string) *BatchGetFormDataByIdListShrinkRequest {
	s.FormInstanceIdListShrink = &v
	return s
}

func (s *BatchGetFormDataByIdListShrinkRequest) SetFormUuid(v string) *BatchGetFormDataByIdListShrinkRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchGetFormDataByIdListShrinkRequest) SetNeedFormInstanceValue(v bool) *BatchGetFormDataByIdListShrinkRequest {
	s.NeedFormInstanceValue = &v
	return s
}

func (s *BatchGetFormDataByIdListShrinkRequest) SetSystemToken(v string) *BatchGetFormDataByIdListShrinkRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchGetFormDataByIdListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
