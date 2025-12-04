// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSQLAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSQLAccountResponseBody
	GetRequestId() *string
}

type CreateSQLAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSQLAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSQLAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSQLAccountResponseBody) SetRequestId(v string) *CreateSQLAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSQLAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
