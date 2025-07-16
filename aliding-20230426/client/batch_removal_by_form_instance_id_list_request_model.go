// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemovalByFormInstanceIdListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *BatchRemovalByFormInstanceIdListRequest
	GetAppType() *string
	SetAsynchronousExecution(v bool) *BatchRemovalByFormInstanceIdListRequest
	GetAsynchronousExecution() *bool
	SetExecuteExpression(v bool) *BatchRemovalByFormInstanceIdListRequest
	GetExecuteExpression() *bool
	SetFormInstanceIdList(v []*string) *BatchRemovalByFormInstanceIdListRequest
	GetFormInstanceIdList() []*string
	SetFormUuid(v string) *BatchRemovalByFormInstanceIdListRequest
	GetFormUuid() *string
	SetSystemToken(v string) *BatchRemovalByFormInstanceIdListRequest
	GetSystemToken() *string
}

type BatchRemovalByFormInstanceIdListRequest struct {
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
	FormInstanceIdList []*string `json:"FormInstanceIdList,omitempty" xml:"FormInstanceIdList,omitempty" type:"Repeated"`
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

func (s BatchRemovalByFormInstanceIdListRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRemovalByFormInstanceIdListRequest) GoString() string {
	return s.String()
}

func (s *BatchRemovalByFormInstanceIdListRequest) GetAppType() *string {
	return s.AppType
}

func (s *BatchRemovalByFormInstanceIdListRequest) GetAsynchronousExecution() *bool {
	return s.AsynchronousExecution
}

func (s *BatchRemovalByFormInstanceIdListRequest) GetExecuteExpression() *bool {
	return s.ExecuteExpression
}

func (s *BatchRemovalByFormInstanceIdListRequest) GetFormInstanceIdList() []*string {
	return s.FormInstanceIdList
}

func (s *BatchRemovalByFormInstanceIdListRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchRemovalByFormInstanceIdListRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetAppType(v string) *BatchRemovalByFormInstanceIdListRequest {
	s.AppType = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetAsynchronousExecution(v bool) *BatchRemovalByFormInstanceIdListRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetExecuteExpression(v bool) *BatchRemovalByFormInstanceIdListRequest {
	s.ExecuteExpression = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetFormInstanceIdList(v []*string) *BatchRemovalByFormInstanceIdListRequest {
	s.FormInstanceIdList = v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetFormUuid(v string) *BatchRemovalByFormInstanceIdListRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetSystemToken(v string) *BatchRemovalByFormInstanceIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) Validate() error {
	return dara.Validate(s)
}
