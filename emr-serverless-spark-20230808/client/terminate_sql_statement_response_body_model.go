// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateSqlStatementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TerminateSqlStatementResponseBody
	GetRequestId() *string
}

type TerminateSqlStatementResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s TerminateSqlStatementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateSqlStatementResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateSqlStatementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateSqlStatementResponseBody) SetRequestId(v string) *TerminateSqlStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateSqlStatementResponseBody) Validate() error {
	return dara.Validate(s)
}
