// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteAlertRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteAlertRuleResponseBody
	GetRequestId() *string
}

type DeleteAlertRuleResponseBody struct {
	// Indicates whether the call was successful.
	//
	// 	- `true`: The call was successful.
	//
	// 	- `false`: The call failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertRuleResponseBody) SetIsSuccess(v bool) *DeleteAlertRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertRuleResponseBody) SetRequestId(v string) *DeleteAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
