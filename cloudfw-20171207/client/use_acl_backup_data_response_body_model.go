// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUseAclBackupDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UseAclBackupDataResponseBody
	GetRequestId() *string
}

type UseAclBackupDataResponseBody struct {
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UseAclBackupDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UseAclBackupDataResponseBody) GoString() string {
	return s.String()
}

func (s *UseAclBackupDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UseAclBackupDataResponseBody) SetRequestId(v string) *UseAclBackupDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UseAclBackupDataResponseBody) Validate() error {
	return dara.Validate(s)
}
