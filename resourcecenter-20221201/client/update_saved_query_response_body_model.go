// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSavedQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSavedQueryResponseBody
	GetRequestId() *string
}

type UpdateSavedQueryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSavedQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSavedQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSavedQueryResponseBody) SetRequestId(v string) *UpdateSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSavedQueryResponseBody) Validate() error {
	return dara.Validate(s)
}
