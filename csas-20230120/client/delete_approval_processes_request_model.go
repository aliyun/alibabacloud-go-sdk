// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApprovalProcessesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProcessIds(v []*string) *DeleteApprovalProcessesRequest
	GetProcessIds() []*string
}

type DeleteApprovalProcessesRequest struct {
	// This parameter is required.
	ProcessIds []*string `json:"ProcessIds,omitempty" xml:"ProcessIds,omitempty" type:"Repeated"`
}

func (s DeleteApprovalProcessesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApprovalProcessesRequest) GoString() string {
	return s.String()
}

func (s *DeleteApprovalProcessesRequest) GetProcessIds() []*string {
	return s.ProcessIds
}

func (s *DeleteApprovalProcessesRequest) SetProcessIds(v []*string) *DeleteApprovalProcessesRequest {
	s.ProcessIds = v
	return s
}

func (s *DeleteApprovalProcessesRequest) Validate() error {
	return dara.Validate(s)
}
