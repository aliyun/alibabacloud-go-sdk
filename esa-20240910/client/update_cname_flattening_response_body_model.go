// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCnameFlatteningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCnameFlatteningResponseBody
	GetRequestId() *string
}

type UpdateCnameFlatteningResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCnameFlatteningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCnameFlatteningResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCnameFlatteningResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCnameFlatteningResponseBody) SetRequestId(v string) *UpdateCnameFlatteningResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCnameFlatteningResponseBody) Validate() error {
	return dara.Validate(s)
}
