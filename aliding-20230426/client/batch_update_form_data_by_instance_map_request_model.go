// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *BatchUpdateFormDataByInstanceMapRequest
	GetAppType() *string
	SetAsynchronousExecution(v bool) *BatchUpdateFormDataByInstanceMapRequest
	GetAsynchronousExecution() *bool
	SetFormUuid(v string) *BatchUpdateFormDataByInstanceMapRequest
	GetFormUuid() *string
	SetIgnoreEmpty(v bool) *BatchUpdateFormDataByInstanceMapRequest
	GetIgnoreEmpty() *bool
	SetNoExecuteExpression(v bool) *BatchUpdateFormDataByInstanceMapRequest
	GetNoExecuteExpression() *bool
	SetSystemToken(v string) *BatchUpdateFormDataByInstanceMapRequest
	GetSystemToken() *string
	SetUpdateFormDataJsonMap(v map[string]interface{}) *BatchUpdateFormDataByInstanceMapRequest
	GetUpdateFormDataJsonMap() map[string]interface{}
	SetUseLatestFormSchemaVersion(v bool) *BatchUpdateFormDataByInstanceMapRequest
	GetUseLatestFormSchemaVersion() *bool
}

type BatchUpdateFormDataByInstanceMapRequest struct {
	// example:
	//
	// String
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// false
	AsynchronousExecution *bool `json:"AsynchronousExecution,omitempty" xml:"AsynchronousExecution,omitempty"`
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// false
	IgnoreEmpty *bool `json:"IgnoreEmpty,omitempty" xml:"IgnoreEmpty,omitempty"`
	// example:
	//
	// false
	NoExecuteExpression *bool `json:"NoExecuteExpression,omitempty" xml:"NoExecuteExpression,omitempty"`
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken           *string                `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	UpdateFormDataJsonMap map[string]interface{} `json:"UpdateFormDataJsonMap,omitempty" xml:"UpdateFormDataJsonMap,omitempty"`
	// example:
	//
	// false
	UseLatestFormSchemaVersion *bool `json:"UseLatestFormSchemaVersion,omitempty" xml:"UseLatestFormSchemaVersion,omitempty"`
}

func (s BatchUpdateFormDataByInstanceMapRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapRequest) GetAppType() *string {
	return s.AppType
}

func (s *BatchUpdateFormDataByInstanceMapRequest) GetAsynchronousExecution() *bool {
	return s.AsynchronousExecution
}

func (s *BatchUpdateFormDataByInstanceMapRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchUpdateFormDataByInstanceMapRequest) GetIgnoreEmpty() *bool {
	return s.IgnoreEmpty
}

func (s *BatchUpdateFormDataByInstanceMapRequest) GetNoExecuteExpression() *bool {
	return s.NoExecuteExpression
}

func (s *BatchUpdateFormDataByInstanceMapRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *BatchUpdateFormDataByInstanceMapRequest) GetUpdateFormDataJsonMap() map[string]interface{} {
	return s.UpdateFormDataJsonMap
}

func (s *BatchUpdateFormDataByInstanceMapRequest) GetUseLatestFormSchemaVersion() *bool {
	return s.UseLatestFormSchemaVersion
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetAppType(v string) *BatchUpdateFormDataByInstanceMapRequest {
	s.AppType = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetAsynchronousExecution(v bool) *BatchUpdateFormDataByInstanceMapRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetFormUuid(v string) *BatchUpdateFormDataByInstanceMapRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetIgnoreEmpty(v bool) *BatchUpdateFormDataByInstanceMapRequest {
	s.IgnoreEmpty = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetNoExecuteExpression(v bool) *BatchUpdateFormDataByInstanceMapRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetSystemToken(v string) *BatchUpdateFormDataByInstanceMapRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetUpdateFormDataJsonMap(v map[string]interface{}) *BatchUpdateFormDataByInstanceMapRequest {
	s.UpdateFormDataJsonMap = v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetUseLatestFormSchemaVersion(v bool) *BatchUpdateFormDataByInstanceMapRequest {
	s.UseLatestFormSchemaVersion = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) Validate() error {
	return dara.Validate(s)
}
