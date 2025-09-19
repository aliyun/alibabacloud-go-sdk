// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPolicyMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBackupPolicyMachineResponseBody
	GetRequestId() *string
}

type DeleteBackupPolicyMachineResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D0D6E6E4-CB8C-4897-B852-46AEFDA04B21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackupPolicyMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPolicyMachineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupPolicyMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupPolicyMachineResponseBody) SetRequestId(v string) *DeleteBackupPolicyMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupPolicyMachineResponseBody) Validate() error {
	return dara.Validate(s)
}
