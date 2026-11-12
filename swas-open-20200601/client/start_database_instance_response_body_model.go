// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDatabaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartDatabaseInstanceResponseBody
	GetRequestId() *string
}

type StartDatabaseInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDatabaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDatabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartDatabaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDatabaseInstanceResponseBody) SetRequestId(v string) *StartDatabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDatabaseInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
