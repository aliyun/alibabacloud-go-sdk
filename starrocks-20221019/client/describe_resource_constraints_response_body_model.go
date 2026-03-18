// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceConstraintsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeResourceConstraintsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *DescribeResourceConstraintsResponseBodyData) *DescribeResourceConstraintsResponseBody
	GetData() *DescribeResourceConstraintsResponseBodyData
	SetErrMessage(v string) *DescribeResourceConstraintsResponseBody
	GetErrMessage() *string
	SetErrorCode(v string) *DescribeResourceConstraintsResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DescribeResourceConstraintsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeResourceConstraintsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeResourceConstraintsResponseBody
	GetSuccess() *bool
}

type DescribeResourceConstraintsResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                      `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *DescribeResourceConstraintsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// B67D142D-D54E-184F-A306-22BDC01B2XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeResourceConstraintsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeResourceConstraintsResponseBody) GetData() *DescribeResourceConstraintsResponseBodyData {
	return s.Data
}

func (s *DescribeResourceConstraintsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeResourceConstraintsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeResourceConstraintsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeResourceConstraintsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceConstraintsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeResourceConstraintsResponseBody) SetAccessDeniedDetail(v string) *DescribeResourceConstraintsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBody) SetData(v *DescribeResourceConstraintsResponseBodyData) *DescribeResourceConstraintsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeResourceConstraintsResponseBody) SetErrMessage(v string) *DescribeResourceConstraintsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBody) SetErrorCode(v string) *DescribeResourceConstraintsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBody) SetHttpStatusCode(v int32) *DescribeResourceConstraintsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBody) SetRequestId(v string) *DescribeResourceConstraintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBody) SetSuccess(v bool) *DescribeResourceConstraintsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourceConstraintsResponseBodyData struct {
	AgentCu                         []*int32                                                                      `json:"AgentCu,omitempty" xml:"AgentCu,omitempty" type:"Repeated"`
	BeCu                            []*int32                                                                      `json:"BeCu,omitempty" xml:"BeCu,omitempty" type:"Repeated"`
	BeCuOnEcs                       []*int32                                                                      `json:"BeCuOnEcs,omitempty" xml:"BeCuOnEcs,omitempty" type:"Repeated"`
	BeNumber                        *DescribeResourceConstraintsResponseBodyDataBeNumber                          `json:"BeNumber,omitempty" xml:"BeNumber,omitempty" type:"Struct"`
	BeStorageConstraints            []*DescribeResourceConstraintsResponseBodyDataBeStorageConstraints            `json:"BeStorageConstraints,omitempty" xml:"BeStorageConstraints,omitempty" type:"Repeated"`
	BigDataInstanceTypeConstraints  []*DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints  `json:"BigDataInstanceTypeConstraints,omitempty" xml:"BigDataInstanceTypeConstraints,omitempty" type:"Repeated"`
	FeCu                            []*int32                                                                      `json:"FeCu,omitempty" xml:"FeCu,omitempty" type:"Repeated"`
	FeCuOnEcs                       []*int32                                                                      `json:"FeCuOnEcs,omitempty" xml:"FeCuOnEcs,omitempty" type:"Repeated"`
	FeNumber                        *DescribeResourceConstraintsResponseBodyDataFeNumber                          `json:"FeNumber,omitempty" xml:"FeNumber,omitempty" type:"Struct"`
	FeSpecType                      []*DescribeResourceConstraintsResponseBodyDataFeSpecType                      `json:"FeSpecType,omitempty" xml:"FeSpecType,omitempty" type:"Repeated"`
	FeStorage                       *DescribeResourceConstraintsResponseBodyDataFeStorage                         `json:"FeStorage,omitempty" xml:"FeStorage,omitempty" type:"Struct"`
	HaFeResourceSpec                *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec                  `json:"HaFeResourceSpec,omitempty" xml:"HaFeResourceSpec,omitempty" type:"Struct"`
	LocalSSDInstanceTypeConstraints []*DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints `json:"LocalSSDInstanceTypeConstraints,omitempty" xml:"LocalSSDInstanceTypeConstraints,omitempty" type:"Repeated"`
	NormalFeResourceSpec            *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec              `json:"NormalFeResourceSpec,omitempty" xml:"NormalFeResourceSpec,omitempty" type:"Struct"`
	// example:
	//
	// standard
	SpecType                      []*DescribeResourceConstraintsResponseBodyDataSpecType                    `json:"SpecType,omitempty" xml:"SpecType,omitempty" type:"Repeated"`
	SplitDiskThresholdMap         map[string]map[string]interface{}                                         `json:"SplitDiskThresholdMap,omitempty" xml:"SplitDiskThresholdMap,omitempty"`
	VersionConstraint             *DescribeResourceConstraintsResponseBodyDataVersionConstraint             `json:"VersionConstraint,omitempty" xml:"VersionConstraint,omitempty" type:"Struct"`
	ZoneSupportedEedTypes         map[string][]*string                                                      `json:"ZoneSupportedEedTypes,omitempty" xml:"ZoneSupportedEedTypes,omitempty"`
	ZoneSupportedSpecTypes        map[string][]*string                                                      `json:"ZoneSupportedSpecTypes,omitempty" xml:"ZoneSupportedSpecTypes,omitempty"`
	CompactionServiceCuConstraint *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint `json:"compactionServiceCuConstraint,omitempty" xml:"compactionServiceCuConstraint,omitempty" type:"Struct"`
}

func (s DescribeResourceConstraintsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyData) GetAgentCu() []*int32 {
	return s.AgentCu
}

func (s *DescribeResourceConstraintsResponseBodyData) GetBeCu() []*int32 {
	return s.BeCu
}

func (s *DescribeResourceConstraintsResponseBodyData) GetBeCuOnEcs() []*int32 {
	return s.BeCuOnEcs
}

func (s *DescribeResourceConstraintsResponseBodyData) GetBeNumber() *DescribeResourceConstraintsResponseBodyDataBeNumber {
	return s.BeNumber
}

func (s *DescribeResourceConstraintsResponseBodyData) GetBeStorageConstraints() []*DescribeResourceConstraintsResponseBodyDataBeStorageConstraints {
	return s.BeStorageConstraints
}

func (s *DescribeResourceConstraintsResponseBodyData) GetBigDataInstanceTypeConstraints() []*DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints {
	return s.BigDataInstanceTypeConstraints
}

func (s *DescribeResourceConstraintsResponseBodyData) GetFeCu() []*int32 {
	return s.FeCu
}

func (s *DescribeResourceConstraintsResponseBodyData) GetFeCuOnEcs() []*int32 {
	return s.FeCuOnEcs
}

func (s *DescribeResourceConstraintsResponseBodyData) GetFeNumber() *DescribeResourceConstraintsResponseBodyDataFeNumber {
	return s.FeNumber
}

func (s *DescribeResourceConstraintsResponseBodyData) GetFeSpecType() []*DescribeResourceConstraintsResponseBodyDataFeSpecType {
	return s.FeSpecType
}

func (s *DescribeResourceConstraintsResponseBodyData) GetFeStorage() *DescribeResourceConstraintsResponseBodyDataFeStorage {
	return s.FeStorage
}

func (s *DescribeResourceConstraintsResponseBodyData) GetHaFeResourceSpec() *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec {
	return s.HaFeResourceSpec
}

func (s *DescribeResourceConstraintsResponseBodyData) GetLocalSSDInstanceTypeConstraints() []*DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints {
	return s.LocalSSDInstanceTypeConstraints
}

func (s *DescribeResourceConstraintsResponseBodyData) GetNormalFeResourceSpec() *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec {
	return s.NormalFeResourceSpec
}

func (s *DescribeResourceConstraintsResponseBodyData) GetSpecType() []*DescribeResourceConstraintsResponseBodyDataSpecType {
	return s.SpecType
}

func (s *DescribeResourceConstraintsResponseBodyData) GetSplitDiskThresholdMap() map[string]map[string]interface{} {
	return s.SplitDiskThresholdMap
}

func (s *DescribeResourceConstraintsResponseBodyData) GetVersionConstraint() *DescribeResourceConstraintsResponseBodyDataVersionConstraint {
	return s.VersionConstraint
}

func (s *DescribeResourceConstraintsResponseBodyData) GetZoneSupportedEedTypes() map[string][]*string {
	return s.ZoneSupportedEedTypes
}

func (s *DescribeResourceConstraintsResponseBodyData) GetZoneSupportedSpecTypes() map[string][]*string {
	return s.ZoneSupportedSpecTypes
}

func (s *DescribeResourceConstraintsResponseBodyData) GetCompactionServiceCuConstraint() *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint {
	return s.CompactionServiceCuConstraint
}

func (s *DescribeResourceConstraintsResponseBodyData) SetAgentCu(v []*int32) *DescribeResourceConstraintsResponseBodyData {
	s.AgentCu = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetBeCu(v []*int32) *DescribeResourceConstraintsResponseBodyData {
	s.BeCu = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetBeCuOnEcs(v []*int32) *DescribeResourceConstraintsResponseBodyData {
	s.BeCuOnEcs = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetBeNumber(v *DescribeResourceConstraintsResponseBodyDataBeNumber) *DescribeResourceConstraintsResponseBodyData {
	s.BeNumber = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetBeStorageConstraints(v []*DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) *DescribeResourceConstraintsResponseBodyData {
	s.BeStorageConstraints = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetBigDataInstanceTypeConstraints(v []*DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) *DescribeResourceConstraintsResponseBodyData {
	s.BigDataInstanceTypeConstraints = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetFeCu(v []*int32) *DescribeResourceConstraintsResponseBodyData {
	s.FeCu = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetFeCuOnEcs(v []*int32) *DescribeResourceConstraintsResponseBodyData {
	s.FeCuOnEcs = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetFeNumber(v *DescribeResourceConstraintsResponseBodyDataFeNumber) *DescribeResourceConstraintsResponseBodyData {
	s.FeNumber = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetFeSpecType(v []*DescribeResourceConstraintsResponseBodyDataFeSpecType) *DescribeResourceConstraintsResponseBodyData {
	s.FeSpecType = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetFeStorage(v *DescribeResourceConstraintsResponseBodyDataFeStorage) *DescribeResourceConstraintsResponseBodyData {
	s.FeStorage = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetHaFeResourceSpec(v *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec) *DescribeResourceConstraintsResponseBodyData {
	s.HaFeResourceSpec = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetLocalSSDInstanceTypeConstraints(v []*DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) *DescribeResourceConstraintsResponseBodyData {
	s.LocalSSDInstanceTypeConstraints = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetNormalFeResourceSpec(v *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec) *DescribeResourceConstraintsResponseBodyData {
	s.NormalFeResourceSpec = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetSpecType(v []*DescribeResourceConstraintsResponseBodyDataSpecType) *DescribeResourceConstraintsResponseBodyData {
	s.SpecType = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetSplitDiskThresholdMap(v map[string]map[string]interface{}) *DescribeResourceConstraintsResponseBodyData {
	s.SplitDiskThresholdMap = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetVersionConstraint(v *DescribeResourceConstraintsResponseBodyDataVersionConstraint) *DescribeResourceConstraintsResponseBodyData {
	s.VersionConstraint = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetZoneSupportedEedTypes(v map[string][]*string) *DescribeResourceConstraintsResponseBodyData {
	s.ZoneSupportedEedTypes = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetZoneSupportedSpecTypes(v map[string][]*string) *DescribeResourceConstraintsResponseBodyData {
	s.ZoneSupportedSpecTypes = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) SetCompactionServiceCuConstraint(v *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) *DescribeResourceConstraintsResponseBodyData {
	s.CompactionServiceCuConstraint = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyData) Validate() error {
	if s.BeNumber != nil {
		if err := s.BeNumber.Validate(); err != nil {
			return err
		}
	}
	if s.BeStorageConstraints != nil {
		for _, item := range s.BeStorageConstraints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.BigDataInstanceTypeConstraints != nil {
		for _, item := range s.BigDataInstanceTypeConstraints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FeNumber != nil {
		if err := s.FeNumber.Validate(); err != nil {
			return err
		}
	}
	if s.FeSpecType != nil {
		for _, item := range s.FeSpecType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FeStorage != nil {
		if err := s.FeStorage.Validate(); err != nil {
			return err
		}
	}
	if s.HaFeResourceSpec != nil {
		if err := s.HaFeResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.LocalSSDInstanceTypeConstraints != nil {
		for _, item := range s.LocalSSDInstanceTypeConstraints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NormalFeResourceSpec != nil {
		if err := s.NormalFeResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.SpecType != nil {
		for _, item := range s.SpecType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VersionConstraint != nil {
		if err := s.VersionConstraint.Validate(); err != nil {
			return err
		}
	}
	if s.CompactionServiceCuConstraint != nil {
		if err := s.CompactionServiceCuConstraint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourceConstraintsResponseBodyDataBeNumber struct {
	// example:
	//
	// 1
	Default *int32 `json:"Default,omitempty" xml:"Default,omitempty"`
	// example:
	//
	// 10
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 1
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataBeNumber) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataBeNumber) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataBeNumber) GetDefault() *int32 {
	return s.Default
}

func (s *DescribeResourceConstraintsResponseBodyDataBeNumber) GetMax() *int32 {
	return s.Max
}

func (s *DescribeResourceConstraintsResponseBodyDataBeNumber) GetMin() *int32 {
	return s.Min
}

func (s *DescribeResourceConstraintsResponseBodyDataBeNumber) GetStep() *int32 {
	return s.Step
}

func (s *DescribeResourceConstraintsResponseBodyDataBeNumber) SetDefault(v int32) *DescribeResourceConstraintsResponseBodyDataBeNumber {
	s.Default = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeNumber) SetMax(v int32) *DescribeResourceConstraintsResponseBodyDataBeNumber {
	s.Max = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeNumber) SetMin(v int32) *DescribeResourceConstraintsResponseBodyDataBeNumber {
	s.Min = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeNumber) SetStep(v int32) *DescribeResourceConstraintsResponseBodyDataBeNumber {
	s.Step = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeNumber) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataBeStorageConstraints struct {
	Desc                 *string                                                                              `json:"Desc,omitempty" xml:"Desc,omitempty"`
	DiskNumberConstraint *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint `json:"DiskNumberConstraint,omitempty" xml:"DiskNumberConstraint,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// PL1
	Level           *string                                                                         `json:"Level,omitempty" xml:"Level,omitempty"`
	ValueConstraint *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint `json:"ValueConstraint,omitempty" xml:"ValueConstraint,omitempty" type:"Struct"`
}

