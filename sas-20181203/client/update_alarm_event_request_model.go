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
}

type UpdateAlarmEventRequest struct {
	// The IDs of the alert events.
	AlarmEventIdList []*int64 `json:"AlarmEventIdList,omitempty" xml:"AlarmEventIdList,omitempty" type:"Repeated"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The operation that you want to perform on the alert events. Valid values:
	//
	// 	- **manual_handled**: handle the alert events.
	//
	// 	- **ignore**: igore the alert events.
	//
	// 	- **cancel_ignore**: remove the alert events from the whitelist.
	//
	// example:
	//
	// ignore
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
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

func (s *UpdateAlarmEventRequest) Validate() error {
	return dara.Validate(s)
}
