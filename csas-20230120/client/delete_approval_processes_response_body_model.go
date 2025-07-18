// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApprovalProcessesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApprovalProcessesResponseBody
	GetRequestId() *string
}

type DeleteApprovalProcessesResponseBody struct {
	// example:
	//
	// B608C6AE-623D-55C4-9454-601B88AE937E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApprovalProcessesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApprovalProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApprovalProcessesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApprovalProcessesResponseBody) SetRequestId(v string) *DeleteApprovalProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApprovalProcessesResponseBody) Validate() error {
	return dara.Validate(s)
}
