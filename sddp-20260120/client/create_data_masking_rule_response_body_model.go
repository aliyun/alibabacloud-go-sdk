// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataMaskingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDataMaskingRuleResponseBody
	GetRequestId() *string
}

type CreateDataMaskingRuleResponseBody struct {
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataMaskingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataMaskingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataMaskingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataMaskingRuleResponseBody) SetRequestId(v string) *CreateDataMaskingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataMaskingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
