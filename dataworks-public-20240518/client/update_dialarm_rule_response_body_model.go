// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIAlarmRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDIAlarmRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDIAlarmRuleResponseBody
	GetSuccess() *bool
}

type UpdateDIAlarmRuleResponseBody struct {
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// A6C6B486-E3A2-5D52-9E76-D9380485D946
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDIAlarmRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDIAlarmRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDIAlarmRuleResponseBody) SetRequestId(v string) *UpdateDIAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDIAlarmRuleResponseBody) SetSuccess(v bool) *UpdateDIAlarmRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDIAlarmRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
