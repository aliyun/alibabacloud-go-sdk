// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenBackupResponseBody
	GetRequestId() *string
}

type OpenBackupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenBackupResponseBody) GoString() string {
	return s.String()
}

func (s *OpenBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenBackupResponseBody) SetRequestId(v string) *OpenBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
