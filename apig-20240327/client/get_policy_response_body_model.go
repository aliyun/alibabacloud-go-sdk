// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPolicyResponseBody
	GetCode() *string
	SetData(v *PolicyDetailInfo) *GetPolicyResponseBody
	GetData() *PolicyDetailInfo
	SetMessage(v string) *GetPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPolicyResponseBody
	GetRequestId() *string
}

type GetPolicyResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *PolicyDetailInfo `json:"data,omitempty" xml:"data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F46B9E7-67EF-5C8A-BA52-D38D5B32A***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPolicyResponseBody) GetData() *PolicyDetailInfo {
	return s.Data
}

func (s *GetPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyResponseBody) SetCode(v string) *GetPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *GetPolicyResponseBody) SetData(v *PolicyDetailInfo) *GetPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetPolicyResponseBody) SetMessage(v string) *GetPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *GetPolicyResponseBody) SetRequestId(v string) *GetPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