func (s DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) GetDesc() *string {
	return s.Desc
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) GetDiskNumberConstraint() *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint {
	return s.DiskNumberConstraint
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) GetLevel() *string {
	return s.Level
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) GetValueConstraint() *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint {
	return s.ValueConstraint
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) SetDesc(v string) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints {
	s.Desc = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) SetDiskNumberConstraint(v *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints {
	s.DiskNumberConstraint = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) SetIsDefault(v bool) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints {
	s.IsDefault = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) SetLevel(v string) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints {
	s.Level = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) SetValueConstraint(v *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints {
	s.ValueConstraint = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraints) Validate() error {
	if s.DiskNumberConstraint != nil {
		if err := s.DiskNumberConstraint.Validate(); err != nil {
			return err
		}
	}
	if s.ValueConstraint != nil {
		if err := s.ValueConstraint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint struct {
	// example:
	//
	// 1
	Default *int32 `json:"Default,omitempty" xml:"Default,omitempty"`
	// example:
	//
	// 10
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 1
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) GetDefault() *int32 {
	return s.Default
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) GetMax() *int32 {
	return s.Max
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) GetMin() *int32 {
	return s.Min
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) GetStep() *int32 {
	return s.Step
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) SetDefault(v int32) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint {
	s.Default = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) SetMax(v int32) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint {
	s.Max = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) SetMin(v int32) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint {
	s.Min = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) SetStep(v int32) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint {
	s.Step = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsDiskNumberConstraint) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint struct {
	// example:
	//
	// 1
	Default *int32 `json:"Default,omitempty" xml:"Default,omitempty"`
	// example:
	//
	// 5
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 1
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) GetDefault() *int32 {
	return s.Default
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) GetMax() *int32 {
	return s.Max
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) GetMin() *int32 {
	return s.Min
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) GetStep() *int32 {
	return s.Step
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) SetDefault(v int32) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint {
	s.Default = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) SetMax(v int32) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint {
	s.Max = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) SetMin(v int32) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint {
	s.Min = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) SetStep(v int32) *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint {
	s.Step = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBeStorageConstraintsValueConstraint) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints struct {
	// example:
	//
	// 20
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 8
	DiskNumber *string `json:"DiskNumber,omitempty" xml:"DiskNumber,omitempty"`
	Display    *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// ecs.d2s.5xlarge
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
	// example:
	//
	// local_hdd_2s_5xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// false
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// 88
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 7300
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) GetDiskNumber() *string {
	return s.DiskNumber
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) GetDisplay() *string {
	return s.Display
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) GetEcsInstanceType() *string {
	return s.EcsInstanceType
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) GetIsDefault() *string {
	return s.IsDefault
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) GetStorageSize() *string {
	return s.StorageSize
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) SetCpu(v int32) *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints {
	s.Cpu = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) SetDiskNumber(v string) *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints {
	s.DiskNumber = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) SetDisplay(v string) *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints {
	s.Display = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) SetEcsInstanceType(v string) *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints {
	s.EcsInstanceType = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) SetInstanceType(v string) *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints {
	s.InstanceType = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) SetIsDefault(v string) *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints {
	s.IsDefault = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) SetMemory(v int32) *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints {
	s.Memory = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) SetStorageSize(v string) *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints {
	s.StorageSize = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataBigDataInstanceTypeConstraints) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataFeNumber struct {
	// example:
	//
	// 3
	Default *int32 `json:"Default,omitempty" xml:"Default,omitempty"`
	// example:
	//
	// 11
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 1
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// 2
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataFeNumber) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataFeNumber) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataFeNumber) GetDefault() *int32 {
	return s.Default
}

func (s *DescribeResourceConstraintsResponseBodyDataFeNumber) GetMax() *int32 {
	return s.Max
}

func (s *DescribeResourceConstraintsResponseBodyDataFeNumber) GetMin() *int32 {
	return s.Min
}

func (s *DescribeResourceConstraintsResponseBodyDataFeNumber) GetStep() *int32 {
	return s.Step
}

func (s *DescribeResourceConstraintsResponseBodyDataFeNumber) SetDefault(v int32) *DescribeResourceConstraintsResponseBodyDataFeNumber {
	s.Default = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataFeNumber) SetMax(v int32) *DescribeResourceConstraintsResponseBodyDataFeNumber {
	s.Max = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataFeNumber) SetMin(v int32) *DescribeResourceConstraintsResponseBodyDataFeNumber {
	s.Min = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataFeNumber) SetStep(v int32) *DescribeResourceConstraintsResponseBodyDataFeNumber {
	s.Step = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataFeNumber) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataFeSpecType struct {
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// standard
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataFeSpecType) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataFeSpecType) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataFeSpecType) GetDisplay() *string {
	return s.Display
}

