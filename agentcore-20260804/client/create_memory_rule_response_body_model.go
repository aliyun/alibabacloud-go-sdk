// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMemoryRuleResponseBody
	GetCode() *string
	SetData(v *CreateMemoryRuleResponseBodyData) *CreateMemoryRuleResponseBody
	GetData() *CreateMemoryRuleResponseBodyData
	SetHttpStatusCode(v int32) *CreateMemoryRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateMemoryRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMemoryRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMemoryRuleResponseBody
	GetSuccess() *bool
}

type CreateMemoryRuleResponseBody struct {
	// code
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *CreateMemoryRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateMemoryRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMemoryRuleResponseBody) GetData() *CreateMemoryRuleResponseBodyData {
	return s.Data
}

func (s *CreateMemoryRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateMemoryRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMemoryRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemoryRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMemoryRuleResponseBody) SetCode(v string) *CreateMemoryRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMemoryRuleResponseBody) SetData(v *CreateMemoryRuleResponseBodyData) *CreateMemoryRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateMemoryRuleResponseBody) SetHttpStatusCode(v int32) *CreateMemoryRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMemoryRuleResponseBody) SetMessage(v string) *CreateMemoryRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMemoryRuleResponseBody) SetRequestId(v string) *CreateMemoryRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemoryRuleResponseBody) SetSuccess(v bool) *CreateMemoryRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMemoryRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryRuleResponseBodyData struct {
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
	// Extract long-term preferences explicitly expressed by users
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
	Strategy *CreateMemoryRuleResponseBodyDataStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
	// updatedAt
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CreateMemoryRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMemoryRuleResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateMemoryRuleResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateMemoryRuleResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMemoryRuleResponseBodyData) GetMemoryStoreId() *string {
	return s.MemoryStoreId
}

func (s *CreateMemoryRuleResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateMemoryRuleResponseBodyData) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateMemoryRuleResponseBodyData) GetStrategy() *CreateMemoryRuleResponseBodyDataStrategy {
	return s.Strategy
}

func (s *CreateMemoryRuleResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateMemoryRuleResponseBodyData) SetCreatedAt(v string) *CreateMemoryRuleResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateMemoryRuleResponseBodyData) SetDescription(v string) *CreateMemoryRuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateMemoryRuleResponseBodyData) SetEnabled(v bool) *CreateMemoryRuleResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *CreateMemoryRuleResponseBodyData) SetMemoryStoreId(v string) *CreateMemoryRuleResponseBodyData {
	s.MemoryStoreId = &v
	return s
}

func (s *CreateMemoryRuleResponseBodyData) SetName(v string) *CreateMemoryRuleResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateMemoryRuleResponseBodyData) SetRuleId(v string) *CreateMemoryRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *CreateMemoryRuleResponseBodyData) SetStrategy(v *CreateMemoryRuleResponseBodyDataStrategy) *CreateMemoryRuleResponseBodyData {
	s.Strategy = v
	return s
}

func (s *CreateMemoryRuleResponseBodyData) SetUpdatedAt(v string) *CreateMemoryRuleResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateMemoryRuleResponseBodyData) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryRuleResponseBodyDataStrategy struct {
	// The strategy-level instruction. It takes effect together with the store-level instruction. If a conflict occurs, this field takes precedence.
	//
	// example:
	//
	// Retain only long-term preferences explicitly expressed by users
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

func (s CreateMemoryRuleResponseBodyDataStrategy) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRuleResponseBodyDataStrategy) GoString() string {
	return s.String()
}

func (s *CreateMemoryRuleResponseBodyDataStrategy) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateMemoryRuleResponseBodyDataStrategy) GetType() *string {
	return s.Type
}

func (s *CreateMemoryRuleResponseBodyDataStrategy) SetInstruction(v string) *CreateMemoryRuleResponseBodyDataStrategy {
	s.Instruction = &v
	return s
}

func (s *CreateMemoryRuleResponseBodyDataStrategy) SetType(v string) *CreateMemoryRuleResponseBodyDataStrategy {
	s.Type = &v
	return s
}

func (s *CreateMemoryRuleResponseBodyDataStrategy) Validate() error {
	return dara.Validate(s)
}
