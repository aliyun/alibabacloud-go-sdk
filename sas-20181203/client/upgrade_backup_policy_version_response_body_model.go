// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeBackupPolicyVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeBackupPolicyVersionResponseBody
	GetRequestId() *string
}

type UpgradeBackupPolicyVersionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9F5B8057-3E44-54DC-AC89-6814A4CD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeBackupPolicyVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupPolicyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeBackupPolicyVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeBackupPolicyVersionResponseBody) SetRequestId(v string) *UpgradeBackupPolicyVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeBackupPolicyVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
