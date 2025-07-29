// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotasValue interface {
	dara.Model
	String() string
	GoString() string
	SetQuota(v string) *QuotasValue
	GetQuota() *string
	SetOperationCode(v string) *QuotasValue
	GetOperationCode() *string
	SetAdjustable(v bool) *QuotasValue
	GetAdjustable() *bool
	SetUnit(v string) *QuotasValue
	GetUnit() *string
}

type QuotasValue struct {
	// The value of the quota. If the quota limit is reached, submit an application in the [Quota Center console](https://quotas.console.aliyun.com/products/csk/quotas) to increase the quota.
	//
	// example:
	//
	// 1
	Quota *string `json:"quota,omitempty" xml:"quota,omitempty"`
	// The quota code.
	//
	// example:
	//
	// q_Kubernetes_Cluster
	OperationCode *string `json:"operation_code,omitempty" xml:"operation_code,omitempty"`
	// Indicates whether the quota is adjustable.
	//
	// example:
	//
	// true
	Adjustable *bool `json:"adjustable,omitempty" xml:"adjustable,omitempty"`
	// The unit.
	//
	// example:
	//
	// Cluster
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QuotasValue) String() string {
	return dara.Prettify(s)
}

func (s QuotasValue) GoString() string {
	return s.String()
}

func (s *QuotasValue) GetQuota() *string {
	return s.Quota
}

func (s *QuotasValue) GetOperationCode() *string {
	return s.OperationCode
}

func (s *QuotasValue) GetAdjustable() *bool {
	return s.Adjustable
}

func (s *QuotasValue) GetUnit() *string {
	return s.Unit
}

func (s *QuotasValue) SetQuota(v string) *QuotasValue {
	s.Quota = &v
	return s
}

func (s *QuotasValue) SetOperationCode(v string) *QuotasValue {
	s.OperationCode = &v
	return s
}

func (s *QuotasValue) SetAdjustable(v bool) *QuotasValue {
	s.Adjustable = &v
	return s
}

func (s *QuotasValue) SetUnit(v string) *QuotasValue {
	s.Unit = &v
	return s
}

func (s *QuotasValue) Validate() error {
	return dara.Validate(s)
}
