// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProcessInstanceId(v string) *GetProcessInstanceRequest
	GetProcessInstanceId() *string
}

type GetProcessInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 332066440109224007
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
}

func (s GetProcessInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetProcessInstanceRequest) SetProcessInstanceId(v string) *GetProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetProcessInstanceRequest) Validate() error {
	return dara.Validate(s)
}
