// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBandwidthResourcePackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *CreateBandwidthResourcePackagesResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateBandwidthResourcePackagesResponseBody
	GetRequestId() *string
}

type CreateBandwidthResourcePackagesResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 24251717783****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of a request.
	//
	// example:
	//
	// AE7B699F-625C-587E-BC5F-1395CA969681
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBandwidthResourcePackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBandwidthResourcePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBandwidthResourcePackagesResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateBandwidthResourcePackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBandwidthResourcePackagesResponseBody) SetOrderId(v int64) *CreateBandwidthResourcePackagesResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateBandwidthResourcePackagesResponseBody) SetRequestId(v string) *CreateBandwidthResourcePackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBandwidthResourcePackagesResponseBody) Validate() error {
	return dara.Validate(s)
}
