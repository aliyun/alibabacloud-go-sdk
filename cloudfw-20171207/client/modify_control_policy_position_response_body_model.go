// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlPolicyPositionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyControlPolicyPositionResponseBody
	GetRequestId() *string
}

type ModifyControlPolicyPositionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125EEB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyControlPolicyPositionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlPolicyPositionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPositionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyControlPolicyPositionResponseBody) SetRequestId(v string) *ModifyControlPolicyPositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyControlPolicyPositionResponseBody) Validate() error {
	return dara.Validate(s)
}
