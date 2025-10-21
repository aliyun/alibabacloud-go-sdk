// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePolicyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreatePolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreatePolicyResponseBody
	GetMessage() *string
	SetPolicyId(v int64) *CreatePolicyResponseBody
	GetPolicyId() *int64
	SetPolicyIdentifier(v string) *CreatePolicyResponseBody
	GetPolicyIdentifier() *string
	SetRequestId(v string) *CreatePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePolicyResponseBody
	GetSuccess() *bool
}

type CreatePolicyResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// If there is an error, returns the error message.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// 16
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Policy identifier
	//
	// example:
	//
	// 2tehcwesh4xx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// Request ID
	//
	// example:
	//
	// 74D2A967-2CE0-519B-B623-38D851734EC5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation was successful. true indicates success, false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreatePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePolicyResponseBody) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *CreatePolicyResponseBody) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *CreatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePolicyResponseBody) SetCode(v string) *CreatePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolicyResponseBody) SetHttpStatusCode(v int32) *CreatePolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePolicyResponseBody) SetMessage(v string) *CreatePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePolicyResponseBody) SetPolicyId(v int64) *CreatePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreatePolicyResponseBody) SetPolicyIdentifier(v string) *CreatePolicyResponseBody {
	s.PolicyIdentifier = &v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponseBody) SetSuccess(v bool) *CreatePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
