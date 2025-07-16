// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemovalByFormInstanceIdListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *BatchRemovalByFormInstanceIdListShrinkRequest
	GetAppType() *string
	SetAsynchronousExecution(v bool) *BatchRemovalByFormInstanceIdListShrinkRequest
	GetAsynchronousExecution() *bool
	SetExecuteExpression(v bool) *BatchRemovalByFormInstanceIdListShrinkRequest
	GetExecuteExpression() *bool
	SetFormInstanceIdListShrink(v string) *BatchRemovalByFormInstanceIdListShrinkRequest
	GetFormInstanceIdListShrink() *string
	SetFormUuid(v string) *BatchRemovalByFormInstanceIdListShrinkRequest
	GetFormUuid() *string
	SetSystemToken(v string) *BatchRemovalByFormInstanceIdListShrinkRequest
	GetSystemToken() *string
}

type BatchRemovalByFormInstanceIdListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_XCE0EVXS6DYG3YDYC5RD
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// false
	AsynchronousExecution *bool `json:"AsynchronousExecution,omitempty" xml:"AsynchronousExecution,omitempty"`
	// example:
	//
	// false
	ExecuteExpression *bool `json:"ExecuteExpression,omitempty" xml:"ExecuteExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// []
	FormInstanceIdListShrink *string `json:"FormInstanceIdList,omitempty" xml:"FormInstanceIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s BatchRemovalByFormInstanceIdListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRemovalByFormInstanceIdListShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) GetAsynchronousExecution() *bool {
	return s.AsynchronousExecution
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) GetExecuteExpression() *bool {
	return s.ExecuteExpression
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) GetFormInstanceIdListShrink() *string {
	return s.FormInstanceIdListShrink
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) SetAppType(v string) *BatchRemovalByFormInstanceIdListShrinkRequest {
	s.AppType = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) SetAsynchronousExecution(v bool) *BatchRemovalByFormInstanceIdListShrinkRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) SetExecuteExpression(v bool) *BatchRemovalByFormInstanceIdListShrinkRequest {
	s.ExecuteExpression = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) SetFormInstanceIdListShrink(v string) *BatchRemovalByFormInstanceIdListShrinkRequest {
	s.FormInstanceIdListShrink = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) SetFormUuid(v string) *BatchRemovalByFormInstanceIdListShrinkRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) SetSystemToken(v string) *BatchRemovalByFormInstanceIdListShrinkRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
