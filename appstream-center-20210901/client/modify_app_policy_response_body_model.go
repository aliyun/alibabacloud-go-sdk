// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppPolicyResponseBody
	GetRequestId() *string
}

type ModifyAppPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppPolicyResponseBody) SetRequestId(v string) *ModifyAppPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
