// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAutoScalerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateServiceAutoScalerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceAutoScalerResponseBody
	GetRequestId() *string
}

type CreateServiceAutoScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succeed to auto scale service [foo]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceAutoScalerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceAutoScalerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceAutoScalerResponseBody) SetMessage(v string) *CreateServiceAutoScalerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceAutoScalerResponseBody) SetRequestId(v string) *CreateServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceAutoScalerResponseBody) Validate() error {
	return dara.Validate(s)
}
