// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePolarFsResponseBody
	GetRequestId() *string
}

type DeletePolarFsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C890995A-CF06-4F4D-8DB8-DD26C2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolarFsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolarFsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolarFsResponseBody) SetRequestId(v string) *DeletePolarFsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolarFsResponseBody) Validate() error {
	return dara.Validate(s)
}
