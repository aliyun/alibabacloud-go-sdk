// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallUniBackupAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallUniBackupAgentResponseBody
	GetRequestId() *string
}

type InstallUniBackupAgentResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 53ACA55D-0325-5056-A72D-D0EC0B9C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallUniBackupAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallUniBackupAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallUniBackupAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallUniBackupAgentResponseBody) SetRequestId(v string) *InstallUniBackupAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallUniBackupAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
