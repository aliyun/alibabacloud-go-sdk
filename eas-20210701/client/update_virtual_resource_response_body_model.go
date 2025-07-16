// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirtualResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateVirtualResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVirtualResourceResponseBody
	GetRequestId() *string
}

type UpdateVirtualResourceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Successfully updated virtual resource eas-vr-npovr28onap1xxxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVirtualResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirtualResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVirtualResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVirtualResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVirtualResourceResponseBody) SetMessage(v string) *UpdateVirtualResourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVirtualResourceResponseBody) SetRequestId(v string) *UpdateVirtualResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVirtualResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
