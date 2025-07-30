// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupExpireTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBackupExpireTimeResponseBody
	GetRequestId() *string
}

type ModifyBackupExpireTimeResponseBody struct {
	// example:
	//
	// B560AAD5-5027-51AD-A0D4-FA4DB1A76F40
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupExpireTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupExpireTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupExpireTimeResponseBody) SetRequestId(v string) *ModifyBackupExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupExpireTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
