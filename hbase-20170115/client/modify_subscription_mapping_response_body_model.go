// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySubscriptionMappingResponseBody
	GetRequestId() *string
}

type ModifySubscriptionMappingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySubscriptionMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionMappingResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySubscriptionMappingResponseBody) SetRequestId(v string) *ModifySubscriptionMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySubscriptionMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
