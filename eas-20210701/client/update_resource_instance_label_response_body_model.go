// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceInstanceLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateResourceInstanceLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateResourceInstanceLabelResponseBody
	GetRequestId() *string
}

type UpdateResourceInstanceLabelResponseBody struct {
	// The message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceInstanceLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceInstanceLabelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateResourceInstanceLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceInstanceLabelResponseBody) SetMessage(v string) *UpdateResourceInstanceLabelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateResourceInstanceLabelResponseBody) SetRequestId(v string) *UpdateResourceInstanceLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceInstanceLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
