// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallUniBackupAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UninstallUniBackupAgentResponseBody
	GetRequestId() *string
}

type UninstallUniBackupAgentResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9D97AFC3-AA58-5B8F-BBC4-16D7D8AB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallUniBackupAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallUniBackupAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallUniBackupAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallUniBackupAgentResponseBody) SetRequestId(v string) *UninstallUniBackupAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallUniBackupAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
