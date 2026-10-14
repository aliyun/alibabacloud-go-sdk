// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOptimizeRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtra(v map[string]interface{}) *UpdateOptimizeRunRequest
	GetExtra() map[string]interface{}
	SetStatus(v string) *UpdateOptimizeRunRequest
	GetStatus() *string
	SetClientToken(v string) *UpdateOptimizeRunRequest
	GetClientToken() *string
}

type UpdateOptimizeRunRequest struct {
	// The extra parameters. To retry a failed artifact application, set application.action to retry and do not specify status. Do not specify this field when canceling a run.
	//
	// example:
	//
	// {"application":{"action":"retry"}}
	Extra map[string]interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
	// The status. Only Cancelled is supported when canceling a run. The execution status is maintained by the server. Do not specify this field when retrying an artifact application.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateOptimizeRunRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOptimizeRunRequest) GoString() string {
	return s.String()
}

func (s *UpdateOptimizeRunRequest) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *UpdateOptimizeRunRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateOptimizeRunRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateOptimizeRunRequest) SetExtra(v map[string]interface{}) *UpdateOptimizeRunRequest {
	s.Extra = v
	return s
}

func (s *UpdateOptimizeRunRequest) SetStatus(v string) *UpdateOptimizeRunRequest {
	s.Status = &v
	return s
}

func (s *UpdateOptimizeRunRequest) SetClientToken(v string) *UpdateOptimizeRunRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateOptimizeRunRequest) Validate() error {
	return dara.Validate(s)
}
