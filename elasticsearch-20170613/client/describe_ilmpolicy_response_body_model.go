// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeILMPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeILMPolicyResponseBody
	GetRequestId() *string
	SetResult(v *DescribeILMPolicyResponseBodyResult) *DescribeILMPolicyResponseBody
	GetResult() *DescribeILMPolicyResponseBodyResult
}

type DescribeILMPolicyResponseBody struct {
	// example:
	//
	// FF44681E-FD41-4FDE-B8DF-295DCDD6****
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeILMPolicyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeILMPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeILMPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeILMPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeILMPolicyResponseBody) GetResult() *DescribeILMPolicyResponseBodyResult {
	return s.Result
}

func (s *DescribeILMPolicyResponseBody) SetRequestId(v string) *DescribeILMPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeILMPolicyResponseBody) SetResult(v *DescribeILMPolicyResponseBodyResult) *DescribeILMPolicyResponseBody {
	s.Result = v
	return s
}

func (s *DescribeILMPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeILMPolicyResponseBodyResult struct {
	// example:
	//
	// ilm-history-ilm-policy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// {"cold":{"minAge":"30d","actions":{"allocate":{"numberOfReplicas":1,"require":{"boxType":"warm"}},"setPriority":{"priority":100}}},"hot":{"minAge":"0s","actions":{"rollover":{"maxAge":"30d","maxDocs":10000,"maxSize":"50gb"},"setPriority":{"priority":1000}}},"delete":{"minAge":"30d","actions":{"delete":{}}}}
	Phases map[string]interface{} `json:"phases,omitempty" xml:"phases,omitempty"`
}

func (s DescribeILMPolicyResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeILMPolicyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeILMPolicyResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeILMPolicyResponseBodyResult) GetPhases() map[string]interface{} {
	return s.Phases
}

func (s *DescribeILMPolicyResponseBodyResult) SetName(v string) *DescribeILMPolicyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeILMPolicyResponseBodyResult) SetPhases(v map[string]interface{}) *DescribeILMPolicyResponseBodyResult {
	s.Phases = v
	return s
}

func (s *DescribeILMPolicyResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
