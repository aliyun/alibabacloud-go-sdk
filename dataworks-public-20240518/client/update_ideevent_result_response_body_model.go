// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIDEEventResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIDEEventResultResponseBody
	GetRequestId() *string
}

type UpdateIDEEventResultResponseBody struct {
	// The request ID. Used for troubleshooting errors.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIDEEventResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIDEEventResultResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIDEEventResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIDEEventResultResponseBody) SetRequestId(v string) *UpdateIDEEventResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIDEEventResultResponseBody) Validate() error {
	return dara.Validate(s)
}
