// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBInstanceResponseBody
	GetRequestId() *string
}

type DeleteDBInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 72651AF9-7897-41A7-8B67-6C15C7F26A0A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBInstanceResponseBody) SetRequestId(v string) *DeleteDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
