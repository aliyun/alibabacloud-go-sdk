// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAclBackupDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddAclBackupDataResponseBody
	GetRequestId() *string
}

type AddAclBackupDataResponseBody struct {
	// example:
	//
	// 244EA37C-D2AB-54A7-B6E3-7ED0E9A1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAclBackupDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAclBackupDataResponseBody) GoString() string {
	return s.String()
}

func (s *AddAclBackupDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAclBackupDataResponseBody) SetRequestId(v string) *AddAclBackupDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAclBackupDataResponseBody) Validate() error {
	return dara.Validate(s)
}
