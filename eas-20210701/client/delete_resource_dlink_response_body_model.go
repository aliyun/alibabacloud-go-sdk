// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceDLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteResourceDLinkResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteResourceDLinkResponseBody
	GetRequestId() *string
}

type DeleteResourceDLinkResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Network interfaces are deleting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceDLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceDLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceDLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteResourceDLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceDLinkResponseBody) SetMessage(v string) *DeleteResourceDLinkResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceDLinkResponseBody) SetRequestId(v string) *DeleteResourceDLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceDLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
