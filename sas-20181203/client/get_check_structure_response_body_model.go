// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckStructureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckStructureResponse(v []*GetCheckStructureResponseBodyCheckStructureResponse) *GetCheckStructureResponseBody
	GetCheckStructureResponse() []*GetCheckStructureResponseBodyCheckStructureResponse
	SetRequestId(v string) *GetCheckStructureResponseBody
	GetRequestId() *string
}

type GetCheckStructureResponseBody struct {
	// The structure information about check items provided by the configuration assessment feature.
	CheckStructureResponse []*GetCheckStructureResponseBodyCheckStructureResponse `json:"CheckStructureResponse,omitempty" xml:"CheckStructureResponse,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 379a9b8f-107b-4630-9e95-2299a1ea****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCheckStructureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckStructureResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckStructureResponseBody) GetCheckStructureResponse() []*GetCheckStructureResponseBodyCheckStructureResponse {
	return s.CheckStructureResponse
}

func (s *GetCheckStructureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckStructureResponseBody) SetCheckStructureResponse(v []*GetCheckStructureResponseBodyCheckStructureResponse) *GetCheckStructureResponseBody {
	s.CheckStructureResponse = v
	return s
}

func (s *GetCheckStructureResponseBody) SetRequestId(v string) *GetCheckStructureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckStructureResponseBody) Validate() error {
	if s.CheckStructureResponse != nil {
		for _, item := range s.CheckStructureResponse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCheckStructureResponseBodyCheckStructureResponse struct {
	// The type of the check item.
	//
	// 	- RISK: security risk.
	//
	// 	- IDENTITY_PERMISSION: Cloud Infrastructure Entitlement Management (CIEM).
	//
	// 	- COMPLIANCE: security compliance.
	//
	// example:
	//
	// RISK
	StandardType *string `json:"StandardType,omitempty" xml:"StandardType,omitempty"`
	// The structure information about the check items of the business type.
	Standards []*GetCheckStructureResponseBodyCheckStructureResponseStandards `json:"Standards,omitempty" xml:"Standards,omitempty" type:"Repeated"`
}

func (s GetCheckStructureResponseBodyCheckStructureResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckStructureResponseBodyCheckStructureResponse) GoString() string {
	return s.String()
}

func (s *GetCheckStructureResponseBodyCheckStructureResponse) GetStandardType() *string {
	return s.StandardType
}

func (s *GetCheckStructureResponseBodyCheckStructureResponse) GetStandards() []*GetCheckStructureResponseBodyCheckStructureResponseStandards {
	return s.Standards
}

func (s *GetCheckStructureResponseBodyCheckStructureResponse) SetStandardType(v string) *GetCheckStructureResponseBodyCheckStructureResponse {
	s.StandardType = &v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponse) SetStandards(v []*GetCheckStructureResponseBodyCheckStructureResponseStandards) *GetCheckStructureResponseBodyCheckStructureResponse {
	s.Standards = v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponse) Validate() error {
	if s.Standards != nil {
		for _, item := range s.Standards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCheckStructureResponseBodyCheckStructureResponseStandards struct {
	// The standard ID of the check item.
	//
	// example:
	//
	// 8
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The standards of the check items.
	Requirements []*GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements `json:"Requirements,omitempty" xml:"Requirements,omitempty" type:"Repeated"`
	// The display name of the standard for the check item.
	//
	// example:
	//
	// Alibaba Cloud best security practices
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The standard type of the check item. Valid values:
	//
	// 	- RISK: security risk.
	//
	// 	- IDENTITY_PERMISSION: CIEM.
	//
	// 	- COMPLIANCE: security compliance.
	//
	// example:
	//
	// IDENTITY_PERMISSION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCheckStructureResponseBodyCheckStructureResponseStandards) String() string {
	return dara.Prettify(s)
}

func (s GetCheckStructureResponseBodyCheckStructureResponseStandards) GoString() string {
	return s.String()
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandards) GetId() *int64 {
	return s.Id
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandards) GetRequirements() []*GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements {
	return s.Requirements
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandards) GetShowName() *string {
	return s.ShowName
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandards) GetType() *string {
	return s.Type
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandards) SetId(v int64) *GetCheckStructureResponseBodyCheckStructureResponseStandards {
	s.Id = &v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandards) SetRequirements(v []*GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) *GetCheckStructureResponseBodyCheckStructureResponseStandards {
	s.Requirements = v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandards) SetShowName(v string) *GetCheckStructureResponseBodyCheckStructureResponseStandards {
	s.ShowName = &v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandards) SetType(v string) *GetCheckStructureResponseBodyCheckStructureResponseStandards {
	s.Type = &v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandards) Validate() error {
	if s.Requirements != nil {
		for _, item := range s.Requirements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements struct {
	// The ID of the requirement item for the check item.
	//
	// example:
	//
	// 46
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the sections of check items.
	Sections []*GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections `json:"Sections,omitempty" xml:"Sections,omitempty" type:"Repeated"`
	// The display name of the requirement item for the check item.
	//
	// example:
	//
	// Networking
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The total number of check items for the requirement.
	//
	// example:
	//
	// 36
	TotalCheckCount *int32 `json:"TotalCheckCount,omitempty" xml:"TotalCheckCount,omitempty"`
}

func (s GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) String() string {
	return dara.Prettify(s)
}

func (s GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) GoString() string {
	return s.String()
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) GetId() *int64 {
	return s.Id
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) GetSections() []*GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections {
	return s.Sections
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) GetShowName() *string {
	return s.ShowName
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) GetTotalCheckCount() *int32 {
	return s.TotalCheckCount
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) SetId(v int64) *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements {
	s.Id = &v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) SetSections(v []*GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections) *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements {
	s.Sections = v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) SetShowName(v string) *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements {
	s.ShowName = &v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) SetTotalCheckCount(v int32) *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements {
	s.TotalCheckCount = &v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirements) Validate() error {
	if s.Sections != nil {
		for _, item := range s.Sections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections struct {
	// The ID of the section for the check item.
	//
	// example:
	//
	// 177
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The display name of the section for the check item.
	//
	// example:
	//
	// Access Control
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections) String() string {
	return dara.Prettify(s)
}

func (s GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections) GoString() string {
	return s.String()
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections) GetId() *int64 {
	return s.Id
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections) GetShowName() *string {
	return s.ShowName
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections) SetId(v int64) *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections {
	s.Id = &v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections) SetShowName(v string) *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections {
	s.ShowName = &v
	return s
}

func (s *GetCheckStructureResponseBodyCheckStructureResponseStandardsRequirementsSections) Validate() error {
	return dara.Validate(s)
}
