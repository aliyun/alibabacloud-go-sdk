// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnSubTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDcdnSubTaskResponseBody
	GetRequestId() *string
}

type CreateDcdnSubTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDcdnSubTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDcdnSubTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDcdnSubTaskResponseBody) SetRequestId(v string) *CreateDcdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDcdnSubTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
