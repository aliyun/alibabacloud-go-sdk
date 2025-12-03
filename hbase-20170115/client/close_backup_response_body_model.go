// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseBackupResponseBody
	GetRequestId() *string
}

type CloseBackupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CloseBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseBackupResponseBody) SetRequestId(v string) *CloseBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
