// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAlertRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAlertRuleResponseBody
	GetSuccess() *bool
}

type DeleteAlertRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8754EE08-4AA2-5F77-ADD7-754DBBDA9F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAlertRuleResponseBody) SetRequestId(v string) *DeleteAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertRuleResponseBody) SetSuccess(v bool) *DeleteAlertRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
