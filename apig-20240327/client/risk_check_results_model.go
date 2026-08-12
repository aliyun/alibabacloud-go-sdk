// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRiskCheckResults interface {
	dara.Model
	String() string
	GoString() string
	SetCheckTime(v int64) *RiskCheckResults
	GetCheckTime() *int64
	SetGatewayId(v string) *RiskCheckResults
	GetGatewayId() *string
	SetMetadata(v *RiskCheckResultsMetadata) *RiskCheckResults
	GetMetadata() *RiskCheckResultsMetadata
	SetRiskDetails(v []*RiskCheckResultsRiskDetails) *RiskCheckResults
	GetRiskDetails() []*RiskCheckResultsRiskDetails
	SetRiskLevel(v string) *RiskCheckResults
	GetRiskLevel() *string
	SetScore(v int32) *RiskCheckResults
	GetScore() *int32
	SetSnapshotTime(v int64) *RiskCheckResults
	GetSnapshotTime() *int64
	SetStatus(v string) *RiskCheckResults
	GetStatus() *string
	SetTotalRisk(v int32) *RiskCheckResults
	GetTotalRisk() *int32
}

type RiskCheckResults struct {
	// The check time
	//
	// example:
	//
	// 1234567890000
	CheckTime *int64 `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	// The gateway ID
	//
	// example:
	//
	// gw-0364f863b1a04474911b48cd6d51d03d
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The instance metadata
	Metadata *RiskCheckResultsMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The risk details list
	RiskDetails []*RiskCheckResultsRiskDetails `json:"riskDetails,omitempty" xml:"riskDetails,omitempty" type:"Repeated"`
	// The risk level
	//
	// example:
	//
	// MEDIUM
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// The risk score
	//
	// example:
	//
	// 85
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
	// The snapshot time
	//
	// example:
	//
	// 1234567890000
	SnapshotTime *int64 `json:"snapshotTime,omitempty" xml:"snapshotTime,omitempty"`
	// The execution status
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total number of risks
	//
	// example:
	//
	// 1
	TotalRisk *int32 `json:"totalRisk,omitempty" xml:"totalRisk,omitempty"`
}

func (s RiskCheckResults) String() string {
	return dara.Prettify(s)
}

func (s RiskCheckResults) GoString() string {
	return s.String()
}

func (s *RiskCheckResults) GetCheckTime() *int64 {
	return s.CheckTime
}

func (s *RiskCheckResults) GetGatewayId() *string {
	return s.GatewayId
}

func (s *RiskCheckResults) GetMetadata() *RiskCheckResultsMetadata {
	return s.Metadata
}

func (s *RiskCheckResults) GetRiskDetails() []*RiskCheckResultsRiskDetails {
	return s.RiskDetails
}

func (s *RiskCheckResults) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *RiskCheckResults) GetScore() *int32 {
	return s.Score
}

func (s *RiskCheckResults) GetSnapshotTime() *int64 {
	return s.SnapshotTime
}

func (s *RiskCheckResults) GetStatus() *string {
	return s.Status
}

func (s *RiskCheckResults) GetTotalRisk() *int32 {
	return s.TotalRisk
}

func (s *RiskCheckResults) SetCheckTime(v int64) *RiskCheckResults {
	s.CheckTime = &v
	return s
}

func (s *RiskCheckResults) SetGatewayId(v string) *RiskCheckResults {
	s.GatewayId = &v
	return s
}

func (s *RiskCheckResults) SetMetadata(v *RiskCheckResultsMetadata) *RiskCheckResults {
	s.Metadata = v
	return s
}

func (s *RiskCheckResults) SetRiskDetails(v []*RiskCheckResultsRiskDetails) *RiskCheckResults {
	s.RiskDetails = v
	return s
}

func (s *RiskCheckResults) SetRiskLevel(v string) *RiskCheckResults {
	s.RiskLevel = &v
	return s
}

func (s *RiskCheckResults) SetScore(v int32) *RiskCheckResults {
	s.Score = &v
	return s
}

func (s *RiskCheckResults) SetSnapshotTime(v int64) *RiskCheckResults {
	s.SnapshotTime = &v
	return s
}

func (s *RiskCheckResults) SetStatus(v string) *RiskCheckResults {
	s.Status = &v
	return s
}

func (s *RiskCheckResults) SetTotalRisk(v int32) *RiskCheckResults {
	s.TotalRisk = &v
	return s
}

func (s *RiskCheckResults) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	if s.RiskDetails != nil {
		for _, item := range s.RiskDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RiskCheckResultsMetadata struct {
	// The cluster type
	//
	// example:
	//
	// Ingress
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	// The replica count
	//
	// example:
	//
	// 3
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// The specification
	//
	// example:
	//
	// apigw.small.x1
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// The version
	//
	// example:
	//
	// 2.0.14
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s RiskCheckResultsMetadata) String() string {
	return dara.Prettify(s)
}

func (s RiskCheckResultsMetadata) GoString() string {
	return s.String()
}

func (s *RiskCheckResultsMetadata) GetClusterType() *string {
	return s.ClusterType
}

func (s *RiskCheckResultsMetadata) GetReplica() *int32 {
	return s.Replica
}

func (s *RiskCheckResultsMetadata) GetSpec() *string {
	return s.Spec
}

func (s *RiskCheckResultsMetadata) GetVersion() *string {
	return s.Version
}

func (s *RiskCheckResultsMetadata) SetClusterType(v string) *RiskCheckResultsMetadata {
	s.ClusterType = &v
	return s
}

func (s *RiskCheckResultsMetadata) SetReplica(v int32) *RiskCheckResultsMetadata {
	s.Replica = &v
	return s
}

func (s *RiskCheckResultsMetadata) SetSpec(v string) *RiskCheckResultsMetadata {
	s.Spec = &v
	return s
}

func (s *RiskCheckResultsMetadata) SetVersion(v string) *RiskCheckResultsMetadata {
	s.Version = &v
	return s
}

func (s *RiskCheckResultsMetadata) Validate() error {
	return dara.Validate(s)
}

type RiskCheckResultsRiskDetails struct {
	// The check module
	//
	// example:
	//
	// BaseInfo
	CheckModule *string `json:"checkModule,omitempty" xml:"checkModule,omitempty"`
	// The risk detailed data
	Data map[string]*string `json:"data,omitempty" xml:"data,omitempty"`
	// The risk description
	//
	// example:
	//
	// The version is outdated
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Whether to mute notifications
	//
	// example:
	//
	// false
	IsNoticeMute *bool `json:"isNoticeMute,omitempty" xml:"isNoticeMute,omitempty"`
	// The risk code
	//
	// example:
	//
	// 30010010001
	RiskCode *string `json:"riskCode,omitempty" xml:"riskCode,omitempty"`
	// The risk level
	//
	// example:
	//
	// HIGH
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// The risk title
	//
	// example:
	//
	// Version Risk
	RiskName *string `json:"riskName,omitempty" xml:"riskName,omitempty"`
	// The risk type
	//
	// example:
	//
	// SYSTEM
	RiskType *string `json:"riskType,omitempty" xml:"riskType,omitempty"`
	// The risk situation
	//
	// example:
	//
	// Current version is 2 releases behind
	Situation *string `json:"situation,omitempty" xml:"situation,omitempty"`
	// The fix suggestion
	//
	// example:
	//
	// Upgrade to the latest version
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
}

func (s RiskCheckResultsRiskDetails) String() string {
	return dara.Prettify(s)
}

func (s RiskCheckResultsRiskDetails) GoString() string {
	return s.String()
}

func (s *RiskCheckResultsRiskDetails) GetCheckModule() *string {
	return s.CheckModule
}

func (s *RiskCheckResultsRiskDetails) GetData() map[string]*string {
	return s.Data
}

func (s *RiskCheckResultsRiskDetails) GetDescription() *string {
	return s.Description
}

func (s *RiskCheckResultsRiskDetails) GetIsNoticeMute() *bool {
	return s.IsNoticeMute
}

func (s *RiskCheckResultsRiskDetails) GetRiskCode() *string {
	return s.RiskCode
}

func (s *RiskCheckResultsRiskDetails) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *RiskCheckResultsRiskDetails) GetRiskName() *string {
	return s.RiskName
}

func (s *RiskCheckResultsRiskDetails) GetRiskType() *string {
	return s.RiskType
}

func (s *RiskCheckResultsRiskDetails) GetSituation() *string {
	return s.Situation
}

func (s *RiskCheckResultsRiskDetails) GetSuggestion() *string {
	return s.Suggestion
}

func (s *RiskCheckResultsRiskDetails) SetCheckModule(v string) *RiskCheckResultsRiskDetails {
	s.CheckModule = &v
	return s
}

func (s *RiskCheckResultsRiskDetails) SetData(v map[string]*string) *RiskCheckResultsRiskDetails {
	s.Data = v
	return s
}

func (s *RiskCheckResultsRiskDetails) SetDescription(v string) *RiskCheckResultsRiskDetails {
	s.Description = &v
	return s
}

func (s *RiskCheckResultsRiskDetails) SetIsNoticeMute(v bool) *RiskCheckResultsRiskDetails {
	s.IsNoticeMute = &v
	return s
}

func (s *RiskCheckResultsRiskDetails) SetRiskCode(v string) *RiskCheckResultsRiskDetails {
	s.RiskCode = &v
	return s
}

func (s *RiskCheckResultsRiskDetails) SetRiskLevel(v string) *RiskCheckResultsRiskDetails {
	s.RiskLevel = &v
	return s
}

func (s *RiskCheckResultsRiskDetails) SetRiskName(v string) *RiskCheckResultsRiskDetails {
	s.RiskName = &v
	return s
}

func (s *RiskCheckResultsRiskDetails) SetRiskType(v string) *RiskCheckResultsRiskDetails {
	s.RiskType = &v
	return s
}

func (s *RiskCheckResultsRiskDetails) SetSituation(v string) *RiskCheckResultsRiskDetails {
	s.Situation = &v
	return s
}

func (s *RiskCheckResultsRiskDetails) SetSuggestion(v string) *RiskCheckResultsRiskDetails {
	s.Suggestion = &v
	return s
}

func (s *RiskCheckResultsRiskDetails) Validate() error {
	return dara.Validate(s)
}
