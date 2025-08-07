// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLogBackupPolicyResponseBody
	GetRequestId() *string
}

type ModifyLogBackupPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 345174B4-FAB3-412E-A326-BEDDA9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLogBackupPolicyResponseBody) SetRequestId(v string) *ModifyLogBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLogBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
