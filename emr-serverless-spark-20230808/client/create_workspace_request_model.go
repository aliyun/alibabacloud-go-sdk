// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *CreateWorkspaceRequest
	GetAutoRenew() *string
	SetAutoRenewPeriod(v string) *CreateWorkspaceRequest
	GetAutoRenewPeriod() *string
	SetAutoRenewPeriodUnit(v string) *CreateWorkspaceRequest
	GetAutoRenewPeriodUnit() *string
	SetAutoStartSessionCluster(v bool) *CreateWorkspaceRequest
	GetAutoStartSessionCluster() *bool
	SetClientToken(v string) *CreateWorkspaceRequest
	GetClientToken() *string
	SetDlfCatalogId(v string) *CreateWorkspaceRequest
	GetDlfCatalogId() *string
	SetDlfType(v string) *CreateWorkspaceRequest
	GetDlfType() *string
	SetDuration(v string) *CreateWorkspaceRequest
	GetDuration() *string
	SetOssBucket(v string) *CreateWorkspaceRequest
	GetOssBucket() *string
	SetPaymentDurationUnit(v string) *CreateWorkspaceRequest
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *CreateWorkspaceRequest
	GetPaymentType() *string
	SetRamRoleName(v string) *CreateWorkspaceRequest
	GetRamRoleName() *string
	SetReleaseType(v string) *CreateWorkspaceRequest
	GetReleaseType() *string
	SetResourceSpec(v *CreateWorkspaceRequestResourceSpec) *CreateWorkspaceRequest
	GetResourceSpec() *CreateWorkspaceRequestResourceSpec
	SetTag(v []*CreateWorkspaceRequestTag) *CreateWorkspaceRequest
	GetTag() []*CreateWorkspaceRequestTag
	SetWorkspaceName(v string) *CreateWorkspaceRequest
	GetWorkspaceName() *string
	SetRegionId(v string) *CreateWorkspaceRequest
	GetRegionId() *string
}

type CreateWorkspaceRequest struct {
	// Specifies whether to enable auto-renewal. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The auto-renewal duration. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// 100
	AutoRenewPeriod *string `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// The unit of the auto-renewal duration. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// month
	AutoRenewPeriodUnit *string `json:"autoRenewPeriodUnit,omitempty" xml:"autoRenewPeriodUnit,omitempty"`
	// Specifies whether to automatically start a session.
	//
	// example:
	//
	// false
	AutoStartSessionCluster *bool `json:"autoStartSessionCluster,omitempty" xml:"autoStartSessionCluster,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 8e6aae2810c8f67229ca70bb31cd****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The information of the Data Lake Formation (DLF) catalog.
	//
	// example:
	//
	// 123xxxxx
	DlfCatalogId *string `json:"dlfCatalogId,omitempty" xml:"dlfCatalogId,omitempty"`
	// The version of DLF.
	//
	// example:
	//
	// dlf1.0
	DlfType *string `json:"dlfType,omitempty" xml:"dlfType,omitempty"`
	// The subscription period. This parameter is required only if the paymentType parameter is set to Pre.
	//
	// example:
	//
	// 12452
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// The name of the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// oss://test-bucket/
	OssBucket *string `json:"ossBucket,omitempty" xml:"ossBucket,omitempty"`
	// The unit of the subscription duration.
	//
	// example:
	//
	// 1000
	PaymentDurationUnit *string `json:"paymentDurationUnit,omitempty" xml:"paymentDurationUnit,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Pre
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The name of the role used to run Spark jobs.
	//
	// example:
	//
	// AliyunEMRSparkJobRunDefaultRole
	RamRoleName *string `json:"ramRoleName,omitempty" xml:"ramRoleName,omitempty"`
	// The type of the version.
	//
	// example:
	//
	// pro
	ReleaseType *string `json:"releaseType,omitempty" xml:"releaseType,omitempty"`
	// The resource specifications.
	ResourceSpec *CreateWorkspaceRequestResourceSpec `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty" type:"Struct"`
	// if can be null:
	// false
	Tag []*CreateWorkspaceRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// The name of the workspace.
	//
	// example:
	//
	// default
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateWorkspaceRequest) GetAutoRenewPeriod() *string {
	return s.AutoRenewPeriod
}

