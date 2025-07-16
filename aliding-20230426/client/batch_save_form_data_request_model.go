// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSaveFormDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *BatchSaveFormDataRequest
	GetAppType() *string
	SetAsynchronousExecution(v bool) *BatchSaveFormDataRequest
	GetAsynchronousExecution() *bool
	SetFormDataJsonList(v []*string) *BatchSaveFormDataRequest
	GetFormDataJsonList() []*string
	SetFormUuid(v string) *BatchSaveFormDataRequest
	GetFormUuid() *string
	SetKeepRunningAfterException(v bool) *BatchSaveFormDataRequest
	GetKeepRunningAfterException() *bool
	SetNoExecuteExpression(v bool) *BatchSaveFormDataRequest
	GetNoExecuteExpression() *bool
	SetSystemToken(v string) *BatchSaveFormDataRequest
	GetSystemToken() *string
}

type BatchSaveFormDataRequest struct {
	// example:
	//
	// APP_XCE0EVXS6DYG3YDYC5RD
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// false
	AsynchronousExecution *bool     `json:"AsynchronousExecution,omitempty" xml:"AsynchronousExecution,omitempty"`
	FormDataJsonList      []*string `json:"FormDataJsonList,omitempty" xml:"FormDataJsonList,omitempty" type:"Repeated"`
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

func (s BatchSaveFormDataRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSaveFormDataRequest) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataRequest) GetAppType() *string {
	return s.AppType
}

func (s *BatchSaveFormDataRequest) GetAsynchronousExecution() *bool {
	return s.AsynchronousExecution
}

func (s *BatchSaveFormDataRequest) GetFormDataJsonList() []*string {
	return s.FormDataJsonList
}

func (s *BatchSaveFormDataRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchSaveFormDataRequest) GetKeepRunningAfterException() *bool {
	return s.KeepRunningAfterException
}

func (s *BatchSaveFormDataRequest) GetNoExecuteExpression() *bool {
	return s.NoExecuteExpression
}

func (s *BatchSaveFormDataRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *BatchSaveFormDataRequest) SetAppType(v string) *BatchSaveFormDataRequest {
	s.AppType = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetAsynchronousExecution(v bool) *BatchSaveFormDataRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetFormDataJsonList(v []*string) *BatchSaveFormDataRequest {
	s.FormDataJsonList = v
	return s
}

func (s *BatchSaveFormDataRequest) SetFormUuid(v string) *BatchSaveFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetKeepRunningAfterException(v bool) *BatchSaveFormDataRequest {
	s.KeepRunningAfterException = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetNoExecuteExpression(v bool) *BatchSaveFormDataRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetSystemToken(v string) *BatchSaveFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchSaveFormDataRequest) Validate() error {
	return dara.Validate(s)
}
