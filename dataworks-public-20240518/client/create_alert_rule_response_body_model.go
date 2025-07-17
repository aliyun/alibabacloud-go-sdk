// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateAlertRuleResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateAlertRuleResponseBody
	GetRequestId() *string
}

type CreateAlertRuleResponseBody struct {
	// The rule ID.
	//
	// example:
	//
	// 123123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A6C6B486-E3A2-5D52-9E76-D9380485D946
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlertRuleResponseBody) SetId(v int64) *CreateAlertRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateAlertRuleResponseBody) SetRequestId(v string) *CreateAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
