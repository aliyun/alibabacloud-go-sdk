// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMemoryRuleResponseBody
	GetCode() *string
	SetData(v *UpdateMemoryRuleResponseBodyData) *UpdateMemoryRuleResponseBody
	GetData() *UpdateMemoryRuleResponseBodyData
	SetHttpStatusCode(v int32) *UpdateMemoryRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMemoryRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMemoryRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMemoryRuleResponseBody
	GetSuccess() *bool
}

type UpdateMemoryRuleResponseBody struct {
	// code
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *UpdateMemoryRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// message
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// requestId
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateMemoryRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMemoryRuleResponseBody) GetData() *UpdateMemoryRuleResponseBodyData {
	return s.Data
}

func (s *UpdateMemoryRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMemoryRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMemoryRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMemoryRuleResponseBody) SetCode(v string) *UpdateMemoryRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMemoryRuleResponseBody) SetData(v *UpdateMemoryRuleResponseBodyData) *UpdateMemoryRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMemoryRuleResponseBody) SetHttpStatusCode(v int32) *UpdateMemoryRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMemoryRuleResponseBody) SetMessage(v string) *UpdateMemoryRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMemoryRuleResponseBody) SetRequestId(v string) *UpdateMemoryRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryRuleResponseBody) SetSuccess(v bool) *UpdateMemoryRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMemoryRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryRuleResponseBodyData struct {
	// createdAt
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// description
	//
	// example:
	//
	// Extract long-term preferences explicitly expressed by the user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// enabled
	//
	// This parameter is required.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// memoryStoreId
	//
	// This parameter is required.
	//
	// example:
	//
	// ms-2f5f9c5a8d0e4c3b
	MemoryStoreId *string `json:"memoryStoreId,omitempty" xml:"memoryStoreId,omitempty"`
	// name
	//
	// This parameter is required.
	//
	// example:
	//
	// Long-term preferences
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// ruleId
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-2f5f9c5a8d0e4c3b
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// strategy
	//
	// This parameter is required.
	Strategy *UpdateMemoryRuleResponseBodyDataStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
	// updatedAt
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s UpdateMemoryRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRuleResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateMemoryRuleResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateMemoryRuleResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMemoryRuleResponseBodyData) GetMemoryStoreId() *string {
	return s.MemoryStoreId
}

func (s *UpdateMemoryRuleResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateMemoryRuleResponseBodyData) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateMemoryRuleResponseBodyData) GetStrategy() *UpdateMemoryRuleResponseBodyDataStrategy {
	return s.Strategy
}

func (s *UpdateMemoryRuleResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateMemoryRuleResponseBodyData) SetCreatedAt(v string) *UpdateMemoryRuleResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateMemoryRuleResponseBodyData) SetDescription(v string) *UpdateMemoryRuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateMemoryRuleResponseBodyData) SetEnabled(v bool) *UpdateMemoryRuleResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *UpdateMemoryRuleResponseBodyData) SetMemoryStoreId(v string) *UpdateMemoryRuleResponseBodyData {
	s.MemoryStoreId = &v
	return s
}

func (s *UpdateMemoryRuleResponseBodyData) SetName(v string) *UpdateMemoryRuleResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateMemoryRuleResponseBodyData) SetRuleId(v string) *UpdateMemoryRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *UpdateMemoryRuleResponseBodyData) SetStrategy(v *UpdateMemoryRuleResponseBodyDataStrategy) *UpdateMemoryRuleResponseBodyData {
	s.Strategy = v
	return s
}

func (s *UpdateMemoryRuleResponseBodyData) SetUpdatedAt(v string) *UpdateMemoryRuleResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateMemoryRuleResponseBodyData) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryRuleResponseBodyDataStrategy struct {
	// The strategy-level instruction. It takes effect together with the store-level instruction. In case of a conflict, this field takes precedence.
	//
	// example:
	//
	// Retain only long-term preferences explicitly expressed by the user
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// type
	//
	// This parameter is required.
	//
	// example:
	//
	// auto
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateMemoryRuleResponseBodyDataStrategy) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRuleResponseBodyDataStrategy) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRuleResponseBodyDataStrategy) GetInstruction() *string {
	return s.Instruction
}

func (s *UpdateMemoryRuleResponseBodyDataStrategy) GetType() *string {
	return s.Type
}

func (s *UpdateMemoryRuleResponseBodyDataStrategy) SetInstruction(v string) *UpdateMemoryRuleResponseBodyDataStrategy {
	s.Instruction = &v
	return s
}

func (s *UpdateMemoryRuleResponseBodyDataStrategy) SetType(v string) *UpdateMemoryRuleResponseBodyDataStrategy {
	s.Type = &v
	return s
}

func (s *UpdateMemoryRuleResponseBodyDataStrategy) Validate() error {
	return dara.Validate(s)
}
