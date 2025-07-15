// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskEncryptionServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CreateDiskEncryptionServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateDiskEncryptionServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDiskEncryptionServiceResponseBody
	GetSuccess() *bool
}

type CreateDiskEncryptionServiceResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 214552063030752
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDiskEncryptionServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskEncryptionServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiskEncryptionServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDiskEncryptionServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiskEncryptionServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDiskEncryptionServiceResponseBody) SetOrderId(v string) *CreateDiskEncryptionServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDiskEncryptionServiceResponseBody) SetRequestId(v string) *CreateDiskEncryptionServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiskEncryptionServiceResponseBody) SetSuccess(v bool) *CreateDiskEncryptionServiceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDiskEncryptionServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
