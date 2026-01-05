// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProvisionedProductId(v string) *LaunchProductResponseBody
	GetProvisionedProductId() *string
	SetRequestId(v string) *LaunchProductResponseBody
	GetRequestId() *string
}

type LaunchProductResponseBody struct {
	// The ID of the instance
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LaunchProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LaunchProductResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchProductResponseBody) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *LaunchProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LaunchProductResponseBody) SetProvisionedProductId(v string) *LaunchProductResponseBody {
	s.ProvisionedProductId = &v
	return s
}

func (s *LaunchProductResponseBody) SetRequestId(v string) *LaunchProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *LaunchProductResponseBody) Validate() error {
	return dara.Validate(s)
}
