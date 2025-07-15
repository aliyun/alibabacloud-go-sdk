// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceCenterPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModifyResults(v []*ModifyResourceCenterPolicyResponseBodyModifyResults) *ModifyResourceCenterPolicyResponseBody
	GetModifyResults() []*ModifyResourceCenterPolicyResponseBodyModifyResults
	SetRequestId(v string) *ModifyResourceCenterPolicyResponseBody
	GetRequestId() *string
}

type ModifyResourceCenterPolicyResponseBody struct {
	// The modification results.
	ModifyResults []*ModifyResourceCenterPolicyResponseBodyModifyResults `json:"ModifyResults,omitempty" xml:"ModifyResults,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 51592A88-0F2C-55E6-AD2C-2AD9C10D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyResourceCenterPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceCenterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceCenterPolicyResponseBody) GetModifyResults() []*ModifyResourceCenterPolicyResponseBodyModifyResults {
	return s.ModifyResults
}

func (s *ModifyResourceCenterPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyResourceCenterPolicyResponseBody) SetModifyResults(v []*ModifyResourceCenterPolicyResponseBodyModifyResults) *ModifyResourceCenterPolicyResponseBody {
	s.ModifyResults = v
	return s
}

func (s *ModifyResourceCenterPolicyResponseBody) SetRequestId(v string) *ModifyResourceCenterPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResourceCenterPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyResourceCenterPolicyResponseBodyModifyResults struct {
	// The verification result.
	//
	// example:
	//
	// true
	CheckResult *bool `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// ecd-e254cpyt9bb*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ModifyResourceCenterPolicyResponseBodyModifyResults) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceCenterPolicyResponseBodyModifyResults) GoString() string {
	return s.String()
}

func (s *ModifyResourceCenterPolicyResponseBodyModifyResults) GetCheckResult() *bool {
	return s.CheckResult
}

func (s *ModifyResourceCenterPolicyResponseBodyModifyResults) GetResourceId() *string {
	return s.ResourceId
}

func (s *ModifyResourceCenterPolicyResponseBodyModifyResults) SetCheckResult(v bool) *ModifyResourceCenterPolicyResponseBodyModifyResults {
	s.CheckResult = &v
	return s
}

func (s *ModifyResourceCenterPolicyResponseBodyModifyResults) SetResourceId(v string) *ModifyResourceCenterPolicyResponseBodyModifyResults {
	s.ResourceId = &v
	return s
}

func (s *ModifyResourceCenterPolicyResponseBodyModifyResults) Validate() error {
	return dara.Validate(s)
}
