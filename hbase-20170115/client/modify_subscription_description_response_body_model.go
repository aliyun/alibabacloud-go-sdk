// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySubscriptionDescriptionResponseBody
	GetRequestId() *string
}

type ModifySubscriptionDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySubscriptionDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySubscriptionDescriptionResponseBody) SetRequestId(v string) *ModifySubscriptionDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySubscriptionDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
