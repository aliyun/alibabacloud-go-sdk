// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetBackupResponseBody
	GetRequestId() *string
}

type ResetBackupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetBackupResponseBody) GoString() string {
	return s.String()
}

func (s *ResetBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetBackupResponseBody) SetRequestId(v string) *ResetBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
