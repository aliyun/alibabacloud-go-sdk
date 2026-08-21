// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeletePrivateAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeletePrivateAccessPolicyResponseBody
	GetRequestId() *string
}

type BatchDeletePrivateAccessPolicyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 019F3F8C-1127-5152-80E0-4F9D45DB5756
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeletePrivateAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeletePrivateAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeletePrivateAccessPolicyResponseBody) SetRequestId(v string) *BatchDeletePrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeletePrivateAccessPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
