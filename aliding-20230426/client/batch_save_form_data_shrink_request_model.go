// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSaveFormDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *BatchSaveFormDataShrinkRequest
	GetAppType() *string
	SetAsynchronousExecution(v bool) *BatchSaveFormDataShrinkRequest
	GetAsynchronousExecution() *bool
	SetFormDataJsonListShrink(v string) *BatchSaveFormDataShrinkRequest
	GetFormDataJsonListShrink() *string
	SetFormUuid(v string) *BatchSaveFormDataShrinkRequest
	GetFormUuid() *string
	SetKeepRunningAfterException(v bool) *BatchSaveFormDataShrinkRequest
	GetKeepRunningAfterException() *bool
	SetNoExecuteExpression(v bool) *BatchSaveFormDataShrinkRequest
	GetNoExecuteExpression() *bool
	SetSystemToken(v string) *BatchSaveFormDataShrinkRequest
	GetSystemToken() *string
}

type BatchSaveFormDataShrinkRequest struct {
	// example:
	//
	// APP_XCE0EVXS6DYG3YDYC5RD
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// false
	AsynchronousExecution  *bool   `json:"AsynchronousExecution,omitempty" xml:"AsynchronousExecution,omitempty"`
	FormDataJsonListShrink *string `json:"FormDataJsonList,omitempty" xml:"FormDataJsonList,omitempty"`
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// false
	KeepRunningAfterException *bool `json:"KeepRunningAfterException,omitempty" xml:"KeepRunningAfterException,omitempty"`
	// example:
	//
	// false
	NoExecuteExpression *bool `json:"NoExecuteExpression,omitempty" xml:"NoExecuteExpression,omitempty"`
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s BatchSaveFormDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSaveFormDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *BatchSaveFormDataShrinkRequest) GetAsynchronousExecution() *bool {
	return s.AsynchronousExecution
}

func (s *BatchSaveFormDataShrinkRequest) GetFormDataJsonListShrink() *string {
	return s.FormDataJsonListShrink
}

func (s *BatchSaveFormDataShrinkRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchSaveFormDataShrinkRequest) GetKeepRunningAfterException() *bool {
	return s.KeepRunningAfterException
}

func (s *BatchSaveFormDataShrinkRequest) GetNoExecuteExpression() *bool {
	return s.NoExecuteExpression
}

func (s *BatchSaveFormDataShrinkRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *BatchSaveFormDataShrinkRequest) SetAppType(v string) *BatchSaveFormDataShrinkRequest {
	s.AppType = &v
	return s
}

func (s *BatchSaveFormDataShrinkRequest) SetAsynchronousExecution(v bool) *BatchSaveFormDataShrinkRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchSaveFormDataShrinkRequest) SetFormDataJsonListShrink(v string) *BatchSaveFormDataShrinkRequest {
	s.FormDataJsonListShrink = &v
	return s
}

func (s *BatchSaveFormDataShrinkRequest) SetFormUuid(v string) *BatchSaveFormDataShrinkRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchSaveFormDataShrinkRequest) SetKeepRunningAfterException(v bool) *BatchSaveFormDataShrinkRequest {
	s.KeepRunningAfterException = &v
	return s
}

func (s *BatchSaveFormDataShrinkRequest) SetNoExecuteExpression(v bool) *BatchSaveFormDataShrinkRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *BatchSaveFormDataShrinkRequest) SetSystemToken(v string) *BatchSaveFormDataShrinkRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchSaveFormDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
