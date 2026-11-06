// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreApplicationByBackupSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestoreApplicationByBackupSetResponseBody
	GetRequestId() *string
}

type RestoreApplicationByBackupSetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreApplicationByBackupSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreApplicationByBackupSetResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreApplicationByBackupSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreApplicationByBackupSetResponseBody) SetRequestId(v string) *RestoreApplicationByBackupSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreApplicationByBackupSetResponseBody) Validate() error {
	return dara.Validate(s)
}
