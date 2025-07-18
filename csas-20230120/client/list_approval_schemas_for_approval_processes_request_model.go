// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalSchemasForApprovalProcessesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProcessIds(v []*string) *ListApprovalSchemasForApprovalProcessesRequest
	GetProcessIds() []*string
}

type ListApprovalSchemasForApprovalProcessesRequest struct {
	// This parameter is required.
	ProcessIds []*string `json:"ProcessIds,omitempty" xml:"ProcessIds,omitempty" type:"Repeated"`
}

func (s ListApprovalSchemasForApprovalProcessesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalSchemasForApprovalProcessesRequest) GoString() string {
	return s.String()
}

func (s *ListApprovalSchemasForApprovalProcessesRequest) GetProcessIds() []*string {
	return s.ProcessIds
}

func (s *ListApprovalSchemasForApprovalProcessesRequest) SetProcessIds(v []*string) *ListApprovalSchemasForApprovalProcessesRequest {
	s.ProcessIds = v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesRequest) Validate() error {
	return dara.Validate(s)
}
