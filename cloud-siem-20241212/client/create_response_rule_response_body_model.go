// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResponseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateResponseRuleResponseBody
	GetRequestId() *string
	SetResponseRuleId(v string) *CreateResponseRuleResponseBody
	GetResponseRuleId() *string
}

type CreateResponseRuleResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 399827
	ResponseRuleId *string `json:"ResponseRuleId,omitempty" xml:"ResponseRuleId,omitempty"`
}

func (s CreateResponseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResponseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResponseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResponseRuleResponseBody) GetResponseRuleId() *string {
	return s.ResponseRuleId
}

func (s *CreateResponseRuleResponseBody) SetRequestId(v string) *CreateResponseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResponseRuleResponseBody) SetResponseRuleId(v string) *CreateResponseRuleResponseBody {
	s.ResponseRuleId = &v
	return s
}

func (s *CreateResponseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
