// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOptimizeRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v interface{}) *CreateOptimizeRunRequest
	GetInput() interface{}
	SetPreview(v bool) *CreateOptimizeRunRequest
	GetPreview() *bool
	SetClientToken(v string) *CreateOptimizeRunRequest
	GetClientToken() *string
}

type CreateOptimizeRunRequest struct {
	// The JSON input defined by the task type, preserving the original object or array structure. ExperienceMining accepts a start/end/processTimeMaxNS object. SkillOptimize accepts an inline input array, a skills/window object for SLS, or a skills/records object for ATIF or xTrace. The server validates the input based on the task type and configuration. Memory extraction continues to use the original ContextStore API.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"start":1727136000,"end":1727139600,"processTimeMaxNS":1727139600000000000}
	Input interface{} `json:"input,omitempty" xml:"input,omitempty"`
	// Specifies whether to execute an independent pre-run of the complete algorithm. The pre-run generates candidate results and evidence without publishing them to the target dataset or advancing the formal run history. A formal run executes the algorithm independently.
	//
	// example:
	//
	// false
	Preview *bool `json:"preview,omitempty" xml:"preview,omitempty"`
	// The idempotency token. Use the same value when you retry the same creation request.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateOptimizeRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOptimizeRunRequest) GoString() string {
	return s.String()
}

func (s *CreateOptimizeRunRequest) GetInput() interface{} {
	return s.Input
}

func (s *CreateOptimizeRunRequest) GetPreview() *bool {
	return s.Preview
}

func (s *CreateOptimizeRunRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOptimizeRunRequest) SetInput(v interface{}) *CreateOptimizeRunRequest {
	s.Input = v
	return s
}

func (s *CreateOptimizeRunRequest) SetPreview(v bool) *CreateOptimizeRunRequest {
	s.Preview = &v
	return s
}

func (s *CreateOptimizeRunRequest) SetClientToken(v string) *CreateOptimizeRunRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOptimizeRunRequest) Validate() error {
	return dara.Validate(s)
}
