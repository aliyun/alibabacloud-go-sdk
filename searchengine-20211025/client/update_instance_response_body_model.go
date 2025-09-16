// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceResponseBody
	GetRequestId() *string
	SetResult(v *UpdateInstanceResponseBodyResult) *UpdateInstanceResponseBody
	GetResult() *UpdateInstanceResponseBodyResult
}

type UpdateInstanceResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 90D6B8F5-FE97-4509-9AAB-367836C51818
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results returned.
	Result *UpdateInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceResponseBody) GetResult() *UpdateInstanceResponseBodyResult {
	return s.Result
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetResult(v *UpdateInstanceResponseBodyResult) *UpdateInstanceResponseBody {
	s.Result = v
	return s
}

func (s *UpdateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceResponseBodyResult struct {
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The commodity code of the instance.
	//
	// example:
	//
	// ha3-code
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance was created
	//
	// example:
	//
	// 2018-12-06T11:17:49.0
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the instance
	//
	// example:
	//
	// Test instance
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The time when the instance expires
	//
	// example:
	//
	// 2019-01-06T16:00:00.0
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// Indicates whether an overdue payment is involved
	//
	// example:
	//
	// false
	InDebt *bool `json:"inDebt,omitempty" xml:"inDebt,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ha-cn-0ju2s170b03
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock status
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aeky6hthboewpuy
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the instance was last updated
	//
	// example:
	//
	// 2018-12-06T11:17:49.0
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s UpdateInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *UpdateInstanceResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *UpdateInstanceResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateInstanceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *UpdateInstanceResponseBodyResult) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *UpdateInstanceResponseBodyResult) GetInDebt() *bool {
	return s.InDebt
}

func (s *UpdateInstanceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceResponseBodyResult) GetLockMode() *string {
	return s.LockMode
}

func (s *UpdateInstanceResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateInstanceResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *UpdateInstanceResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateInstanceResponseBodyResult) SetChargeType(v string) *UpdateInstanceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetCommodityCode(v string) *UpdateInstanceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetCreateTime(v string) *UpdateInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetDescription(v string) *UpdateInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetExpiredTime(v string) *UpdateInstanceResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetInDebt(v bool) *UpdateInstanceResponseBodyResult {
	s.InDebt = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetInstanceId(v string) *UpdateInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetLockMode(v string) *UpdateInstanceResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetResourceGroupId(v string) *UpdateInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetStatus(v string) *UpdateInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetUpdateTime(v string) *UpdateInstanceResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
