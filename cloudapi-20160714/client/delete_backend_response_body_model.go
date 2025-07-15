// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBackendResponseBody
	GetRequestId() *string
}

type DeleteBackendResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AD00F8C0-311B-54A9-ADE2-2436771012DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackendResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackendResponseBody) SetRequestId(v string) *DeleteBackendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackendResponseBody) Validate() error {
	return dara.Validate(s)
}
