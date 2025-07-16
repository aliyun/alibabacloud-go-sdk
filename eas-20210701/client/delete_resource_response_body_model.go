// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteResourceResponseBody
	GetRequestId() *string
}

type DeleteResourceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Resource [eas-r-asdasdasd] is deleted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82-9624-EC2B1779848E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceResponseBody) SetMessage(v string) *DeleteResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
