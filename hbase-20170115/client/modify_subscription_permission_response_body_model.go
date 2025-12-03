// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySubscriptionPermissionResponseBody
	GetRequestId() *string
	SetStatus(v string) *ModifySubscriptionPermissionResponseBody
	GetStatus() *string
}

type ModifySubscriptionPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifySubscriptionPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySubscriptionPermissionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ModifySubscriptionPermissionResponseBody) SetRequestId(v string) *ModifySubscriptionPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySubscriptionPermissionResponseBody) SetStatus(v string) *ModifySubscriptionPermissionResponseBody {
	s.Status = &v
	return s
}

func (s *ModifySubscriptionPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
