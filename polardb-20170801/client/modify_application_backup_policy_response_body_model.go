// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApplicationBackupPolicyResponseBody
	GetRequestId() *string
}

type ModifyApplicationBackupPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApplicationBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApplicationBackupPolicyResponseBody) SetRequestId(v string) *ModifyApplicationBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
