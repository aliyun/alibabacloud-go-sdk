// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySupabaseBackupPolicyResponseBody
	GetRequestId() *string
}

type ModifySupabaseBackupPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySupabaseBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySupabaseBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySupabaseBackupPolicyResponseBody) SetRequestId(v string) *ModifySupabaseBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySupabaseBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
