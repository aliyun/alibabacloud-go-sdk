// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvancedBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAdvancedBackupPolicyResponseBody
	GetRequestId() *string
}

type ModifyAdvancedBackupPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2F42BB4E-461F-5B55-A37C-53B1141C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAdvancedBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvancedBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAdvancedBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAdvancedBackupPolicyResponseBody) SetRequestId(v string) *ModifyAdvancedBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAdvancedBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
