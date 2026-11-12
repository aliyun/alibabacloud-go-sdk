// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDatabaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartDatabaseInstanceResponseBody
	GetRequestId() *string
}

type RestartDatabaseInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDatabaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDatabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDatabaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDatabaseInstanceResponseBody) SetRequestId(v string) *RestartDatabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDatabaseInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
