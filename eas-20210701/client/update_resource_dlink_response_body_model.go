// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceDLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateResourceDLinkResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateResourceDLinkResponseBody
	GetRequestId() *string
}

type UpdateResourceDLinkResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Network interfaces are updating
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceDLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceDLinkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceDLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateResourceDLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceDLinkResponseBody) SetMessage(v string) *UpdateResourceDLinkResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateResourceDLinkResponseBody) SetRequestId(v string) *UpdateResourceDLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceDLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
