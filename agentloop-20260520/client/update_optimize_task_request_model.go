// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOptimizeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v map[string]interface{}) *UpdateOptimizeTaskRequest
	GetConfig() map[string]interface{}
	SetStatus(v string) *UpdateOptimizeTaskRequest
	GetStatus() *string
	SetClientToken(v string) *UpdateOptimizeTaskRequest
	GetClientToken() *string
}

type UpdateOptimizeTaskRequest struct {
	// example:
	//
	// {"ioVersion":2,"inputs":{"experienceLibrary":{"type":"dataset","dataset":"experiences"},"trajectory":{"type":"sls","project":"my-project","logstore":"agent-trajectory","format":"atif"},"prompt":{"type":"inline","format":"markdown","content":"# Guidance\\\\nExtract evidence-backed tool contracts."}},"outputs":{"changePlan":{"type":"dataset","dataset":"experience-plans"}},"operators":["tool_contract"],"application":{"mode":"manual"}}
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// Disabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateOptimizeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOptimizeTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateOptimizeTaskRequest) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *UpdateOptimizeTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateOptimizeTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateOptimizeTaskRequest) SetConfig(v map[string]interface{}) *UpdateOptimizeTaskRequest {
	s.Config = v
	return s
}

func (s *UpdateOptimizeTaskRequest) SetStatus(v string) *UpdateOptimizeTaskRequest {
	s.Status = &v
	return s
}

func (s *UpdateOptimizeTaskRequest) SetClientToken(v string) *UpdateOptimizeTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateOptimizeTaskRequest) Validate() error {
	return dara.Validate(s)
}
