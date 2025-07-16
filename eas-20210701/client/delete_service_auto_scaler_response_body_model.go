// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceAutoScalerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteServiceAutoScalerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceAutoScalerResponseBody
	GetRequestId() *string
}

type DeleteServiceAutoScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Autoscaler for service [foo] deleted successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceAutoScalerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceAutoScalerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceAutoScalerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceAutoScalerResponseBody) SetMessage(v string) *DeleteServiceAutoScalerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceAutoScalerResponseBody) SetRequestId(v string) *DeleteServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceAutoScalerResponseBody) Validate() error {
	return dara.Validate(s)
}
