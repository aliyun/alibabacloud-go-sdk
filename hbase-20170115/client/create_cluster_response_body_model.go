// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateClusterResponseBody
	GetClusterId() *string
	SetOrderId(v string) *CreateClusterResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateClusterResponseBody
	GetRequestId() *string
}

type CreateClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetOrderId(v string) *CreateClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
