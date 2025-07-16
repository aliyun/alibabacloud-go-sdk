// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *BatchUpdateFormDataByInstanceIdShrinkRequest
	GetAppType() *string
	SetAsynchronousExecution(v bool) *BatchUpdateFormDataByInstanceIdShrinkRequest
	GetAsynchronousExecution() *bool
	SetFormInstanceIdListShrink(v string) *BatchUpdateFormDataByInstanceIdShrinkRequest
	GetFormInstanceIdListShrink() *string
	SetFormUuid(v string) *BatchUpdateFormDataByInstanceIdShrinkRequest
	GetFormUuid() *string
	SetIgnoreEmpty(v bool) *BatchUpdateFormDataByInstanceIdShrinkRequest
	GetIgnoreEmpty() *bool
	SetNoExecuteExpression(v bool) *BatchUpdateFormDataByInstanceIdShrinkRequest
	GetNoExecuteExpression() *bool
	SetSystemToken(v string) *BatchUpdateFormDataByInstanceIdShrinkRequest
	GetSystemToken() *string
	SetUpdateFormDataJson(v string) *BatchUpdateFormDataByInstanceIdShrinkRequest
	GetUpdateFormDataJson() *string
	SetUseLatestFormSchemaVersion(v bool) *BatchUpdateFormDataByInstanceIdShrinkRequest
	GetUseLatestFormSchemaVersion() *bool
}

type BatchUpdateFormDataByInstanceIdShrinkRequest struct {
	// example:
	//
	// String
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// true
	AsynchronousExecution *bool `json:"AsynchronousExecution,omitempty" xml:"AsynchronousExecution,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [ "FINST-J8766S91O2UYN87ZX3XOF1MY8MBA2912BSV0L24" ]
	FormInstanceIdListShrink *string `json:"FormInstanceIdList,omitempty" xml:"FormInstanceIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// true
	IgnoreEmpty *bool `json:"IgnoreEmpty,omitempty" xml:"IgnoreEmpty,omitempty"`
	// example:
	//
	// false
	NoExecuteExpression *bool `json:"NoExecuteExpression,omitempty" xml:"NoExecuteExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\"countrySelectField_l0c1cwiu\":[{\"value\":\"US\"}],\"addressField_l0c1cwiy\":{\"address\":\"111\",\"regionIds\":[460000,469027,469023401],\"regionText\":[{\"en_US\":\"hai+nan+sheng\",\"zh_CN\":\"海南省\"},{\"en_US\":\"cheng+mai+xian\",\"zh_CN\":\"澄迈县\"},{\"en_US\":\"guo+ying+hong+gang+nong+chang\",\"zh_CN\":\"国营红岗农场\"}]}}
	UpdateFormDataJson *string `json:"UpdateFormDataJson,omitempty" xml:"UpdateFormDataJson,omitempty"`
	// example:
	//
	// false
	UseLatestFormSchemaVersion *bool `json:"UseLatestFormSchemaVersion,omitempty" xml:"UseLatestFormSchemaVersion,omitempty"`
}

func (s BatchUpdateFormDataByInstanceIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) GetAsynchronousExecution() *bool {
	return s.AsynchronousExecution
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) GetFormInstanceIdListShrink() *string {
	return s.FormInstanceIdListShrink
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) GetIgnoreEmpty() *bool {
	return s.IgnoreEmpty
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) GetNoExecuteExpression() *bool {
	return s.NoExecuteExpression
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) GetUpdateFormDataJson() *string {
	return s.UpdateFormDataJson
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) GetUseLatestFormSchemaVersion() *bool {
	return s.UseLatestFormSchemaVersion
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) SetAppType(v string) *BatchUpdateFormDataByInstanceIdShrinkRequest {
	s.AppType = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) SetAsynchronousExecution(v bool) *BatchUpdateFormDataByInstanceIdShrinkRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) SetFormInstanceIdListShrink(v string) *BatchUpdateFormDataByInstanceIdShrinkRequest {
	s.FormInstanceIdListShrink = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) SetFormUuid(v string) *BatchUpdateFormDataByInstanceIdShrinkRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) SetIgnoreEmpty(v bool) *BatchUpdateFormDataByInstanceIdShrinkRequest {
	s.IgnoreEmpty = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) SetNoExecuteExpression(v bool) *BatchUpdateFormDataByInstanceIdShrinkRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) SetSystemToken(v string) *BatchUpdateFormDataByInstanceIdShrinkRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) SetUpdateFormDataJson(v string) *BatchUpdateFormDataByInstanceIdShrinkRequest {
	s.UpdateFormDataJson = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) SetUseLatestFormSchemaVersion(v bool) *BatchUpdateFormDataByInstanceIdShrinkRequest {
	s.UseLatestFormSchemaVersion = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
