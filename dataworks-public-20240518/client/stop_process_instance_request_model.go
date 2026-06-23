// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopProcessInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StopProcessInstanceRequest
	GetClientToken() *string
	SetProcessInstanceId(v string) *StopProcessInstanceRequest
	GetProcessInstanceId() *string
}

type StopProcessInstanceRequest struct {
	// A token to ensure that the request is idempotent. We recommend that you use a universally unique identifier (UUID).
	//
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The process instance ID. This operation supports approval forms from both the old and new versions of Security Center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 176906667488145
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
}

func (s StopProcessInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopProcessInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopProcessInstanceRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *StopProcessInstanceRequest) SetClientToken(v string) *StopProcessInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopProcessInstanceRequest) SetProcessInstanceId(v string) *StopProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *StopProcessInstanceRequest) Validate() error {
	return dara.Validate(s)
}
