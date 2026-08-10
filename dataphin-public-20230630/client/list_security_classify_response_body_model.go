// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityClassifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClassifyListResult(v *ListSecurityClassifyResponseBodyClassifyListResult) *ListSecurityClassifyResponseBody
	GetClassifyListResult() *ListSecurityClassifyResponseBodyClassifyListResult
	SetCode(v string) *ListSecurityClassifyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListSecurityClassifyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSecurityClassifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSecurityClassifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSecurityClassifyResponseBody
	GetSuccess() *bool
}

type ListSecurityClassifyResponseBody struct {
	// The result of the data classification list.
	ClassifyListResult *ListSecurityClassifyResponseBodyClassifyListResult `json:"ClassifyListResult,omitempty" xml:"ClassifyListResult,omitempty" type:"Struct"`
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend error.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSecurityClassifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityClassifyResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityClassifyResponseBody) GetClassifyListResult() *ListSecurityClassifyResponseBodyClassifyListResult {
	return s.ClassifyListResult
}

func (s *ListSecurityClassifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSecurityClassifyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSecurityClassifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSecurityClassifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecurityClassifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSecurityClassifyResponseBody) SetClassifyListResult(v *ListSecurityClassifyResponseBodyClassifyListResult) *ListSecurityClassifyResponseBody {
	s.ClassifyListResult = v
	return s
}

func (s *ListSecurityClassifyResponseBody) SetCode(v string) *ListSecurityClassifyResponseBody {
	s.Code = &v
	return s
}

func (s *ListSecurityClassifyResponseBody) SetHttpStatusCode(v int32) *ListSecurityClassifyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSecurityClassifyResponseBody) SetMessage(v string) *ListSecurityClassifyResponseBody {
	s.Message = &v
	return s
}

func (s *ListSecurityClassifyResponseBody) SetRequestId(v string) *ListSecurityClassifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityClassifyResponseBody) SetSuccess(v bool) *ListSecurityClassifyResponseBody {
	s.Success = &v
	return s
}

func (s *ListSecurityClassifyResponseBody) Validate() error {
	if s.ClassifyListResult != nil {
		if err := s.ClassifyListResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecurityClassifyResponseBodyClassifyListResult struct {
	// The list of data classifications.
	ClassifyList []*ListSecurityClassifyResponseBodyClassifyListResultClassifyList `json:"ClassifyList,omitempty" xml:"ClassifyList,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityClassifyResponseBodyClassifyListResult) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityClassifyResponseBodyClassifyListResult) GoString() string {
	return s.String()
}

func (s *ListSecurityClassifyResponseBodyClassifyListResult) GetClassifyList() []*ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	return s.ClassifyList
}

func (s *ListSecurityClassifyResponseBodyClassifyListResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSecurityClassifyResponseBodyClassifyListResult) SetClassifyList(v []*ListSecurityClassifyResponseBodyClassifyListResultClassifyList) *ListSecurityClassifyResponseBodyClassifyListResult {
	s.ClassifyList = v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResult) SetTotalCount(v int32) *ListSecurityClassifyResponseBodyClassifyListResult {
	s.TotalCount = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResult) Validate() error {
	if s.ClassifyList != nil {
		for _, item := range s.ClassifyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSecurityClassifyResponseBodyClassifyListResultClassifyList struct {
	// The catalog path of the classification.
	//
	// example:
	//
	// /Root/Personal Information
	CatalogPath *string `json:"CatalogPath,omitempty" xml:"CatalogPath,omitempty"`
	// The classification description.
	//
	// example:
	//
	// Personal sensitive information classification
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of effective fields.
	//
	// example:
	//
	// 100
	EffectiveFieldCount *int32 `json:"EffectiveFieldCount,omitempty" xml:"EffectiveFieldCount,omitempty"`
	// The classification ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether a masking rule is bound.
	//
	// example:
	//
	// false
	IsBindDesensitizeRule *bool `json:"IsBindDesensitizeRule,omitempty" xml:"IsBindDesensitizeRule,omitempty"`
	// The level ID.
	//
	// example:
	//
	// 1
	LevelId *int64 `json:"LevelId,omitempty" xml:"LevelId,omitempty"`
	// The level name.
	//
	// example:
	//
	// L3
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// The classification name.
	//
	// example:
	//
	// Personal Information
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The short name of the classification.
	//
	// example:
	//
	// PI
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
	// The status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSecurityClassifyResponseBodyClassifyListResultClassifyList) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GoString() string {
	return s.String()
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetCatalogPath() *string {
	return s.CatalogPath
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetDescription() *string {
	return s.Description
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetEffectiveFieldCount() *int32 {
	return s.EffectiveFieldCount
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetId() *int64 {
	return s.Id
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetIsBindDesensitizeRule() *bool {
	return s.IsBindDesensitizeRule
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetLevelId() *int64 {
	return s.LevelId
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetLevelName() *string {
	return s.LevelName
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetName() *string {
	return s.Name
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetShortName() *string {
	return s.ShortName
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) GetStatus() *string {
	return s.Status
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetCatalogPath(v string) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.CatalogPath = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetDescription(v string) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.Description = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetEffectiveFieldCount(v int32) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.EffectiveFieldCount = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetId(v int64) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.Id = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetIsBindDesensitizeRule(v bool) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.IsBindDesensitizeRule = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetLevelId(v int64) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.LevelId = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetLevelName(v string) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.LevelName = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetName(v string) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.Name = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetPriority(v int32) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.Priority = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetShortName(v string) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.ShortName = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) SetStatus(v string) *ListSecurityClassifyResponseBodyClassifyListResultClassifyList {
	s.Status = &v
	return s
}

func (s *ListSecurityClassifyResponseBodyClassifyListResultClassifyList) Validate() error {
	return dara.Validate(s)
}
