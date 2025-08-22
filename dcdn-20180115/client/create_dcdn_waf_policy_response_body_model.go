// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnWafPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *CreateDcdnWafPolicyResponseBody
	GetPolicyId() *int64
	SetRequestId(v string) *CreateDcdnWafPolicyResponseBody
	GetRequestId() *string
}

type CreateDcdnWafPolicyResponseBody struct {
	// The ID of the protection policy that you created.
	//
	// example:
	//
	// 10000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDcdnWafPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnWafPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDcdnWafPolicyResponseBody) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *CreateDcdnWafPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDcdnWafPolicyResponseBody) SetPolicyId(v int64) *CreateDcdnWafPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateDcdnWafPolicyResponseBody) SetRequestId(v string) *CreateDcdnWafPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDcdnWafPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