func (s *DescribeResourceConstraintsResponseBodyDataFeSpecType) GetName() *string {
	return s.Name
}

func (s *DescribeResourceConstraintsResponseBodyDataFeSpecType) SetDisplay(v string) *DescribeResourceConstraintsResponseBodyDataFeSpecType {
	s.Display = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataFeSpecType) SetName(v string) *DescribeResourceConstraintsResponseBodyDataFeSpecType {
	s.Name = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataFeSpecType) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataFeStorage struct {
	// example:
	//
	// 500
	Default *int32 `json:"Default,omitempty" xml:"Default,omitempty"`
	// example:
	//
	// 5000
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 200
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// 100
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataFeStorage) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataFeStorage) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataFeStorage) GetDefault() *int32 {
	return s.Default
}

func (s *DescribeResourceConstraintsResponseBodyDataFeStorage) GetMax() *int32 {
	return s.Max
}

func (s *DescribeResourceConstraintsResponseBodyDataFeStorage) GetMin() *int32 {
	return s.Min
}

func (s *DescribeResourceConstraintsResponseBodyDataFeStorage) GetStep() *int32 {
	return s.Step
}

func (s *DescribeResourceConstraintsResponseBodyDataFeStorage) SetDefault(v int32) *DescribeResourceConstraintsResponseBodyDataFeStorage {
	s.Default = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataFeStorage) SetMax(v int32) *DescribeResourceConstraintsResponseBodyDataFeStorage {
	s.Max = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataFeStorage) SetMin(v int32) *DescribeResourceConstraintsResponseBodyDataFeStorage {
	s.Min = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataFeStorage) SetStep(v int32) *DescribeResourceConstraintsResponseBodyDataFeStorage {
	s.Step = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataFeStorage) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec struct {
	// example:
	//
	// 16
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// 5
	NodeNumber *int32 `json:"NodeNumber,omitempty" xml:"NodeNumber,omitempty"`
	// example:
	//
	// 100
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec) GetCu() *int32 {
	return s.Cu
}

