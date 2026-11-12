// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDatabaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopDatabaseInstanceResponseBody
	GetRequestId() *string
}

type StopDatabaseInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDatabaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDatabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopDatabaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDatabaseInstanceResponseBody) SetRequestId(v string) *StopDatabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDatabaseInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
