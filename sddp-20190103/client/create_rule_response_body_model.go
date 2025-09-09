// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int32) *CreateRuleResponseBody
	GetId() *int32
	SetRequestId(v string) *CreateRuleResponseBody
	GetRequestId() *string
}

type CreateRuleResponseBody struct {
	// The unique ID of the sensitive data detection rule.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) GetId() *int32 {
	return s.Id
}

func (s *CreateRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRuleResponseBody) SetId(v int32) *CreateRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
