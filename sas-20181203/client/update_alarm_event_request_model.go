// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlarmEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmEventIdList(v []*int64) *UpdateAlarmEventRequest
	GetAlarmEventIdList() []*int64
	SetLang(v string) *UpdateAlarmEventRequest
	GetLang() *string
	SetOperationCode(v string) *UpdateAlarmEventRequest
	GetOperationCode() *string
	SetResourceDirectoryAccountId(v int64) *UpdateAlarmEventRequest
	GetResourceDirectoryAccountId() *int64
}

type UpdateAlarmEventRequest struct {
	// The list of alert event IDs.
	//
	// You can call ListHoneypotAlarmEvents to obtain valid alert event IDs. Before calling this operation, ensure that honeypots are deployed and honeypot alert event data exists.
	//
	// This parameter is required. If this parameter is not specified, the API returns HTTP 400 with the error code IllegalParam.
	AlarmEventIdList []*int64 `json:"AlarmEventIdList,omitempty" xml:"AlarmEventIdList,omitempty" type:"Repeated"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The processing method for the alert event. Valid values:
	//
	// - **manual_handled**: Handled.
	//
	// - **ignore**: Ignored.
	//
	// - **cancel_ignore**: Removed from the whitelist.
	//
	// This parameter is required. If this parameter is not specified, the API returns HTTP 400 with the error code IllegalParam.
	//
	// example:
	//
	// ignore
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The Alibaba Cloud account ID of the member accounts in the resource folder.
	//
	// >You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s UpdateAlarmEventRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmEventRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlarmEventRequest) GetAlarmEventIdList() []*int64 {
	return s.AlarmEventIdList
}

func (s *UpdateAlarmEventRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAlarmEventRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *UpdateAlarmEventRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *UpdateAlarmEventRequest) SetAlarmEventIdList(v []*int64) *UpdateAlarmEventRequest {
	s.AlarmEventIdList = v
	return s
}

func (s *UpdateAlarmEventRequest) SetLang(v string) *UpdateAlarmEventRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAlarmEventRequest) SetOperationCode(v string) *UpdateAlarmEventRequest {
	s.OperationCode = &v
	return s
}

func (s *UpdateAlarmEventRequest) SetResourceDirectoryAccountId(v int64) *UpdateAlarmEventRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *UpdateAlarmEventRequest) Validate() error {
	return dara.Validate(s)
}
