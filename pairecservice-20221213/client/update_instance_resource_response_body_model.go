// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceResourceResponseBody
	GetRequestId() *string
}

type UpdateInstanceResourceResponseBody struct {
	// example:
	//
	// 3AAA45F6-0798-5461-9360-81D133823CE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceResourceResponseBody) SetRequestId(v string) *UpdateInstanceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
