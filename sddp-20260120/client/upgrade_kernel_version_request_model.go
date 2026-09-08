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
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// rds_20220731
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1893456000000
	SwitchTime *int64 `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
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
