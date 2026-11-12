// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommandAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCommandAttributeResponseBody
	GetRequestId() *string
}

type UpdateCommandAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCommandAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommandAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCommandAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCommandAttributeResponseBody) SetRequestId(v string) *UpdateCommandAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCommandAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
