// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTLSCipherPolicyAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetTLSCipherPolicyAttributeResponseBody
	GetRequestId() *string
	SetTaskId(v string) *SetTLSCipherPolicyAttributeResponseBody
	GetTaskId() *string
}

type SetTLSCipherPolicyAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af****-18f6aed5
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SetTLSCipherPolicyAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetTLSCipherPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetTLSCipherPolicyAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetTLSCipherPolicyAttributeResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *SetTLSCipherPolicyAttributeResponseBody) SetRequestId(v string) *SetTLSCipherPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeResponseBody) SetTaskId(v string) *SetTLSCipherPolicyAttributeResponseBody {
	s.TaskId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
