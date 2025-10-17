// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDatabaseZonalResponseBody
	GetRequestId() *string
}

type CreateDatabaseZonalResponseBody struct {
	// example:
	//
	// 93E98F25-BE02-40DA-83E3-F77F8D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatabaseZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseZonalResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatabaseZonalResponseBody) SetRequestId(v string) *CreateDatabaseZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatabaseZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
