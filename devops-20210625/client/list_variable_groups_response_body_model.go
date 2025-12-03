// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariableGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListVariableGroupsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListVariableGroupsResponseBody
	GetErrorMessage() *string
	SetNextToken(v string) *ListVariableGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVariableGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListVariableGroupsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListVariableGroupsResponseBody
	GetTotalCount() *int64
	SetVariableGroups(v []*ListVariableGroupsResponseBodyVariableGroups) *ListVariableGroupsResponseBody
	GetVariableGroups() []*ListVariableGroupsResponseBodyVariableGroups
}

type ListVariableGroupsResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// assassa
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 30
	TotalCount     *int64                                          `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	VariableGroups []*ListVariableGroupsResponseBodyVariableGroups `json:"variableGroups,omitempty" xml:"variableGroups,omitempty" type:"Repeated"`
}

func (s ListVariableGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVariableGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListVariableGroupsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListVariableGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVariableGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVariableGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListVariableGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListVariableGroupsResponseBody) GetVariableGroups() []*ListVariableGroupsResponseBodyVariableGroups {
	return s.VariableGroups
}

func (s *ListVariableGroupsResponseBody) SetErrorCode(v string) *ListVariableGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetErrorMessage(v string) *ListVariableGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetNextToken(v string) *ListVariableGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetRequestId(v string) *ListVariableGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetSuccess(v bool) *ListVariableGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetTotalCount(v int64) *ListVariableGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetVariableGroups(v []*ListVariableGroupsResponseBodyVariableGroups) *ListVariableGroupsResponseBody {
	s.VariableGroups = v
	return s
}

func (s *ListVariableGroupsResponseBody) Validate() error {
	if s.VariableGroups != nil {
		for _, item := range s.VariableGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVariableGroupsResponseBodyVariableGroups struct {
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 13232343434343
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// 变量组
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 12234
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 13232343434343
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	// example:
	//
	// 变量组
	Name             *string                                                         `json:"name,omitempty" xml:"name,omitempty"`
	RelatedPipelines []*ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines `json:"relatedPipelines,omitempty" xml:"relatedPipelines,omitempty" type:"Repeated"`
	// example:
	//
	// 1586863220000
	UpdateTime *int64                                                   `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Variables  []*ListVariableGroupsResponseBodyVariableGroupsVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s ListVariableGroupsResponseBodyVariableGroups) String() string {
	return dara.Prettify(s)
}

func (s ListVariableGroupsResponseBodyVariableGroups) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsResponseBodyVariableGroups) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListVariableGroupsResponseBodyVariableGroups) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *ListVariableGroupsResponseBodyVariableGroups) GetDescription() *string {
	return s.Description
}

func (s *ListVariableGroupsResponseBodyVariableGroups) GetId() *int64 {
	return s.Id
}

func (s *ListVariableGroupsResponseBodyVariableGroups) GetModifierAccountId() *string {
	return s.ModifierAccountId
}

func (s *ListVariableGroupsResponseBodyVariableGroups) GetName() *string {
	return s.Name
}

func (s *ListVariableGroupsResponseBodyVariableGroups) GetRelatedPipelines() []*ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines {
	return s.RelatedPipelines
}

func (s *ListVariableGroupsResponseBodyVariableGroups) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListVariableGroupsResponseBodyVariableGroups) GetVariables() []*ListVariableGroupsResponseBodyVariableGroupsVariables {
	return s.Variables
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetCreateTime(v int64) *ListVariableGroupsResponseBodyVariableGroups {
	s.CreateTime = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetCreatorAccountId(v string) *ListVariableGroupsResponseBodyVariableGroups {
	s.CreatorAccountId = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetDescription(v string) *ListVariableGroupsResponseBodyVariableGroups {
	s.Description = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetId(v int64) *ListVariableGroupsResponseBodyVariableGroups {
	s.Id = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetModifierAccountId(v string) *ListVariableGroupsResponseBodyVariableGroups {
	s.ModifierAccountId = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetName(v string) *ListVariableGroupsResponseBodyVariableGroups {
	s.Name = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetRelatedPipelines(v []*ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) *ListVariableGroupsResponseBodyVariableGroups {
	s.RelatedPipelines = v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetUpdateTime(v int64) *ListVariableGroupsResponseBodyVariableGroups {
	s.UpdateTime = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetVariables(v []*ListVariableGroupsResponseBodyVariableGroupsVariables) *ListVariableGroupsResponseBodyVariableGroups {
	s.Variables = v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) Validate() error {
	if s.RelatedPipelines != nil {
		for _, item := range s.RelatedPipelines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines struct {
	// example:
	//
	// 1234
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 流水线
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) String() string {
	return dara.Prettify(s)
}

func (s ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) GetId() *int64 {
	return s.Id
}

func (s *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) GetName() *string {
	return s.Name
}

func (s *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) SetId(v int64) *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines {
	s.Id = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) SetName(v string) *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines {
	s.Name = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) Validate() error {
	return dara.Validate(s)
}

type ListVariableGroupsResponseBodyVariableGroupsVariables struct {
	// example:
	//
	// true
	IsEncrypted *bool `json:"isEncrypted,omitempty" xml:"isEncrypted,omitempty"`
	// example:
	//
	// name1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListVariableGroupsResponseBodyVariableGroupsVariables) String() string {
	return dara.Prettify(s)
}

func (s ListVariableGroupsResponseBodyVariableGroupsVariables) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsResponseBodyVariableGroupsVariables) GetIsEncrypted() *bool {
	return s.IsEncrypted
}

func (s *ListVariableGroupsResponseBodyVariableGroupsVariables) GetName() *string {
	return s.Name
}

func (s *ListVariableGroupsResponseBodyVariableGroupsVariables) GetValue() *string {
	return s.Value
}

func (s *ListVariableGroupsResponseBodyVariableGroupsVariables) SetIsEncrypted(v bool) *ListVariableGroupsResponseBodyVariableGroupsVariables {
	s.IsEncrypted = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroupsVariables) SetName(v string) *ListVariableGroupsResponseBodyVariableGroupsVariables {
	s.Name = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroupsVariables) SetValue(v string) *ListVariableGroupsResponseBodyVariableGroupsVariables {
	s.Value = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroupsVariables) Validate() error {
	return dara.Validate(s)
}
