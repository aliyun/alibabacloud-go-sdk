// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceLabelResponseBody
	GetRequestId() *string
}

type UpdateServiceLabelResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succeed to update service [service_from_XXXX] labels.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceLabelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceLabelResponseBody) SetMessage(v string) *UpdateServiceLabelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceLabelResponseBody) SetRequestId(v string) *UpdateServiceLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
