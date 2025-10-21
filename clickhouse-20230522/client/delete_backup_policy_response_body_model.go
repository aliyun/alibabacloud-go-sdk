// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBackupPolicyResponseBody
	GetRequestId() *string
}

type DeleteBackupPolicyResponseBody struct {
	// example:
	//
	// 60DDD29D-E5A8-563C-88FB-06D3A1F1C609
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupPolicyResponseBody) SetRequestId(v string) *DeleteBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
