// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateModelVersionResponseBody
	GetRequestId() *string
}

type UpdateModelVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D5BFFEE3-6025-443F-8A03-02D61***C4B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateModelVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelVersionResponseBody) SetRequestId(v string) *UpdateModelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
