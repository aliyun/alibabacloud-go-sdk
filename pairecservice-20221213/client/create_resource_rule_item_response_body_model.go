// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRuleItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateResourceRuleItemResponseBody
	GetRequestId() *string
	SetResourceRuleItemId(v string) *CreateResourceRuleItemResponseBody
	GetResourceRuleItemId() *string
}

type CreateResourceRuleItemResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRuleItemId *string `json:"ResourceRuleItemId,omitempty" xml:"ResourceRuleItemId,omitempty"`
}

func (s CreateResourceRuleItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRuleItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceRuleItemResponseBody) GetResourceRuleItemId() *string {
	return s.ResourceRuleItemId
}

func (s *CreateResourceRuleItemResponseBody) SetRequestId(v string) *CreateResourceRuleItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceRuleItemResponseBody) SetResourceRuleItemId(v string) *CreateResourceRuleItemResponseBody {
	s.ResourceRuleItemId = &v
	return s
}

func (s *CreateResourceRuleItemResponseBody) Validate() error {
	return dara.Validate(s)
}
