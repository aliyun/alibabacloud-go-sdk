// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlPolicyPositionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ModifyControlPolicyPositionRequest
	GetDirection() *string
	SetLang(v string) *ModifyControlPolicyPositionRequest
	GetLang() *string
	SetNewOrder(v string) *ModifyControlPolicyPositionRequest
	GetNewOrder() *string
	SetOldOrder(v string) *ModifyControlPolicyPositionRequest
	GetOldOrder() *string
	SetSourceIp(v string) *ModifyControlPolicyPositionRequest
	GetSourceIp() *string
}

type ModifyControlPolicyPositionRequest struct {
	// The direction of the traffic to which the IPv4 access control policy applies. Valid values:
	//
	// 	- in: inbound traffic
	//
	// 	- out: outbound traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- zh: Chinese (default)
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The new priority of the IPv4 access control policy. You must specify a numeric value for this parameter. The value 1 indicates the highest priority. A larger value indicates a lower priority.
	//
	// >  The new priority cannot exceed the priority range of the IPv4 access control policy. Otherwise, an error occurs when you call this operation. Before you call this operation, we recommend that you use the [DescribePolicyPriorUsed](https://help.aliyun.com/document_detail/138862.html) operation to query the priority range of the IPv4 access control policy in the specified direction.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The original priority of the IPv4 access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	OldOrder *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyControlPolicyPositionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlPolicyPositionRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPositionRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyControlPolicyPositionRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyControlPolicyPositionRequest) GetNewOrder() *string {
	return s.NewOrder
}

func (s *ModifyControlPolicyPositionRequest) GetOldOrder() *string {
	return s.OldOrder
}

func (s *ModifyControlPolicyPositionRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyControlPolicyPositionRequest) SetDirection(v string) *ModifyControlPolicyPositionRequest {
	s.Direction = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetLang(v string) *ModifyControlPolicyPositionRequest {
	s.Lang = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetNewOrder(v string) *ModifyControlPolicyPositionRequest {
	s.NewOrder = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetOldOrder(v string) *ModifyControlPolicyPositionRequest {
	s.OldOrder = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetSourceIp(v string) *ModifyControlPolicyPositionRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) Validate() error {
	return dara.Validate(s)
}
