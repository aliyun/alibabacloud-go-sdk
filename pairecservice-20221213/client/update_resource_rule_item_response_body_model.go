// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRuleItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceRuleItemResponseBody
	GetRequestId() *string
}

type UpdateResourceRuleItemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceRuleItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRuleItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceRuleItemResponseBody) SetRequestId(v string) *UpdateResourceRuleItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceRuleItemResponseBody) Validate() error {
	return dara.Validate(s)
}
