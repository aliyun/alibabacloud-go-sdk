// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUniBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUniBackupPolicyResponseBody
	GetRequestId() *string
}

type ModifyUniBackupPolicyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3F6FAB39-5AF7-5B2C-A17D-16D52AE0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUniBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUniBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUniBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUniBackupPolicyResponseBody) SetRequestId(v string) *ModifyUniBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUniBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
