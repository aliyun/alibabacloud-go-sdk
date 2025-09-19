// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallBackupClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallBackupClientResponseBody
	GetRequestId() *string
}

type InstallBackupClientResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D0D6E6E4-CB8C-4897-B852-46AEFDA04B21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallBackupClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallBackupClientResponseBody) GoString() string {
	return s.String()
}

func (s *InstallBackupClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallBackupClientResponseBody) SetRequestId(v string) *InstallBackupClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallBackupClientResponseBody) Validate() error {
	return dara.Validate(s)
}
