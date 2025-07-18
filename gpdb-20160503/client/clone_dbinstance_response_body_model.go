// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloneDBInstanceResponseBody
	GetRequestId() *string
}

type CloneDBInstanceResponseBody struct {
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CloneDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneDBInstanceResponseBody) SetRequestId(v string) *CloneDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
