// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountCheckPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccountCheckPolicyResponseBody
	GetRequestId() *string
}

type ModifyAccountCheckPolicyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 866F5EB8-4650-4061-87F0-379F6F968BCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountCheckPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountCheckPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountCheckPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountCheckPolicyResponseBody) SetRequestId(v string) *ModifyAccountCheckPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountCheckPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
