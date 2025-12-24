// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupProductionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeGroupProductionsResponseBodyData) *DescribeGroupProductionsResponseBody
	GetData() []*DescribeGroupProductionsResponseBodyData
	SetPage(v *DescribeGroupProductionsResponseBodyPage) *DescribeGroupProductionsResponseBody
	GetPage() *DescribeGroupProductionsResponseBodyPage
	SetRequestId(v string) *DescribeGroupProductionsResponseBody
	GetRequestId() *string
}

type DescribeGroupProductionsResponseBody struct {
	// The information about the groups.
	Data []*DescribeGroupProductionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	Page *DescribeGroupProductionsResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 358E012F-****-599D-9ED0-****61CDE615
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupProductionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupProductionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponseBody) GetData() []*DescribeGroupProductionsResponseBodyData {
	return s.Data
}

func (s *DescribeGroupProductionsResponseBody) GetPage() *DescribeGroupProductionsResponseBodyPage {
	return s.Page
}

func (s *DescribeGroupProductionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupProductionsResponseBody) SetData(v []*DescribeGroupProductionsResponseBodyData) *DescribeGroupProductionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeGroupProductionsResponseBody) SetPage(v *DescribeGroupProductionsResponseBodyPage) *DescribeGroupProductionsResponseBody {
	s.Page = v
	return s
}

func (s *DescribeGroupProductionsResponseBody) SetRequestId(v string) *DescribeGroupProductionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupProductionsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupProductionsResponseBodyData struct {
	// The name of the cloud service.
	//
	// example:
	//
	// Cloud Communication
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The information about the cloud services.
	Productions []*DescribeGroupProductionsResponseBodyDataProductions `json:"Productions,omitempty" xml:"Productions,omitempty" type:"Repeated"`
}

func (s DescribeGroupProductionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupProductionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeGroupProductionsResponseBodyData) GetProductions() []*DescribeGroupProductionsResponseBodyDataProductions {
	return s.Productions
}

func (s *DescribeGroupProductionsResponseBodyData) SetGroupName(v string) *DescribeGroupProductionsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyData) SetProductions(v []*DescribeGroupProductionsResponseBodyDataProductions) *DescribeGroupProductionsResponseBodyData {
	s.Productions = v
	return s
}

func (s *DescribeGroupProductionsResponseBodyData) Validate() error {
	if s.Productions != nil {
		for _, item := range s.Productions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGroupProductionsResponseBodyDataProductions struct {
	// The code of the cloud service.
	//
	// example:
	//
	// DM
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default requested domain name.
	//
	// example:
	//
	// dm.aliyuncs.com
	DefaultDomain *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// The default version.
	//
	// example:
	//
	// 2014-08-15
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the cloud service.
	//
	// example:
	//
	// DM
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The requested domain names.
	FullDomains []*string `json:"FullDomains,omitempty" xml:"FullDomains,omitempty" type:"Repeated"`
	// The name of the group.
	//
	// example:
	//
	// Cloud Communication
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The name of the cloud service.
	//
	// example:
	//
	// DirectMail
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The RAM policies of the cloud service.
	PolicyList []*DescribeGroupProductionsResponseBodyDataProductionsPolicyList `json:"PolicyList,omitempty" xml:"PolicyList,omitempty" type:"Repeated"`
	// The Resource Access Management (RAM) code of the POP to which the resource belongs.
	//
	// example:
	//
	// dm
	RamCode *string `json:"RamCode,omitempty" xml:"RamCode,omitempty"`
	// The short name of the cloud service.
	//
	// example:
	//
	// DM
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
	// The information source of the cloud service.
	//
	// example:
	//
	// next
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The API versions.
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeGroupProductionsResponseBodyDataProductions) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupProductionsResponseBodyDataProductions) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetCode() *string {
	return s.Code
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetDefaultDomain() *string {
	return s.DefaultDomain
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetDescription() *string {
	return s.Description
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetFullDomains() []*string {
	return s.FullDomains
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetGroup() *string {
	return s.Group
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetName() *string {
	return s.Name
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetPolicyList() []*DescribeGroupProductionsResponseBodyDataProductionsPolicyList {
	return s.PolicyList
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetRamCode() *string {
	return s.RamCode
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetShortName() *string {
	return s.ShortName
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetSource() *string {
	return s.Source
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) GetVersions() []*string {
	return s.Versions
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetCode(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Code = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetDefaultDomain(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.DefaultDomain = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetDefaultVersion(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.DefaultVersion = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetDescription(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Description = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetFullDomains(v []*string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.FullDomains = v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetGroup(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Group = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetName(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Name = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetPolicyList(v []*DescribeGroupProductionsResponseBodyDataProductionsPolicyList) *DescribeGroupProductionsResponseBodyDataProductions {
	s.PolicyList = v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetRamCode(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.RamCode = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetShortName(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.ShortName = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetSource(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Source = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetVersions(v []*string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Versions = v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) Validate() error {
	if s.PolicyList != nil {
		for _, item := range s.PolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGroupProductionsResponseBodyDataProductionsPolicyList struct {
	// The name of the RAM policy.
	//
	// example:
	//
	// AliyunRAMReadOnlyAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the RAM policy. Valid values:
	//
	// 	- **All**: permissions.
	//
	// 	- **Read-only**: permissions.
	//
	// example:
	//
	// All
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeGroupProductionsResponseBodyDataProductionsPolicyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupProductionsResponseBodyDataProductionsPolicyList) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponseBodyDataProductionsPolicyList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeGroupProductionsResponseBodyDataProductionsPolicyList) GetType() *string {
	return s.Type
}

func (s *DescribeGroupProductionsResponseBodyDataProductionsPolicyList) SetPolicyName(v string) *DescribeGroupProductionsResponseBodyDataProductionsPolicyList {
	s.PolicyName = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductionsPolicyList) SetType(v string) *DescribeGroupProductionsResponseBodyDataProductionsPolicyList {
	s.Type = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductionsPolicyList) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupProductionsResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGroupProductionsResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupProductionsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGroupProductionsResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupProductionsResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGroupProductionsResponseBodyPage) SetPageNumber(v int32) *DescribeGroupProductionsResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyPage) SetPageSize(v int32) *DescribeGroupProductionsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyPage) SetTotalCount(v int32) *DescribeGroupProductionsResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyPage) Validate() error {
	return dara.Validate(s)
}
