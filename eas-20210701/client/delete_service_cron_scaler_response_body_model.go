// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceCronScalerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteServiceCronScalerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceCronScalerResponseBody
	GetRequestId() *string
}

type DeleteServiceCronScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Cronscaler for service [foo] deleted successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceCronScalerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceCronScalerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceCronScalerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceCronScalerResponseBody) SetMessage(v string) *DeleteServiceCronScalerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceCronScalerResponseBody) SetRequestId(v string) *DeleteServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceCronScalerResponseBody) Validate() error {
	return dara.Validate(s)
}
