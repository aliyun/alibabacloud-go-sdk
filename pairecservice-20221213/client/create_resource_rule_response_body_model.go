// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateResourceRuleResponseBody
	GetRequestId() *string
	SetResourceRuleId(v string) *CreateResourceRuleResponseBody
	GetResourceRuleId() *string
}

type CreateResourceRuleResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRuleId *string `json:"ResourceRuleId,omitempty" xml:"ResourceRuleId,omitempty"`
}

func (s CreateResourceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceRuleResponseBody) GetResourceRuleId() *string {
	return s.ResourceRuleId
}

func (s *CreateResourceRuleResponseBody) SetRequestId(v string) *CreateResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceRuleResponseBody) SetResourceRuleId(v string) *CreateResourceRuleResponseBody {
	s.ResourceRuleId = &v
	return s
}

func (s *CreateResourceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
