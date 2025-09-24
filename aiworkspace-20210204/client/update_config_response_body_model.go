// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateConfigResponseBody
	GetRequestId() *string
}

type UpdateConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A******C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigResponseBody) SetRequestId(v string) *UpdateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
