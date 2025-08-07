// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBackupPolicyResponseBody
	GetRequestId() *string
}

type ModifyBackupPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C5A5DF0E-5968-4DC1-882E-AC2FE7******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupPolicyResponseBody) SetRequestId(v string) *ModifyBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
