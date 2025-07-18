// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProcessId(v string) *GetApprovalProcessRequest
	GetProcessId() *string
}

type GetApprovalProcessRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// approval-process-fcc351b8a95b****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s GetApprovalProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessRequest) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetApprovalProcessRequest) SetProcessId(v string) *GetApprovalProcessRequest {
	s.ProcessId = &v
	return s
}

func (s *GetApprovalProcessRequest) Validate() error {
	return dara.Validate(s)
}
