// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenterPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCenterPolicyResponseBody
	GetRequestId() *string
}

type ModifyCenterPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5CED7F18-43B1-5035-BBB6-2538B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCenterPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCenterPolicyResponseBody) SetRequestId(v string) *ModifyCenterPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCenterPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
