// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBackupResponseBody
	GetRequestId() *string
}

type DeleteBackupResponseBody struct {
	// example:
	//
	// 7A9807F0-1301-4154-9849-6497E94A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupResponseBody) SetRequestId(v string) *DeleteBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
