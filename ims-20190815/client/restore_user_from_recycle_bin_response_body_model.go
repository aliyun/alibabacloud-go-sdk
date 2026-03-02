// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreUserFromRecycleBinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestoreUserFromRecycleBinResponseBody
	GetRequestId() *string
}

type RestoreUserFromRecycleBinResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreUserFromRecycleBinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreUserFromRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreUserFromRecycleBinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreUserFromRecycleBinResponseBody) SetRequestId(v string) *RestoreUserFromRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreUserFromRecycleBinResponseBody) Validate() error {
	return dara.Validate(s)
}
