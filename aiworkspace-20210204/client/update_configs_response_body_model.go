// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateConfigsResponseBody
	GetRequestId() *string
}

type UpdateConfigsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A******C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigsResponseBody) SetRequestId(v string) *UpdateConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
