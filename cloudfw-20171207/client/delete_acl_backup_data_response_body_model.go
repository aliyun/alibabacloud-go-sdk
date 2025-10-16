// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclBackupDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAclBackupDataResponseBody
	GetRequestId() *string
}

type DeleteAclBackupDataResponseBody struct {
	// example:
	//
	// 75E60025-43C5-5635-B7B7-272C5246****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAclBackupDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclBackupDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclBackupDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAclBackupDataResponseBody) SetRequestId(v string) *DeleteAclBackupDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAclBackupDataResponseBody) Validate() error {
	return dara.Validate(s)
}
