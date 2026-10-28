// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListInstanceCatalogRequest
	GetLang() *string
	SetOnlyCustom(v bool) *ListInstanceCatalogRequest
	GetOnlyCustom() *bool
	SetRegionId(v string) *ListInstanceCatalogRequest
	GetRegionId() *string
	SetRequirementIds(v []*int64) *ListInstanceCatalogRequest
	GetRequirementIds() []*int64
	SetStandardIds(v []*int64) *ListInstanceCatalogRequest
	GetStandardIds() []*int64
	SetTaskSources(v []*string) *ListInstanceCatalogRequest
	GetTaskSources() []*string
	SetTypes(v []*string) *ListInstanceCatalogRequest
	GetTypes() []*string
}

type ListInstanceCatalogRequest struct {
	// The language type of the request and response messages. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to filter assets that support custom checks. Valid values:
	//
	// - **true**: Filters assets that support custom checks.
	//
	// - **false**: All check assets. This is the default value.
	//
	// example:
	//
	// true
	OnlyCustom *bool `json:"OnlyCustom,omitempty" xml:"OnlyCustom,omitempty"`
	// The ID of the region where the instance to be queried is located. Valid values:
	//
	// - **cn-hangzhou**: Center.
	//
	// - **ap-southeast-1**: Singapore.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of requirement IDs.
	RequirementIds []*int64 `json:"RequirementIds,omitempty" xml:"RequirementIds,omitempty" type:"Repeated"`
	// The list of standard IDs.
	StandardIds []*int64 `json:"StandardIds,omitempty" xml:"StandardIds,omitempty" type:"Repeated"`
	// The list of task sources.
	TaskSources []*string `json:"TaskSources,omitempty" xml:"TaskSources,omitempty" type:"Repeated"`
	// The list of types to which the check standards belong.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListInstanceCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceCatalogRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceCatalogRequest) GetLang() *string {
	return s.Lang
}

func (s *ListInstanceCatalogRequest) GetOnlyCustom() *bool {
	return s.OnlyCustom
}

func (s *ListInstanceCatalogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceCatalogRequest) GetRequirementIds() []*int64 {
	return s.RequirementIds
}

func (s *ListInstanceCatalogRequest) GetStandardIds() []*int64 {
	return s.StandardIds
}

func (s *ListInstanceCatalogRequest) GetTaskSources() []*string {
	return s.TaskSources
}

func (s *ListInstanceCatalogRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListInstanceCatalogRequest) SetLang(v string) *ListInstanceCatalogRequest {
	s.Lang = &v
	return s
}

func (s *ListInstanceCatalogRequest) SetOnlyCustom(v bool) *ListInstanceCatalogRequest {
	s.OnlyCustom = &v
	return s
}

func (s *ListInstanceCatalogRequest) SetRegionId(v string) *ListInstanceCatalogRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceCatalogRequest) SetRequirementIds(v []*int64) *ListInstanceCatalogRequest {
	s.RequirementIds = v
	return s
}

func (s *ListInstanceCatalogRequest) SetStandardIds(v []*int64) *ListInstanceCatalogRequest {
	s.StandardIds = v
	return s
}

func (s *ListInstanceCatalogRequest) SetTaskSources(v []*string) *ListInstanceCatalogRequest {
	s.TaskSources = v
	return s
}

func (s *ListInstanceCatalogRequest) SetTypes(v []*string) *ListInstanceCatalogRequest {
	s.Types = v
	return s
}

func (s *ListInstanceCatalogRequest) Validate() error {
	return dara.Validate(s)
}
