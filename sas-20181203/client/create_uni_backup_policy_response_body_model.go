// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUniBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUniBackupPolicyResponseBody
	GetRequestId() *string
}

type CreateUniBackupPolicyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2FAEB7D0-C0B9-581C-877A-F80F50AA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUniBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUniBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUniBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUniBackupPolicyResponseBody) SetRequestId(v string) *CreateUniBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUniBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
