// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateManagedTransformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateManagedTransformResponseBody
	GetRequestId() *string
}

type UpdateManagedTransformResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateManagedTransformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedTransformResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateManagedTransformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateManagedTransformResponseBody) SetRequestId(v string) *UpdateManagedTransformResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateManagedTransformResponseBody) Validate() error {
	return dara.Validate(s)
}
