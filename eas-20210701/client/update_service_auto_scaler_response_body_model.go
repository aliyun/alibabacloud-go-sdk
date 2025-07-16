// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceAutoScalerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceAutoScalerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceAutoScalerResponseBody
	GetRequestId() *string
}

type UpdateServiceAutoScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Update auto scale for service [foo] successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceAutoScalerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceAutoScalerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceAutoScalerResponseBody) SetMessage(v string) *UpdateServiceAutoScalerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceAutoScalerResponseBody) SetRequestId(v string) *UpdateServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceAutoScalerResponseBody) Validate() error {
	return dara.Validate(s)
}
