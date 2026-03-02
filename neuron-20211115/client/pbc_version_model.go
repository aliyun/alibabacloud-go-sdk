// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcVersion interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *PbcVersion
	GetCompanyId() *int64
	SetCompanyName(v string) *PbcVersion
	GetCompanyName() *string
	SetDeveloperId(v int64) *PbcVersion
	GetDeveloperId() *int64
	SetId(v int64) *PbcVersion
	GetId() *int64
	SetIsWatched(v bool) *PbcVersion
	GetIsWatched() *bool
	SetName(v string) *PbcVersion
	GetName() *string
	SetPbcId(v int64) *PbcVersion
	GetPbcId() *int64
	SetRequestId(v string) *PbcVersion
	GetRequestId() *string
	SetReviewId(v int64) *PbcVersion
	GetReviewId() *int64
	SetStatus(v string) *PbcVersion
	GetStatus() *string
	SetStepStatus(v string) *PbcVersion
	GetStepStatus() *string
	SetVersion(v string) *PbcVersion
	GetVersion() *string
	SetVisibilityLevel(v string) *PbcVersion
	GetVisibilityLevel() *string
}

type PbcVersion struct {
	// example:
	//
	// 1
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// 企业服务团队
	CompanyName *string `json:"companyName,omitempty" xml:"companyName,omitempty"`
	// example:
	//
	// 1
	DeveloperId *int64 `json:"developerId,omitempty" xml:"developerId,omitempty"`
	// example:
	//
	// 1
	Id        *int64 `json:"id,omitempty" xml:"id,omitempty"`
	IsWatched *bool  `json:"isWatched,omitempty" xml:"isWatched,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// product
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	PbcId     *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	ReviewId *int64 `json:"reviewId,omitempty" xml:"reviewId,omitempty"`
	// example:
	//
	// DEVELOPING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SPEC
	StepStatus *string `json:"stepStatus,omitempty" xml:"stepStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// \"2022-04-01\"
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC
	VisibilityLevel *string `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s PbcVersion) String() string {
	return dara.Prettify(s)
}

func (s PbcVersion) GoString() string {
	return s.String()
}

func (s *PbcVersion) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PbcVersion) GetCompanyName() *string {
	return s.CompanyName
}

func (s *PbcVersion) GetDeveloperId() *int64 {
	return s.DeveloperId
}

func (s *PbcVersion) GetId() *int64 {
	return s.Id
}

func (s *PbcVersion) GetIsWatched() *bool {
	return s.IsWatched
}

func (s *PbcVersion) GetName() *string {
	return s.Name
}

func (s *PbcVersion) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PbcVersion) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcVersion) GetReviewId() *int64 {
	return s.ReviewId
}

func (s *PbcVersion) GetStatus() *string {
	return s.Status
}

func (s *PbcVersion) GetStepStatus() *string {
	return s.StepStatus
}

func (s *PbcVersion) GetVersion() *string {
	return s.Version
}

func (s *PbcVersion) GetVisibilityLevel() *string {
	return s.VisibilityLevel
}

func (s *PbcVersion) SetCompanyId(v int64) *PbcVersion {
	s.CompanyId = &v
	return s
}

func (s *PbcVersion) SetCompanyName(v string) *PbcVersion {
	s.CompanyName = &v
	return s
}

func (s *PbcVersion) SetDeveloperId(v int64) *PbcVersion {
	s.DeveloperId = &v
	return s
}

func (s *PbcVersion) SetId(v int64) *PbcVersion {
	s.Id = &v
	return s
}

func (s *PbcVersion) SetIsWatched(v bool) *PbcVersion {
	s.IsWatched = &v
	return s
}

func (s *PbcVersion) SetName(v string) *PbcVersion {
	s.Name = &v
	return s
}

func (s *PbcVersion) SetPbcId(v int64) *PbcVersion {
	s.PbcId = &v
	return s
}

func (s *PbcVersion) SetRequestId(v string) *PbcVersion {
	s.RequestId = &v
	return s
}

func (s *PbcVersion) SetReviewId(v int64) *PbcVersion {
	s.ReviewId = &v
	return s
}

func (s *PbcVersion) SetStatus(v string) *PbcVersion {
	s.Status = &v
	return s
}

func (s *PbcVersion) SetStepStatus(v string) *PbcVersion {
	s.StepStatus = &v
	return s
}

func (s *PbcVersion) SetVersion(v string) *PbcVersion {
	s.Version = &v
	return s
}

func (s *PbcVersion) SetVisibilityLevel(v string) *PbcVersion {
	s.VisibilityLevel = &v
	return s
}

func (s *PbcVersion) Validate() error {
	return dara.Validate(s)
}
