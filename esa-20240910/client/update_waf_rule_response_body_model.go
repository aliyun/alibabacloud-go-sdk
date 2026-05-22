// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWafRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateWafRuleResponseBody
	GetId() *int64
	SetRequestId(v string) *UpdateWafRuleResponseBody
	GetRequestId() *string
}

type UpdateWafRuleResponseBody struct {
	// WAF rule ID, which can be obtained by calling the [ListWafRules](https://help.aliyun.com/document_detail/2878257.html) interface.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWafRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWafRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWafRuleResponseBody) GetId() *int64 {
	return s.Id
}

func (s *UpdateWafRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWafRuleResponseBody) SetId(v int64) *UpdateWafRuleResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateWafRuleResponseBody) SetRequestId(v string) *UpdateWafRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWafRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
