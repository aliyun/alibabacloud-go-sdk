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
	// F5178C10-1407-4987-9133-DE4DC9119F34
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