func (s *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec) GetNodeNumber() *int32 {
	return s.NodeNumber
}

func (s *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec) SetCu(v int32) *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec {
	s.Cu = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec) SetNodeNumber(v int32) *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec {
	s.NodeNumber = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec) SetStorageSize(v int32) *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec {
	s.StorageSize = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataHaFeResourceSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints struct {
	// example:
	//
	// 16
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 1
	DiskNumber *string `json:"DiskNumber,omitempty" xml:"DiskNumber,omitempty"`
	Display    *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// ecs.i2g.4xlarge
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
	// example:
	//
	// local_ssd_2g_4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// true
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// 64
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 1788
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) GetDiskNumber() *string {
	return s.DiskNumber
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) GetDisplay() *string {
	return s.Display
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) GetEcsInstanceType() *string {
	return s.EcsInstanceType
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) GetIsDefault() *string {
	return s.IsDefault
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) GetStorageSize() *string {
	return s.StorageSize
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) SetCpu(v int32) *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints {
	s.Cpu = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) SetDiskNumber(v string) *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints {
	s.DiskNumber = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) SetDisplay(v string) *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints {
	s.Display = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) SetEcsInstanceType(v string) *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints {
	s.EcsInstanceType = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) SetInstanceType(v string) *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints {
	s.InstanceType = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) SetIsDefault(v string) *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints {
	s.IsDefault = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) SetMemory(v int32) *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints {
	s.Memory = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) SetStorageSize(v string) *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints {
	s.StorageSize = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataLocalSSDInstanceTypeConstraints) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec struct {
	// example:
	//
	// 64
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// 5
	NodeNumber *int32 `json:"NodeNumber,omitempty" xml:"NodeNumber,omitempty"`
	// example:
	//
	// 500
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec) GetCu() *int32 {
	return s.Cu
}

func (s *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec) GetNodeNumber() *int32 {
	return s.NodeNumber
}

func (s *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec) SetCu(v int32) *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec {
	s.Cu = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec) SetNodeNumber(v int32) *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec {
	s.NodeNumber = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec) SetStorageSize(v int32) *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec {
	s.StorageSize = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataNormalFeResourceSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataSpecType struct {
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// standard
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataSpecType) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataSpecType) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataSpecType) GetDisplay() *string {
	return s.Display
}

