// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckStandardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListCheckStandardResponseBody
	GetRequestId() *string
	SetStandards(v []*ListCheckStandardResponseBodyStandards) *ListCheckStandardResponseBody
	GetStandards() []*ListCheckStandardResponseBodyStandards
}

type ListCheckStandardResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// FA91FBDA-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The standards.
	Standards []*ListCheckStandardResponseBodyStandards `json:"Standards,omitempty" xml:"Standards,omitempty" type:"Repeated"`
}

func (s ListCheckStandardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCheckStandardResponseBody) GoString() string {
	return s.String()
}

func (s *ListCheckStandardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCheckStandardResponseBody) GetStandards() []*ListCheckStandardResponseBodyStandards {
	return s.Standards
}

func (s *ListCheckStandardResponseBody) SetRequestId(v string) *ListCheckStandardResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCheckStandardResponseBody) SetStandards(v []*ListCheckStandardResponseBodyStandards) *ListCheckStandardResponseBody {
	s.Standards = v
	return s
}

func (s *ListCheckStandardResponseBody) Validate() error {
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

type ListCheckStandardResponseBodyStandards struct {
	// The cloud service provider that uses the standard. Valid values:
	//
	// 	- **0**: Alibaba Cloud.
	//
	// 	- **3**: Tencent Cloud.
	//
	// 	- **4**: Huawei Cloud.
	//
	// 	- **5**: Microsoft Azure.
	//
	// 	- **7**: AWS.
	//
	// example:
	//
	// 3
	BindVendor *int32 `json:"BindVendor,omitempty" xml:"BindVendor,omitempty"`
	// The ID of the standard.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The requirements.
	Requirements []*ListCheckStandardResponseBodyStandardsRequirements `json:"Requirements,omitempty" xml:"Requirements,omitempty" type:"Repeated"`
	// The display name of the check item.
	//
	// example:
	//
	// Identity and permission management
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The priority for display.
	//
	// example:
	//
	// 1
	ShowPriorityLevel *int32 `json:"ShowPriorityLevel,omitempty" xml:"ShowPriorityLevel,omitempty"`
	// The type of the standard.
	//
	// example:
	//
	// IDENTITY_PERMISSION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCheckStandardResponseBodyStandards) String() string {
	return dara.Prettify(s)
}

func (s ListCheckStandardResponseBodyStandards) GoString() string {
	return s.String()
}

func (s *ListCheckStandardResponseBodyStandards) GetBindVendor() *int32 {
	return s.BindVendor
}

func (s *ListCheckStandardResponseBodyStandards) GetId() *int64 {
	return s.Id
}

func (s *ListCheckStandardResponseBodyStandards) GetRequirements() []*ListCheckStandardResponseBodyStandardsRequirements {
	return s.Requirements
}

func (s *ListCheckStandardResponseBodyStandards) GetShowName() *string {
	return s.ShowName
}

func (s *ListCheckStandardResponseBodyStandards) GetShowPriorityLevel() *int32 {
	return s.ShowPriorityLevel
}

func (s *ListCheckStandardResponseBodyStandards) GetType() *string {
	return s.Type
}

func (s *ListCheckStandardResponseBodyStandards) SetBindVendor(v int32) *ListCheckStandardResponseBodyStandards {
	s.BindVendor = &v
	return s
}

func (s *ListCheckStandardResponseBodyStandards) SetId(v int64) *ListCheckStandardResponseBodyStandards {
	s.Id = &v
	return s
}

func (s *ListCheckStandardResponseBodyStandards) SetRequirements(v []*ListCheckStandardResponseBodyStandardsRequirements) *ListCheckStandardResponseBodyStandards {
	s.Requirements = v
	return s
}

func (s *ListCheckStandardResponseBodyStandards) SetShowName(v string) *ListCheckStandardResponseBodyStandards {
	s.ShowName = &v
	return s
}

func (s *ListCheckStandardResponseBodyStandards) SetShowPriorityLevel(v int32) *ListCheckStandardResponseBodyStandards {
	s.ShowPriorityLevel = &v
	return s
}

func (s *ListCheckStandardResponseBodyStandards) SetType(v string) *ListCheckStandardResponseBodyStandards {
	s.Type = &v
	return s
}

func (s *ListCheckStandardResponseBodyStandards) Validate() error {
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

type ListCheckStandardResponseBodyStandardsRequirements struct {
	// The ID of the requirement.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of check items in the requirement.
	//
	// example:
	//
	// 10
	RiskCheckCount *int64 `json:"RiskCheckCount,omitempty" xml:"RiskCheckCount,omitempty"`
	// The display name of the search condition.
	//
	// example:
	//
	// RAM identity authentication
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The priority for display.
	//
	// example:
	//
	// 1
	ShowPriorityLevel *int32 `json:"ShowPriorityLevel,omitempty" xml:"ShowPriorityLevel,omitempty"`
}

func (s ListCheckStandardResponseBodyStandardsRequirements) String() string {
	return dara.Prettify(s)
}

func (s ListCheckStandardResponseBodyStandardsRequirements) GoString() string {
	return s.String()
}

func (s *ListCheckStandardResponseBodyStandardsRequirements) GetId() *int64 {
	return s.Id
}

func (s *ListCheckStandardResponseBodyStandardsRequirements) GetRiskCheckCount() *int64 {
	return s.RiskCheckCount
}

func (s *ListCheckStandardResponseBodyStandardsRequirements) GetShowName() *string {
	return s.ShowName
}

func (s *ListCheckStandardResponseBodyStandardsRequirements) GetShowPriorityLevel() *int32 {
	return s.ShowPriorityLevel
}

func (s *ListCheckStandardResponseBodyStandardsRequirements) SetId(v int64) *ListCheckStandardResponseBodyStandardsRequirements {
	s.Id = &v
	return s
}

func (s *ListCheckStandardResponseBodyStandardsRequirements) SetRiskCheckCount(v int64) *ListCheckStandardResponseBodyStandardsRequirements {
	s.RiskCheckCount = &v
	return s
}

func (s *ListCheckStandardResponseBodyStandardsRequirements) SetShowName(v string) *ListCheckStandardResponseBodyStandardsRequirements {
	s.ShowName = &v
	return s
}

func (s *ListCheckStandardResponseBodyStandardsRequirements) SetShowPriorityLevel(v int32) *ListCheckStandardResponseBodyStandardsRequirements {
	s.ShowPriorityLevel = &v
	return s
}

func (s *ListCheckStandardResponseBodyStandardsRequirements) Validate() error {
	return dara.Validate(s)
}
