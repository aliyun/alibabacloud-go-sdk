// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBackupPolicyStatusResponseBody
	GetRequestId() *string
}

type ModifyBackupPolicyStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// E342452B-4401-5F74-9A1B-D24479851173
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPolicyStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupPolicyStatusResponseBody) SetRequestId(v string) *ModifyBackupPolicyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupPolicyStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