func (s *DescribeResourceConstraintsResponseBodyDataSpecType) GetName() *string {
	return s.Name
}

func (s *DescribeResourceConstraintsResponseBodyDataSpecType) SetDisplay(v string) *DescribeResourceConstraintsResponseBodyDataSpecType {
	s.Display = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataSpecType) SetName(v string) *DescribeResourceConstraintsResponseBodyDataSpecType {
	s.Name = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataSpecType) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataVersionConstraint struct {
	BetaVersions []*string `json:"BetaVersions,omitempty" xml:"BetaVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 3.3
	DefaultVersion *string   `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Versions       []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeResourceConstraintsResponseBodyDataVersionConstraint) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataVersionConstraint) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataVersionConstraint) GetBetaVersions() []*string {
	return s.BetaVersions
}

func (s *DescribeResourceConstraintsResponseBodyDataVersionConstraint) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *DescribeResourceConstraintsResponseBodyDataVersionConstraint) GetVersions() []*string {
	return s.Versions
}

func (s *DescribeResourceConstraintsResponseBodyDataVersionConstraint) SetBetaVersions(v []*string) *DescribeResourceConstraintsResponseBodyDataVersionConstraint {
	s.BetaVersions = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataVersionConstraint) SetDefaultVersion(v string) *DescribeResourceConstraintsResponseBodyDataVersionConstraint {
	s.DefaultVersion = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataVersionConstraint) SetVersions(v []*string) *DescribeResourceConstraintsResponseBodyDataVersionConstraint {
	s.Versions = v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataVersionConstraint) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint struct {
	// example:
	//
	// 16
	Def *int32 `json:"def,omitempty" xml:"def,omitempty"`
	// example:
	//
	// 256
	Max *int32 `json:"max,omitempty" xml:"max,omitempty"`
	// example:
	//
	// 8
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// example:
	//
	// 8
	Step *int32 `json:"step,omitempty" xml:"step,omitempty"`
}

func (s DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) GetDef() *int32 {
	return s.Def
}

func (s *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) GetMax() *int32 {
	return s.Max
}

func (s *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) GetMin() *int32 {
	return s.Min
}

func (s *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) GetStep() *int32 {
	return s.Step
}

func (s *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) SetDef(v int32) *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint {
	s.Def = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) SetMax(v int32) *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint {
	s.Max = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) SetMin(v int32) *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint {
	s.Min = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) SetStep(v int32) *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint {
	s.Step = &v
	return s
}

func (s *DescribeResourceConstraintsResponseBodyDataCompactionServiceCuConstraint) Validate() error {
	return dara.Validate(s)
}
