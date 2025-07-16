// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceMapShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *BatchUpdateFormDataByInstanceMapShrinkRequest
	GetAppType() *string
	SetAsynchronousExecution(v bool) *BatchUpdateFormDataByInstanceMapShrinkRequest
	GetAsynchronousExecution() *bool
	SetFormUuid(v string) *BatchUpdateFormDataByInstanceMapShrinkRequest
	GetFormUuid() *string
	SetIgnoreEmpty(v bool) *BatchUpdateFormDataByInstanceMapShrinkRequest
	GetIgnoreEmpty() *bool
	SetNoExecuteExpression(v bool) *BatchUpdateFormDataByInstanceMapShrinkRequest
	GetNoExecuteExpression() *bool
	SetSystemToken(v string) *BatchUpdateFormDataByInstanceMapShrinkRequest
	GetSystemToken() *string
	SetUpdateFormDataJsonMapShrink(v string) *BatchUpdateFormDataByInstanceMapShrinkRequest
	GetUpdateFormDataJsonMapShrink() *string
	SetUseLatestFormSchemaVersion(v bool) *BatchUpdateFormDataByInstanceMapShrinkRequest
	GetUseLatestFormSchemaVersion() *bool
}

type BatchUpdateFormDataByInstanceMapShrinkRequest struct {
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
	SystemToken                 *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	UpdateFormDataJsonMapShrink *string `json:"UpdateFormDataJsonMap,omitempty" xml:"UpdateFormDataJsonMap,omitempty"`
	// example:
	//
	// false
	UseLatestFormSchemaVersion *bool `json:"UseLatestFormSchemaVersion,omitempty" xml:"UseLatestFormSchemaVersion,omitempty"`
}

func (s BatchUpdateFormDataByInstanceMapShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) GetAsynchronousExecution() *bool {
	return s.AsynchronousExecution
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) GetIgnoreEmpty() *bool {
	return s.IgnoreEmpty
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) GetNoExecuteExpression() *bool {
	return s.NoExecuteExpression
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) GetUpdateFormDataJsonMapShrink() *string {
	return s.UpdateFormDataJsonMapShrink
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) GetUseLatestFormSchemaVersion() *bool {
	return s.UseLatestFormSchemaVersion
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) SetAppType(v string) *BatchUpdateFormDataByInstanceMapShrinkRequest {
	s.AppType = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) SetAsynchronousExecution(v bool) *BatchUpdateFormDataByInstanceMapShrinkRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) SetFormUuid(v string) *BatchUpdateFormDataByInstanceMapShrinkRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) SetIgnoreEmpty(v bool) *BatchUpdateFormDataByInstanceMapShrinkRequest {
	s.IgnoreEmpty = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) SetNoExecuteExpression(v bool) *BatchUpdateFormDataByInstanceMapShrinkRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) SetSystemToken(v string) *BatchUpdateFormDataByInstanceMapShrinkRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) SetUpdateFormDataJsonMapShrink(v string) *BatchUpdateFormDataByInstanceMapShrinkRequest {
	s.UpdateFormDataJsonMapShrink = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) SetUseLatestFormSchemaVersion(v bool) *BatchUpdateFormDataByInstanceMapShrinkRequest {
	s.UseLatestFormSchemaVersion = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapShrinkRequest) Validate() error {
	return dara.Validate(s)
}
