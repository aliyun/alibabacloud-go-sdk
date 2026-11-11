// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGlobalCloudAccountSystemRolePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyDocument(v string) *GetGlobalCloudAccountSystemRolePolicyResponseBody
	GetPolicyDocument() *string
	SetRequestId(v string) *GetGlobalCloudAccountSystemRolePolicyResponseBody
	GetRequestId() *string
}

type GetGlobalCloudAccountSystemRolePolicyResponseBody struct {
	// The access policy document for the global cloud account system role (JSON character string).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Version":"1","Statement":[]}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The unique identifier of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGlobalCloudAccountSystemRolePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGlobalCloudAccountSystemRolePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponseBody) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponseBody) SetPolicyDocument(v string) *GetGlobalCloudAccountSystemRolePolicyResponseBody {
	s.PolicyDocument = &v
	return s
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponseBody) SetRequestId(v string) *GetGlobalCloudAccountSystemRolePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
