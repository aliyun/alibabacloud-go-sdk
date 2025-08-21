// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnsServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsServiceId(v string) *CreateEnsServiceRequest
	GetEnsServiceId() *string
	SetOrderType(v string) *CreateEnsServiceRequest
	GetOrderType() *string
}

type CreateEnsServiceRequest struct {
	// The ID of the resource that you want to obtain. You can specify only one ID in a request.
	//
	// This parameter is required.
	//
	// example:
	//
	// ens-20190806****
	EnsServiceId *string `json:"EnsServiceId,omitempty" xml:"EnsServiceId,omitempty"`
	// The operation to perform after you preview the created edge service. Valid values:
	//
	// 	- **Buy**: create
	//
	// 	- **Upgrade**: change
	//
	// This parameter is required.
	//
	// example:
	//
	// Buy
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s CreateEnsServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateEnsServiceRequest) GetEnsServiceId() *string {
	return s.EnsServiceId
}

func (s *CreateEnsServiceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateEnsServiceRequest) SetEnsServiceId(v string) *CreateEnsServiceRequest {
	s.EnsServiceId = &v
	return s
}

func (s *CreateEnsServiceRequest) SetOrderType(v string) *CreateEnsServiceRequest {
	s.OrderType = &v
	return s
}

func (s *CreateEnsServiceRequest) Validate() error {
	return dara.Validate(s)
}
