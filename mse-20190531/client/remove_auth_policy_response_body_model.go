// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAuthPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RemoveAuthPolicyResponseBody
	GetCode() *int32
	SetData(v string) *RemoveAuthPolicyResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *RemoveAuthPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveAuthPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveAuthPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveAuthPolicyResponseBody
	GetSuccess() *bool
}

type RemoveAuthPolicyResponseBody struct {
	// example:
	//
	// 500
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9515ACA4-E94D-440D-989E-C379FCED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveAuthPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveAuthPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAuthPolicyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RemoveAuthPolicyResponseBody) GetData() *string {
	return s.Data
}

func (s *RemoveAuthPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveAuthPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveAuthPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveAuthPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveAuthPolicyResponseBody) SetCode(v int32) *RemoveAuthPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveAuthPolicyResponseBody) SetData(v string) *RemoveAuthPolicyResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveAuthPolicyResponseBody) SetHttpStatusCode(v int32) *RemoveAuthPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveAuthPolicyResponseBody) SetMessage(v string) *RemoveAuthPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveAuthPolicyResponseBody) SetRequestId(v string) *RemoveAuthPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAuthPolicyResponseBody) SetSuccess(v bool) *RemoveAuthPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveAuthPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
