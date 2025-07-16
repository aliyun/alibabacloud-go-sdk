// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateVirtualResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVirtualResourceResponseBody
	GetRequestId() *string
	SetVirtualResourceId(v string) *CreateVirtualResourceResponseBody
	GetVirtualResourceId() *string
}

type CreateVirtualResourceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Successfully created virtual resource eas-vr-npovr28onap1xxxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the virtual resource group.
	//
	// example:
	//
	// eas-vr-npovr28onap1xxxxxx
	VirtualResourceId *string `json:"VirtualResourceId,omitempty" xml:"VirtualResourceId,omitempty"`
}

func (s CreateVirtualResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVirtualResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirtualResourceResponseBody) GetVirtualResourceId() *string {
	return s.VirtualResourceId
}

func (s *CreateVirtualResourceResponseBody) SetMessage(v string) *CreateVirtualResourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVirtualResourceResponseBody) SetRequestId(v string) *CreateVirtualResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualResourceResponseBody) SetVirtualResourceId(v string) *CreateVirtualResourceResponseBody {
	s.VirtualResourceId = &v
	return s
}

func (s *CreateVirtualResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
