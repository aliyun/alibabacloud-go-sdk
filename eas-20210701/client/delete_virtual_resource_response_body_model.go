// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteVirtualResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVirtualResourceResponseBody
	GetRequestId() *string
}

type DeleteVirtualResourceResponseBody struct {
	// The information about the operation result.
	//
	// example:
	//
	// Successfully deleted virtual resource eas-vr-npovr28onap1xxxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVirtualResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVirtualResourceResponseBody) SetMessage(v string) *DeleteVirtualResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVirtualResourceResponseBody) SetRequestId(v string) *DeleteVirtualResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVirtualResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
