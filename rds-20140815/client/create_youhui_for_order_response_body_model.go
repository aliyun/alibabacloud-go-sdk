// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYouhuiForOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateYouhuiForOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateYouhuiForOrderResponseBody
	GetRequestId() *string
	SetYouhuiId(v string) *CreateYouhuiForOrderResponseBody
	GetYouhuiId() *string
}

type CreateYouhuiForOrderResponseBody struct {
	// The response parameters.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0688F1D2-CDA8-5617-A43C-ADAC61D80D43
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 221201******
	YouhuiId *string `json:"YouhuiId,omitempty" xml:"YouhuiId,omitempty"`
}

func (s CreateYouhuiForOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateYouhuiForOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateYouhuiForOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateYouhuiForOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateYouhuiForOrderResponseBody) GetYouhuiId() *string {
	return s.YouhuiId
}

func (s *CreateYouhuiForOrderResponseBody) SetMessage(v string) *CreateYouhuiForOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateYouhuiForOrderResponseBody) SetRequestId(v string) *CreateYouhuiForOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateYouhuiForOrderResponseBody) SetYouhuiId(v string) *CreateYouhuiForOrderResponseBody {
	s.YouhuiId = &v
	return s
}

func (s *CreateYouhuiForOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
