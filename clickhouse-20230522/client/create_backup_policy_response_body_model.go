// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateBackupPolicyResponseBody
	GetRequestId() *string
}

type CreateBackupPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupPolicyResponseBody) SetRequestId(v string) *CreateBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
