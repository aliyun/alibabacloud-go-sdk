// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDatabaseAccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetDatabaseAccountPasswordResponseBody
	GetRequestId() *string
}

type ResetDatabaseAccountPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDatabaseAccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetDatabaseAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDatabaseAccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetDatabaseAccountPasswordResponseBody) SetRequestId(v string) *ResetDatabaseAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetDatabaseAccountPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
