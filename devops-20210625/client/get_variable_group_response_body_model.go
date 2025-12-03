// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVariableGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetVariableGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetVariableGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetVariableGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVariableGroupResponseBody
	GetSuccess() *bool
	SetVariableGroup(v *GetVariableGroupResponseBodyVariableGroup) *GetVariableGroupResponseBody
	GetVariableGroup() *GetVariableGroupResponseBodyVariableGroup
}

type GetVariableGroupResponseBody struct {
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
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success       *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
	VariableGroup *GetVariableGroupResponseBodyVariableGroup `json:"variableGroup,omitempty" xml:"variableGroup,omitempty" type:"Struct"`
}

func (s GetVariableGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVariableGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetVariableGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetVariableGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVariableGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVariableGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVariableGroupResponseBody) GetVariableGroup() *GetVariableGroupResponseBodyVariableGroup {
	return s.VariableGroup
}

func (s *GetVariableGroupResponseBody) SetErrorCode(v string) *GetVariableGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetVariableGroupResponseBody) SetErrorMessage(v string) *GetVariableGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetVariableGroupResponseBody) SetRequestId(v string) *GetVariableGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVariableGroupResponseBody) SetSuccess(v bool) *GetVariableGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetVariableGroupResponseBody) SetVariableGroup(v *GetVariableGroupResponseBodyVariableGroup) *GetVariableGroupResponseBody {
	s.VariableGroup = v
	return s
}

func (s *GetVariableGroupResponseBody) Validate() error {
	if s.VariableGroup != nil {
		if err := s.VariableGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVariableGroupResponseBodyVariableGroup struct {
	// example:
	//
	// 13232343434343
	CcreatorAccountId *string `json:"ccreatorAccountId,omitempty" xml:"ccreatorAccountId,omitempty"`
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
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
	Name             *string                                                      `json:"name,omitempty" xml:"name,omitempty"`
	RelatedPipelines []*GetVariableGroupResponseBodyVariableGroupRelatedPipelines `json:"relatedPipelines,omitempty" xml:"relatedPipelines,omitempty" type:"Repeated"`
	// example:
	//
	// 1586863220000
	UpdateTime *int64                                                `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Variables  []*GetVariableGroupResponseBodyVariableGroupVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s GetVariableGroupResponseBodyVariableGroup) String() string {
	return dara.Prettify(s)
}

func (s GetVariableGroupResponseBodyVariableGroup) GoString() string {
	return s.String()
}

func (s *GetVariableGroupResponseBodyVariableGroup) GetCcreatorAccountId() *string {
	return s.CcreatorAccountId
}

func (s *GetVariableGroupResponseBodyVariableGroup) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetVariableGroupResponseBodyVariableGroup) GetDescription() *string {
	return s.Description
}

func (s *GetVariableGroupResponseBodyVariableGroup) GetId() *int64 {
	return s.Id
}

func (s *GetVariableGroupResponseBodyVariableGroup) GetModifierAccountId() *string {
	return s.ModifierAccountId
}

func (s *GetVariableGroupResponseBodyVariableGroup) GetName() *string {
	return s.Name
}

func (s *GetVariableGroupResponseBodyVariableGroup) GetRelatedPipelines() []*GetVariableGroupResponseBodyVariableGroupRelatedPipelines {
	return s.RelatedPipelines
}

func (s *GetVariableGroupResponseBodyVariableGroup) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetVariableGroupResponseBodyVariableGroup) GetVariables() []*GetVariableGroupResponseBodyVariableGroupVariables {
	return s.Variables
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetCcreatorAccountId(v string) *GetVariableGroupResponseBodyVariableGroup {
	s.CcreatorAccountId = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetCreateTime(v int64) *GetVariableGroupResponseBodyVariableGroup {
	s.CreateTime = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetDescription(v string) *GetVariableGroupResponseBodyVariableGroup {
	s.Description = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetId(v int64) *GetVariableGroupResponseBodyVariableGroup {
	s.Id = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetModifierAccountId(v string) *GetVariableGroupResponseBodyVariableGroup {
	s.ModifierAccountId = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetName(v string) *GetVariableGroupResponseBodyVariableGroup {
	s.Name = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetRelatedPipelines(v []*GetVariableGroupResponseBodyVariableGroupRelatedPipelines) *GetVariableGroupResponseBodyVariableGroup {
	s.RelatedPipelines = v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetUpdateTime(v int64) *GetVariableGroupResponseBodyVariableGroup {
	s.UpdateTime = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetVariables(v []*GetVariableGroupResponseBodyVariableGroupVariables) *GetVariableGroupResponseBodyVariableGroup {
	s.Variables = v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) Validate() error {
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

type GetVariableGroupResponseBodyVariableGroupRelatedPipelines struct {
	// example:
	//
	// 1234
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 流水线
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetVariableGroupResponseBodyVariableGroupRelatedPipelines) String() string {
	return dara.Prettify(s)
}

func (s GetVariableGroupResponseBodyVariableGroupRelatedPipelines) GoString() string {
	return s.String()
}

func (s *GetVariableGroupResponseBodyVariableGroupRelatedPipelines) GetId() *int64 {
	return s.Id
}

func (s *GetVariableGroupResponseBodyVariableGroupRelatedPipelines) GetName() *string {
	return s.Name
}

func (s *GetVariableGroupResponseBodyVariableGroupRelatedPipelines) SetId(v int64) *GetVariableGroupResponseBodyVariableGroupRelatedPipelines {
	s.Id = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroupRelatedPipelines) SetName(v string) *GetVariableGroupResponseBodyVariableGroupRelatedPipelines {
	s.Name = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroupRelatedPipelines) Validate() error {
	return dara.Validate(s)
}

type GetVariableGroupResponseBodyVariableGroupVariables struct {
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

func (s GetVariableGroupResponseBodyVariableGroupVariables) String() string {
	return dara.Prettify(s)
}

func (s GetVariableGroupResponseBodyVariableGroupVariables) GoString() string {
	return s.String()
}

func (s *GetVariableGroupResponseBodyVariableGroupVariables) GetIsEncrypted() *bool {
	return s.IsEncrypted
}

func (s *GetVariableGroupResponseBodyVariableGroupVariables) GetName() *string {
	return s.Name
}

func (s *GetVariableGroupResponseBodyVariableGroupVariables) GetValue() *string {
	return s.Value
}

func (s *GetVariableGroupResponseBodyVariableGroupVariables) SetIsEncrypted(v bool) *GetVariableGroupResponseBodyVariableGroupVariables {
	s.IsEncrypted = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroupVariables) SetName(v string) *GetVariableGroupResponseBodyVariableGroupVariables {
	s.Name = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroupVariables) SetValue(v string) *GetVariableGroupResponseBodyVariableGroupVariables {
	s.Value = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroupVariables) Validate() error {
	return dara.Validate(s)
}
