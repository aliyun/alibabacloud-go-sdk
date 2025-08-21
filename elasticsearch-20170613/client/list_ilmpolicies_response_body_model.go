// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListILMPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListILMPoliciesResponseBody
	GetRequestId() *string
	SetResult(v []*ListILMPoliciesResponseBodyResult) *ListILMPoliciesResponseBody
	GetResult() []*ListILMPoliciesResponseBodyResult
}

type ListILMPoliciesResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListILMPoliciesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListILMPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListILMPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListILMPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListILMPoliciesResponseBody) GetResult() []*ListILMPoliciesResponseBodyResult {
	return s.Result
}

func (s *ListILMPoliciesResponseBody) SetRequestId(v string) *ListILMPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListILMPoliciesResponseBody) SetResult(v []*ListILMPoliciesResponseBodyResult) *ListILMPoliciesResponseBody {
	s.Result = v
	return s
}

func (s *ListILMPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListILMPoliciesResponseBodyResult struct {
	// example:
	//
	// policy-1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// {"hot":{"minAge":"0ms","actions":{"rollover":{"maxSize":"50gb","maxAge":"30d"},"setPriority":{"priority":100}}},"delete":{"minAge":"3d","actions":{"delete":{}}}}
	Phases map[string]interface{} `json:"phases,omitempty" xml:"phases,omitempty"`
}

func (s ListILMPoliciesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListILMPoliciesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListILMPoliciesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListILMPoliciesResponseBodyResult) GetPhases() map[string]interface{} {
	return s.Phases
}

func (s *ListILMPoliciesResponseBodyResult) SetName(v string) *ListILMPoliciesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListILMPoliciesResponseBodyResult) SetPhases(v map[string]interface{}) *ListILMPoliciesResponseBodyResult {
	s.Phases = v
	return s
}

func (s *ListILMPoliciesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
