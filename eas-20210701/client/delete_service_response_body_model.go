// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceResponseBody
	GetRequestId() *string
}

type DeleteServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Service [foo] in region [cn-shanghai] is terminating
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceResponseBody) SetMessage(v string) *DeleteServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
