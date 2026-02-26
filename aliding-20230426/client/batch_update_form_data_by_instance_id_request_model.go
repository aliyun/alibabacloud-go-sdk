// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *BatchUpdateFormDataByInstanceIdRequest
	GetAppType() *string
	SetAsynchronousExecution(v bool) *BatchUpdateFormDataByInstanceIdRequest
	GetAsynchronousExecution() *bool
	SetFormInstanceIdList(v []*string) *BatchUpdateFormDataByInstanceIdRequest
	GetFormInstanceIdList() []*string
	SetFormUuid(v string) *BatchUpdateFormDataByInstanceIdRequest
	GetFormUuid() *string
	SetIgnoreEmpty(v bool) *BatchUpdateFormDataByInstanceIdRequest
	GetIgnoreEmpty() *bool
	SetNoExecuteExpression(v bool) *BatchUpdateFormDataByInstanceIdRequest
	GetNoExecuteExpression() *bool
	SetSystemToken(v string) *BatchUpdateFormDataByInstanceIdRequest
	GetSystemToken() *string
	SetUpdateFormDataJson(v string) *BatchUpdateFormDataByInstanceIdRequest
	GetUpdateFormDataJson() *string
	SetUseLatestFormSchemaVersion(v bool) *BatchUpdateFormDataByInstanceIdRequest
	GetUseLatestFormSchemaVersion() *bool
}

type BatchUpdateFormDataByInstanceIdRequest struct {
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
	FormInstanceIdList []*string `json:"FormInstanceIdList,omitempty" xml:"FormInstanceIdList,omitempty" type:"Repeated"`
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
	// {\\"countrySelectField_l0c1cwiu\\":[{\\"value\\":\\"US\\"}],\\"addressField_l0c1cwiy\\":{\\"address\\":\\"111\\",\\"regionIds\\":[460000,469027,469023401],\\"regionText\\":[{\\"en_US\\":\\"hai+nan+sheng\\",\\"zh_CN\\":\\"海南省\\"},{\\"en_US\\":\\"cheng+mai+xian\\",\\"zh_CN\\":\\"澄迈县\\"},{\\"en_US\\":\\"guo+ying+hong+gang+nong+chang\\",\\"zh_CN\\":\\"国营红岗农场\\"}]}}
	UpdateFormDataJson *string `json:"UpdateFormDataJson,omitempty" xml:"UpdateFormDataJson,omitempty"`
	// example:
	//
	// false
	UseLatestFormSchemaVersion *bool `json:"UseLatestFormSchemaVersion,omitempty" xml:"UseLatestFormSchemaVersion,omitempty"`
}

func (s BatchUpdateFormDataByInstanceIdRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdRequest) GetAppType() *string {
	return s.AppType
}

func (s *BatchUpdateFormDataByInstanceIdRequest) GetAsynchronousExecution() *bool {
	return s.AsynchronousExecution
}

func (s *BatchUpdateFormDataByInstanceIdRequest) GetFormInstanceIdList() []*string {
	return s.FormInstanceIdList
}

func (s *BatchUpdateFormDataByInstanceIdRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchUpdateFormDataByInstanceIdRequest) GetIgnoreEmpty() *bool {
	return s.IgnoreEmpty
}

func (s *BatchUpdateFormDataByInstanceIdRequest) GetNoExecuteExpression() *bool {
	return s.NoExecuteExpression
}

func (s *BatchUpdateFormDataByInstanceIdRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *BatchUpdateFormDataByInstanceIdRequest) GetUpdateFormDataJson() *string {
	return s.UpdateFormDataJson
}

func (s *BatchUpdateFormDataByInstanceIdRequest) GetUseLatestFormSchemaVersion() *bool {
	return s.UseLatestFormSchemaVersion
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetAppType(v string) *BatchUpdateFormDataByInstanceIdRequest {
	s.AppType = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetAsynchronousExecution(v bool) *BatchUpdateFormDataByInstanceIdRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetFormInstanceIdList(v []*string) *BatchUpdateFormDataByInstanceIdRequest {
	s.FormInstanceIdList = v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetFormUuid(v string) *BatchUpdateFormDataByInstanceIdRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetIgnoreEmpty(v bool) *BatchUpdateFormDataByInstanceIdRequest {
	s.IgnoreEmpty = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetNoExecuteExpression(v bool) *BatchUpdateFormDataByInstanceIdRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetSystemToken(v string) *BatchUpdateFormDataByInstanceIdRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetUpdateFormDataJson(v string) *BatchUpdateFormDataByInstanceIdRequest {
	s.UpdateFormDataJson = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetUseLatestFormSchemaVersion(v bool) *BatchUpdateFormDataByInstanceIdRequest {
	s.UseLatestFormSchemaVersion = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) Validate() error {
	return dara.Validate(s)
}