func (s *CreateWorkspaceRequest) GetAutoRenewPeriodUnit() *string {
	return s.AutoRenewPeriodUnit
}

func (s *CreateWorkspaceRequest) GetAutoStartSessionCluster() *bool {
	return s.AutoStartSessionCluster
}

func (s *CreateWorkspaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateWorkspaceRequest) GetDlfCatalogId() *string {
	return s.DlfCatalogId
}

func (s *CreateWorkspaceRequest) GetDlfType() *string {
	return s.DlfType
}

func (s *CreateWorkspaceRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateWorkspaceRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *CreateWorkspaceRequest) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *CreateWorkspaceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateWorkspaceRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateWorkspaceRequest) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *CreateWorkspaceRequest) GetResourceSpec() *CreateWorkspaceRequestResourceSpec {
	return s.ResourceSpec
}

func (s *CreateWorkspaceRequest) GetTag() []*CreateWorkspaceRequestTag {
	return s.Tag
}

func (s *CreateWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWorkspaceRequest) SetAutoRenew(v string) *CreateWorkspaceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateWorkspaceRequest) SetAutoRenewPeriod(v string) *CreateWorkspaceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateWorkspaceRequest) SetAutoRenewPeriodUnit(v string) *CreateWorkspaceRequest {
	s.AutoRenewPeriodUnit = &v
	return s
}

func (s *CreateWorkspaceRequest) SetAutoStartSessionCluster(v bool) *CreateWorkspaceRequest {
	s.AutoStartSessionCluster = &v
	return s
}

func (s *CreateWorkspaceRequest) SetClientToken(v string) *CreateWorkspaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDlfCatalogId(v string) *CreateWorkspaceRequest {
	s.DlfCatalogId = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDlfType(v string) *CreateWorkspaceRequest {
	s.DlfType = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDuration(v string) *CreateWorkspaceRequest {
	s.Duration = &v
	return s
}

func (s *CreateWorkspaceRequest) SetOssBucket(v string) *CreateWorkspaceRequest {
	s.OssBucket = &v
	return s
}

func (s *CreateWorkspaceRequest) SetPaymentDurationUnit(v string) *CreateWorkspaceRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateWorkspaceRequest) SetPaymentType(v string) *CreateWorkspaceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateWorkspaceRequest) SetRamRoleName(v string) *CreateWorkspaceRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateWorkspaceRequest) SetReleaseType(v string) *CreateWorkspaceRequest {
	s.ReleaseType = &v
	return s
}

func (s *CreateWorkspaceRequest) SetResourceSpec(v *CreateWorkspaceRequestResourceSpec) *CreateWorkspaceRequest {
	s.ResourceSpec = v
	return s
}

func (s *CreateWorkspaceRequest) SetTag(v []*CreateWorkspaceRequestTag) *CreateWorkspaceRequest {
	s.Tag = v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspaceName(v string) *CreateWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreateWorkspaceRequest) SetRegionId(v string) *CreateWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWorkspaceRequest) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWorkspaceRequestResourceSpec struct {
	// The maximum resource quota for a workspace.
	//
	// example:
	//
	// 1000
	Cu *string `json:"cu,omitempty" xml:"cu,omitempty"`
}

func (s CreateWorkspaceRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequestResourceSpec) GetCu() *string {
	return s.Cu
}

func (s *CreateWorkspaceRequestResourceSpec) SetCu(v string) *CreateWorkspaceRequestResourceSpec {
	s.Cu = &v
	return s
}

func (s *CreateWorkspaceRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}

type CreateWorkspaceRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateWorkspaceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateWorkspaceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateWorkspaceRequestTag) SetKey(v string) *CreateWorkspaceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateWorkspaceRequestTag) SetValue(v string) *CreateWorkspaceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateWorkspaceRequestTag) Validate() error {
	return dara.Validate(s)
}
