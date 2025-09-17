// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkgroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeWorkgroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeWorkgroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWorkgroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWorkgroupsResponseBody
	GetTotalCount() *int32
	SetWorkgroups(v *DescribeWorkgroupsResponseBodyWorkgroups) *DescribeWorkgroupsResponseBody
	GetWorkgroups() *DescribeWorkgroupsResponseBodyWorkgroups
}

type DescribeWorkgroupsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of workgroups.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details of the workgroup.
	Workgroups *DescribeWorkgroupsResponseBodyWorkgroups `json:"Workgroups,omitempty" xml:"Workgroups,omitempty" type:"Struct"`
}

func (s DescribeWorkgroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeWorkgroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWorkgroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWorkgroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWorkgroupsResponseBody) GetWorkgroups() *DescribeWorkgroupsResponseBodyWorkgroups {
	return s.Workgroups
}

func (s *DescribeWorkgroupsResponseBody) SetPageNumber(v int32) *DescribeWorkgroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeWorkgroupsResponseBody) SetPageSize(v int32) *DescribeWorkgroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWorkgroupsResponseBody) SetRequestId(v string) *DescribeWorkgroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWorkgroupsResponseBody) SetTotalCount(v int32) *DescribeWorkgroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWorkgroupsResponseBody) SetWorkgroups(v *DescribeWorkgroupsResponseBodyWorkgroups) *DescribeWorkgroupsResponseBody {
	s.Workgroups = v
	return s
}

func (s *DescribeWorkgroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkgroupsResponseBodyWorkgroups struct {
	Workgroup []*DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup `json:"Workgroup,omitempty" xml:"Workgroup,omitempty" type:"Repeated"`
}

func (s DescribeWorkgroupsResponseBodyWorkgroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsResponseBodyWorkgroups) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsResponseBodyWorkgroups) GetWorkgroup() []*DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup {
	return s.Workgroup
}

func (s *DescribeWorkgroupsResponseBodyWorkgroups) SetWorkgroup(v []*DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) *DescribeWorkgroupsResponseBodyWorkgroups {
	s.Workgroup = v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroups) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup struct {
	// The description of the workgroup.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the workgroup.
	//
	// example:
	//
	// testWorkgroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The state of the workgroup. Valid values:
	//
	// 	- NotStarted
	//
	// 	- InProgress
	//
	// 	- Cutover
	//
	// 	- Completed
	//
	// example:
	//
	// InProgress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information of the workgroup.
	Tags *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The alert information of the workgroup, which can contain multiple types of alerts.
	Warnings *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarnings `json:"Warnings,omitempty" xml:"Warnings,omitempty" type:"Struct"`
	// The ID of the workgroup.
	//
	// example:
	//
	// w-***
	WorkgroupId *string `json:"WorkgroupId,omitempty" xml:"WorkgroupId,omitempty"`
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) GetDescription() *string {
	return s.Description
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) GetName() *string {
	return s.Name
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) GetStatus() *string {
	return s.Status
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) GetTags() *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTags {
	return s.Tags
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) GetWarnings() *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarnings {
	return s.Warnings
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) GetWorkgroupId() *string {
	return s.WorkgroupId
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) SetDescription(v string) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup {
	s.Description = &v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) SetName(v string) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup {
	s.Name = &v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) SetStatus(v string) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup {
	s.Status = &v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) SetTags(v *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTags) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup {
	s.Tags = v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) SetWarnings(v *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarnings) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup {
	s.Warnings = v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) SetWorkgroupId(v string) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup {
	s.WorkgroupId = &v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroup) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTags struct {
	Tag []*DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTags) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTags) GetTag() []*DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag {
	return s.Tag
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTags) SetTag(v []*DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTags {
	s.Tag = v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTags) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag struct {
	// The tag key of the workgroup.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the workgroup.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag) SetKey(v string) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag) SetValue(v string) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarnings struct {
	Warning []*DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning `json:"Warning,omitempty" xml:"Warning,omitempty" type:"Repeated"`
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarnings) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarnings) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarnings) GetWarning() []*DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning {
	return s.Warning
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarnings) SetWarning(v []*DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarnings {
	s.Warning = v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarnings) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning struct {
	// The migration sources for which alerts are generated.
	SourceIds *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarningSourceIds `json:"SourceIds,omitempty" xml:"SourceIds,omitempty" type:"Struct"`
	// The type of the alert. Valid values:
	//
	// 	- InError: A migration job failed.
	//
	// 	- UnRelated: No migration job is created for a migration source.
	//
	// 	- NotPassed: A migration job failed to pass the migration test.
	//
	// example:
	//
	// InError
	WarningType *string `json:"WarningType,omitempty" xml:"WarningType,omitempty"`
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning) GetSourceIds() *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarningSourceIds {
	return s.SourceIds
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning) GetWarningType() *string {
	return s.WarningType
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning) SetSourceIds(v *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarningSourceIds) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning {
	s.SourceIds = v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning) SetWarningType(v string) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning {
	s.WarningType = &v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarning) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarningSourceIds struct {
	SourceId []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarningSourceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarningSourceIds) GoString() string {
	return s.String()
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarningSourceIds) GetSourceId() []*string {
	return s.SourceId
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarningSourceIds) SetSourceId(v []*string) *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarningSourceIds {
	s.SourceId = v
	return s
}

func (s *DescribeWorkgroupsResponseBodyWorkgroupsWorkgroupWarningsWarningSourceIds) Validate() error {
	return dara.Validate(s)
}
