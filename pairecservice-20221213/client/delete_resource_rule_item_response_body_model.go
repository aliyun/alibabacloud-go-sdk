// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRuleItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceRuleItemResponseBody
	GetRequestId() *string
}

type DeleteResourceRuleItemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceRuleItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRuleItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceRuleItemResponseBody) SetRequestId(v string) *DeleteResourceRuleItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceRuleItemResponseBody) Validate() error {
	return dara.Validate(s)
}
