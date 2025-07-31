// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListDataQualityTemplatesResponseBodyPageInfo) *ListDataQualityTemplatesResponseBody
	GetPageInfo() *ListDataQualityTemplatesResponseBodyPageInfo
	SetRequestId(v string) *ListDataQualityTemplatesResponseBody
	GetRequestId() *string
}

type ListDataQualityTemplatesResponseBody struct {
	PageInfo *ListDataQualityTemplatesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityTemplatesResponseBody) GetPageInfo() *ListDataQualityTemplatesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListDataQualityTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataQualityTemplatesResponseBody) SetPageInfo(v *ListDataQualityTemplatesResponseBodyPageInfo) *ListDataQualityTemplatesResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListDataQualityTemplatesResponseBody) SetRequestId(v string) *ListDataQualityTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityTemplatesResponseBodyPageInfo struct {
	DataQualityTemplates []*ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates `json:"DataQualityTemplates,omitempty" xml:"DataQualityTemplates,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityTemplatesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityTemplatesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityTemplatesResponseBodyPageInfo) GetDataQualityTemplates() []*ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates {
	return s.DataQualityTemplates
}

func (s *ListDataQualityTemplatesResponseBodyPageInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityTemplatesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityTemplatesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataQualityTemplatesResponseBodyPageInfo) SetDataQualityTemplates(v []*ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) *ListDataQualityTemplatesResponseBodyPageInfo {
	s.DataQualityTemplates = v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfo) SetPageNumber(v int32) *ListDataQualityTemplatesResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfo) SetPageSize(v int32) *ListDataQualityTemplatesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfo) SetTotalCount(v int32) *ListDataQualityTemplatesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates struct {
	// example:
	//
	// 1729816478147
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 7892346529452
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 819cf1f8-29be-4f94-a9d0-c5c06c0c3d2a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1729816478147
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 205250754596036836
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// 205250754596036836
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 7635
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// {
	//
	//     "assertion": "anomaly detection fro id_not_null_cnt",
	//
	//     "id_not_null_cnt": {
	//
	//         "query": "SELECT COUNT(*) AS cnt FROM ${tableName} WHERE dt = \\"$[yyyymmdd-1]\\";"
	//
	//     },
	//
	//     "identity": "819cf1f8-29be-4f94-a9d0-c5c06c0c3d2a"
	//
	// }
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) GoString() string {
	return s.String()
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) GetId() *string {
	return s.Id
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) GetOwner() *string {
	return s.Owner
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) GetSpec() *string {
	return s.Spec
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) SetCreateTime(v int64) *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) SetCreateUser(v string) *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates {
	s.CreateUser = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) SetId(v string) *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates {
	s.Id = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) SetModifyTime(v int64) *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates {
	s.ModifyTime = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) SetModifyUser(v string) *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates {
	s.ModifyUser = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) SetOwner(v string) *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates {
	s.Owner = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) SetProjectId(v int64) *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) SetSpec(v string) *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates {
	s.Spec = &v
	return s
}

func (s *ListDataQualityTemplatesResponseBodyPageInfoDataQualityTemplates) Validate() error {
	return dara.Validate(s)
}
