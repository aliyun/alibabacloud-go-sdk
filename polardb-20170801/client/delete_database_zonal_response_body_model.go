// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatabaseZonalResponseBody
	GetRequestId() *string
}

type DeleteDatabaseZonalResponseBody struct {
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatabaseZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseZonalResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatabaseZonalResponseBody) SetRequestId(v string) *DeleteDatabaseZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatabaseZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
