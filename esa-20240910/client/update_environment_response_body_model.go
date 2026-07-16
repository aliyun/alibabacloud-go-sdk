// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEnvironmentResponseBody
	GetRequestId() *string
}

type UpdateEnvironmentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 52DD50C2-C381-13BD-A269-5FAEEB848ACD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnvironmentResponseBody) SetRequestId(v string) *UpdateEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) Validate() error {
	return dara.Validate(s)
}
