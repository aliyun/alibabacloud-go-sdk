// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEditingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEditingProjectResponseBody
	GetRequestId() *string
}

type UpdateEditingProjectResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEditingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEditingProjectResponseBody) SetRequestId(v string) *UpdateEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEditingProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
