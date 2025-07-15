// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceGroupImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceGroupImageResponseBody
	GetRequestId() *string
}

type UpdateInstanceGroupImageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 55726272-E40B-530D-914F-5126B19C79B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceGroupImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceGroupImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceGroupImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceGroupImageResponseBody) SetRequestId(v string) *UpdateInstanceGroupImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceGroupImageResponseBody) Validate() error {
	return dara.Validate(s)
}
