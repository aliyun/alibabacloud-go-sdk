// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFilterResponseBody
	GetRequestId() *string
}

type UpdateFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3C5CDBF6-4DB7-53E9-ADDC-5919E3FACF6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFilterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFilterResponseBody) SetRequestId(v string) *UpdateFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFilterResponseBody) Validate() error {
	return dara.Validate(s)
}
