// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallBackupClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UninstallBackupClientResponseBody
	GetRequestId() *string
}

type UninstallBackupClientResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8eec3b63-18af-454b-8c17-aabcf7190b70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallBackupClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallBackupClientResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallBackupClientResponseBody) SetRequestId(v string) *UninstallBackupClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallBackupClientResponseBody) Validate() error {
	return dara.Validate(s)
}
