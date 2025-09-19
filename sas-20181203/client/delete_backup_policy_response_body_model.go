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
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
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
