// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDatabaseResponseBody
	GetRequestId() *string
}

type CreateDatabaseResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 93E98F25-BE02-40DA-83E3-F77F8D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatabaseResponseBody) SetRequestId(v string) *CreateDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
