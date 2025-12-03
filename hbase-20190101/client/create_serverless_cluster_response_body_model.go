// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerlessClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateServerlessClusterResponseBody
	GetClusterId() *string
	SetOrderId(v string) *CreateServerlessClusterResponseBody
	GetOrderId() *string
	SetPassWord(v string) *CreateServerlessClusterResponseBody
	GetPassWord() *string
	SetRequestId(v string) *CreateServerlessClusterResponseBody
	GetRequestId() *string
}

type CreateServerlessClusterResponseBody struct {
	// example:
	//
	// sh-bp1a969y7681****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 23232453233*****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// *********
	PassWord *string `json:"PassWord,omitempty" xml:"PassWord,omitempty"`
	// example:
	//
	// 3E19E345-101D-4014-946C-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServerlessClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServerlessClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerlessClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateServerlessClusterResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateServerlessClusterResponseBody) GetPassWord() *string {
	return s.PassWord
}

func (s *CreateServerlessClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServerlessClusterResponseBody) SetClusterId(v string) *CreateServerlessClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateServerlessClusterResponseBody) SetOrderId(v string) *CreateServerlessClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateServerlessClusterResponseBody) SetPassWord(v string) *CreateServerlessClusterResponseBody {
	s.PassWord = &v
	return s
}

func (s *CreateServerlessClusterResponseBody) SetRequestId(v string) *CreateServerlessClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerlessClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
