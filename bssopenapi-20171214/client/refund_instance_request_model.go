// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RefundInstanceRequest
	GetClientToken() *string
	SetImmediatelyRelease(v string) *RefundInstanceRequest
	GetImmediatelyRelease() *string
	SetInstanceId(v string) *RefundInstanceRequest
	GetInstanceId() *string
	SetProductCode(v string) *RefundInstanceRequest
	GetProductCode() *string
	SetProductType(v string) *RefundInstanceRequest
	GetProductType() *string
}

type RefundInstanceRequest struct {
	// This parameter is required for scenarios that need idempotence. The UUID that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 793F021C-B589-1225-82A9-99232AEBE494
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required for unsubscription scenarios. Valid values: 1 and 0. A value of 1 specifies that the instance is immediately released. A value of 0 specifies that the instance is shut down based on the shutdown policy. This parameter is supported only for specified services. Default value: 1.
	//
	// example:
	//
	// 1
	ImmediatelyRelease *string `json:"ImmediatelyRelease,omitempty" xml:"ImmediatelyRelease,omitempty"`
	// The ID of the instance. This parameter is required for unsubscription scenarios. Do not specify a custom name for this parameter.
	//
	// example:
	//
	// i-bp1etb69sqxgl4*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The code of the service. This parameter is required for unsubscription scenarios.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service. This parameter is required for unsubscription scenarios. Unless otherwise specified, set this parameter to an empty string.
	//
	// example:
	//
	// ”“
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s RefundInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RefundInstanceRequest) GoString() string {
	return s.String()
}

func (s *RefundInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RefundInstanceRequest) GetImmediatelyRelease() *string {
	return s.ImmediatelyRelease
}

func (s *RefundInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RefundInstanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *RefundInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *RefundInstanceRequest) SetClientToken(v string) *RefundInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RefundInstanceRequest) SetImmediatelyRelease(v string) *RefundInstanceRequest {
	s.ImmediatelyRelease = &v
	return s
}

func (s *RefundInstanceRequest) SetInstanceId(v string) *RefundInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundInstanceRequest) SetProductCode(v string) *RefundInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *RefundInstanceRequest) SetProductType(v string) *RefundInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *RefundInstanceRequest) Validate() error {
	return dara.Validate(s)
}
