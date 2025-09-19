// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSasContainerWebDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSasContainerWebDefenseRuleResponseBody
	GetRequestId() *string
}

type DeleteSasContainerWebDefenseRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DDCAE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSasContainerWebDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSasContainerWebDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSasContainerWebDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSasContainerWebDefenseRuleResponseBody) SetRequestId(v string) *DeleteSasContainerWebDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSasContainerWebDefenseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
