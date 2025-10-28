// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCheckPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicys(v []*DeleteCheckPolicyResponseBodyPolicys) *DeleteCheckPolicyResponseBody
	GetPolicys() []*DeleteCheckPolicyResponseBodyPolicys
	SetRequestId(v string) *DeleteCheckPolicyResponseBody
	GetRequestId() *string
}

type DeleteCheckPolicyResponseBody struct {
	// List of deleted policy details.
	Policys []*DeleteCheckPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The unique ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCheckPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCheckPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCheckPolicyResponseBody) GetPolicys() []*DeleteCheckPolicyResponseBodyPolicys {
	return s.Policys
}

func (s *DeleteCheckPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCheckPolicyResponseBody) SetPolicys(v []*DeleteCheckPolicyResponseBodyPolicys) *DeleteCheckPolicyResponseBody {
	s.Policys = v
	return s
}

func (s *DeleteCheckPolicyResponseBody) SetRequestId(v string) *DeleteCheckPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCheckPolicyResponseBody) Validate() error {
	if s.Policys != nil {
		for _, item := range s.Policys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteCheckPolicyResponseBodyPolicys struct {
	// Deleted policy ID.
	//
	// example:
	//
	// 1000000000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the custom policy.
	//
	// example:
	//
	// testPolicyName
	PolicyShowName *string `json:"PolicyShowName,omitempty" xml:"PolicyShowName,omitempty"`
	// Policy type for custom check rule:
	//
	// 	- **STANDARD**: Standard-level policy
	//
	// 	- **REQUIREMENT**: Requirement-level policy
	//
	// 	- **SECTION**: Section-level policy
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DeleteCheckPolicyResponseBodyPolicys) String() string {
	return dara.Prettify(s)
}

func (s DeleteCheckPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DeleteCheckPolicyResponseBodyPolicys) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DeleteCheckPolicyResponseBodyPolicys) GetPolicyShowName() *string {
	return s.PolicyShowName
}

func (s *DeleteCheckPolicyResponseBodyPolicys) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteCheckPolicyResponseBodyPolicys) SetPolicyId(v int64) *DeleteCheckPolicyResponseBodyPolicys {
	s.PolicyId = &v
	return s
}

func (s *DeleteCheckPolicyResponseBodyPolicys) SetPolicyShowName(v string) *DeleteCheckPolicyResponseBodyPolicys {
	s.PolicyShowName = &v
	return s
}

func (s *DeleteCheckPolicyResponseBodyPolicys) SetPolicyType(v string) *DeleteCheckPolicyResponseBodyPolicys {
	s.PolicyType = &v
	return s
}

func (s *DeleteCheckPolicyResponseBodyPolicys) Validate() error {
	return dara.Validate(s)
}
