// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceInstanceLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteResourceInstanceLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteResourceInstanceLabelResponseBody
	GetRequestId() *string
}

type DeleteResourceInstanceLabelResponseBody struct {
	// The message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceInstanceLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceInstanceLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstanceLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteResourceInstanceLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceInstanceLabelResponseBody) SetMessage(v string) *DeleteResourceInstanceLabelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceInstanceLabelResponseBody) SetRequestId(v string) *DeleteResourceInstanceLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceInstanceLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
