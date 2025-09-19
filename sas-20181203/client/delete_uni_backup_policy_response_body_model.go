// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUniBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUniBackupPolicyResponseBody
	GetRequestId() *string
}

type DeleteUniBackupPolicyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4F07F4F1-88ED-5569-B519-FFCC9B7E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUniBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUniBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUniBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUniBackupPolicyResponseBody) SetRequestId(v string) *DeleteUniBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUniBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
