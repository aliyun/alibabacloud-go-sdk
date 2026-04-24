// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribePolarFsQuotaResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribePolarFsQuotaResponseBody
	GetPageRecordCount() *string
	SetPageSize(v string) *DescribePolarFsQuotaResponseBody
	GetPageSize() *string
	SetPath(v string) *DescribePolarFsQuotaResponseBody
	GetPath() *string
	SetPolarFsInstanceId(v string) *DescribePolarFsQuotaResponseBody
	GetPolarFsInstanceId() *string
	SetPolicyItems(v []*DescribePolarFsQuotaResponseBodyPolicyItems) *DescribePolarFsQuotaResponseBody
	GetPolicyItems() []*DescribePolarFsQuotaResponseBodyPolicyItems
	SetQuotaItems(v []*DescribePolarFsQuotaResponseBodyQuotaItems) *DescribePolarFsQuotaResponseBody
	GetQuotaItems() []*DescribePolarFsQuotaResponseBodyQuotaItems
	SetRequestId(v string) *DescribePolarFsQuotaResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribePolarFsQuotaResponseBody
	GetTotalRecordCount() *string
}

type DescribePolarFsQuotaResponseBody struct {
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// /data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string                                        `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	PolicyItems       []*DescribePolarFsQuotaResponseBodyPolicyItems `json:"PolicyItems,omitempty" xml:"PolicyItems,omitempty" type:"Repeated"`
	QuotaItems        []*DescribePolarFsQuotaResponseBodyQuotaItems  `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribePolarFsQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarFsQuotaResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribePolarFsQuotaResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribePolarFsQuotaResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribePolarFsQuotaResponseBody) GetPath() *string {
	return s.Path
}

func (s *DescribePolarFsQuotaResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsQuotaResponseBody) GetPolicyItems() []*DescribePolarFsQuotaResponseBodyPolicyItems {
	return s.PolicyItems
}

func (s *DescribePolarFsQuotaResponseBody) GetQuotaItems() []*DescribePolarFsQuotaResponseBodyQuotaItems {
	return s.QuotaItems
}

func (s *DescribePolarFsQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarFsQuotaResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribePolarFsQuotaResponseBody) SetPageNumber(v string) *DescribePolarFsQuotaResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBody) SetPageRecordCount(v string) *DescribePolarFsQuotaResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBody) SetPageSize(v string) *DescribePolarFsQuotaResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBody) SetPath(v string) *DescribePolarFsQuotaResponseBody {
	s.Path = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBody) SetPolarFsInstanceId(v string) *DescribePolarFsQuotaResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBody) SetPolicyItems(v []*DescribePolarFsQuotaResponseBodyPolicyItems) *DescribePolarFsQuotaResponseBody {
	s.PolicyItems = v
	return s
}

func (s *DescribePolarFsQuotaResponseBody) SetQuotaItems(v []*DescribePolarFsQuotaResponseBodyQuotaItems) *DescribePolarFsQuotaResponseBody {
	s.QuotaItems = v
	return s
}

func (s *DescribePolarFsQuotaResponseBody) SetRequestId(v string) *DescribePolarFsQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBody) SetTotalRecordCount(v string) *DescribePolarFsQuotaResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBody) Validate() error {
	if s.PolicyItems != nil {
		for _, item := range s.PolicyItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QuotaItems != nil {
		for _, item := range s.QuotaItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarFsQuotaResponseBodyPolicyItems struct {
	// example:
	//
	// 7200
	AccessTTL *int64 `json:"AccessTTL,omitempty" xml:"AccessTTL,omitempty"`
	// example:
	//
	// 7200
	ChangeTTL *int64 `json:"ChangeTTL,omitempty" xml:"ChangeTTL,omitempty"`
	// example:
	//
	// NULL
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// NULL
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// example:
	//
	// 77
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// example:
	//
	// 73
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// /a*
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// example:
	//
	// xxxxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s DescribePolarFsQuotaResponseBodyPolicyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsQuotaResponseBodyPolicyItems) GoString() string {
	return s.String()
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetAccessTTL() *int64 {
	return s.AccessTTL
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetChangeTTL() *int64 {
	return s.ChangeTTL
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetExclude() *string {
	return s.Exclude
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetInclude() *string {
	return s.Include
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetName() *string {
	return s.Name
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetAccessTTL(v int64) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.AccessTTL = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetChangeTTL(v int64) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.ChangeTTL = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetDescription(v string) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.Description = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetEnabled(v bool) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.Enabled = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetExclude(v string) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.Exclude = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetFileCountLimit(v int64) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.FileCountLimit = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetId(v int64) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.Id = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetInclude(v string) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.Include = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetName(v string) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.Name = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetPriority(v int32) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.Priority = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) SetSizeLimit(v int64) *DescribePolarFsQuotaResponseBodyPolicyItems {
	s.SizeLimit = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyPolicyItems) Validate() error {
	return dara.Validate(s)
}

type DescribePolarFsQuotaResponseBodyQuotaItems struct {
	// example:
	//
	// 1073741824
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// 100
	Inodes *int64 `json:"Inodes,omitempty" xml:"Inodes,omitempty"`
	// example:
	//
	// /data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 104857600
	UsedCapacity *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
	// example:
	//
	// 1
	UsedInodes *int64 `json:"UsedInodes,omitempty" xml:"UsedInodes,omitempty"`
}

func (s DescribePolarFsQuotaResponseBodyQuotaItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsQuotaResponseBodyQuotaItems) GoString() string {
	return s.String()
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) GetInodes() *int64 {
	return s.Inodes
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) GetPath() *string {
	return s.Path
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) GetUsedCapacity() *int64 {
	return s.UsedCapacity
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) GetUsedInodes() *int64 {
	return s.UsedInodes
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) SetCapacity(v int64) *DescribePolarFsQuotaResponseBodyQuotaItems {
	s.Capacity = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) SetInodes(v int64) *DescribePolarFsQuotaResponseBodyQuotaItems {
	s.Inodes = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) SetPath(v string) *DescribePolarFsQuotaResponseBodyQuotaItems {
	s.Path = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) SetUsedCapacity(v int64) *DescribePolarFsQuotaResponseBodyQuotaItems {
	s.UsedCapacity = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) SetUsedInodes(v int64) *DescribePolarFsQuotaResponseBodyQuotaItems {
	s.UsedInodes = &v
	return s
}

func (s *DescribePolarFsQuotaResponseBodyQuotaItems) Validate() error {
	return dara.Validate(s)
}
