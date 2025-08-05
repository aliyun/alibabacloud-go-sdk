// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlPolicyPriorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyControlPolicyPriorityResponseBody
	GetRequestId() *string
}

type ModifyControlPolicyPriorityResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 586F34E8-3F16-4C08-9FFC-8FFDC64B9D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyControlPolicyPriorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlPolicyPriorityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPriorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyControlPolicyPriorityResponseBody) SetRequestId(v string) *ModifyControlPolicyPriorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyControlPolicyPriorityResponseBody) Validate() error {
	return dara.Validate(s)
}
