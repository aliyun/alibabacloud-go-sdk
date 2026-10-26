// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeKernelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpgradeKernelVersionRequest
	GetInstanceId() *string
	SetKernelVersion(v string) *UpgradeKernelVersionRequest
	GetKernelVersion() *string
	SetLang(v string) *UpgradeKernelVersionRequest
	GetLang() *string
	SetProductCode(v string) *UpgradeKernelVersionRequest
	GetProductCode() *string
	SetProductId(v int64) *UpgradeKernelVersionRequest
	GetProductId() *int64
	SetSwitchTime(v int64) *UpgradeKernelVersionRequest
	GetSwitchTime() *int64
	SetUpgradeTime(v string) *UpgradeKernelVersionRequest
	GetUpgradeTime() *string
}

type UpgradeKernelVersionRequest struct {
	// The database instance identity. You can obtain the value from Items[].InstanceId in the response of the ListDataMaskingInstances operation.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The target Milvus version. You can obtain the value from KernelVersions[].KernelVersion in the response of the ListMiniEngineVersions operation.
	//
	// example:
	//
	// rds_20220731
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// The language of the response description. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The cloud service code. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also specified, the two values must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The identifier of the cloud service type. Valid values:
	//
	// - 5: RDS.
	//
	// - 8: POLARDB.
	//
	// - 27: PolarDBX2.
	//
	// If this parameter is specified together with ProductCode, the values must match.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The specified upgrade time. The value is a UNIX timestamp in milliseconds that have elapsed since 1970-01-01 00:00:00 UTC. Specify this parameter only when UpgradeTime is set to SpecifyTime, and the time must be later than the current time.
	//
	// example:
	//
	// 1893456000000
	SwitchTime *int64 `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The execution time of the kernel upgrade. Valid values:
	//
	// - Immediate: Executes the upgrade immediately.
	//
	// - MaintainTime: Executes the upgrade during the maintenance window.
	//
	// - SpecifyTime: Executes the upgrade at a specified time.
	//
	// PolarDB-X 2.0 does not support SpecifyTime.
	//
	// example:
	//
	// MaintainTime
	UpgradeTime *string `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty"`
}

func (s UpgradeKernelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeKernelVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeKernelVersionRequest) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *UpgradeKernelVersionRequest) GetLang() *string {
	return s.Lang
}

func (s *UpgradeKernelVersionRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *UpgradeKernelVersionRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *UpgradeKernelVersionRequest) GetSwitchTime() *int64 {
	return s.SwitchTime
}

func (s *UpgradeKernelVersionRequest) GetUpgradeTime() *string {
	return s.UpgradeTime
}

func (s *UpgradeKernelVersionRequest) SetInstanceId(v string) *UpgradeKernelVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetKernelVersion(v string) *UpgradeKernelVersionRequest {
	s.KernelVersion = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetLang(v string) *UpgradeKernelVersionRequest {
	s.Lang = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetProductCode(v string) *UpgradeKernelVersionRequest {
	s.ProductCode = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetProductId(v int64) *UpgradeKernelVersionRequest {
	s.ProductId = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetSwitchTime(v int64) *UpgradeKernelVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetUpgradeTime(v string) *UpgradeKernelVersionRequest {
	s.UpgradeTime = &v
	return s
}

func (s *UpgradeKernelVersionRequest) Validate() error {
	return dara.Validate(s)
}
