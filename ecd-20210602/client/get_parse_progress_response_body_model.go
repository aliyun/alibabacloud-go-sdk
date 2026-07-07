// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParseProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetParseProgressResponseBodyData) *GetParseProgressResponseBody
	GetData() *GetParseProgressResponseBodyData
	SetRequestId(v string) *GetParseProgressResponseBody
	GetRequestId() *string
}

type GetParseProgressResponseBody struct {
	// The response data object for parsing the skill package.
	Data *GetParseProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetParseProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetParseProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetParseProgressResponseBody) GetData() *GetParseProgressResponseBodyData {
	return s.Data
}

func (s *GetParseProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetParseProgressResponseBody) SetData(v *GetParseProgressResponseBodyData) *GetParseProgressResponseBody {
	s.Data = v
	return s
}

func (s *GetParseProgressResponseBody) SetRequestId(v string) *GetParseProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParseProgressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetParseProgressResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The error code returned when an execution exception occurs.
	//
	// example:
	//
	// Package.ReadFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when an execution exception occurs.
	//
	// example:
	//
	// Failed to read skill package
	ErrorMessage    *string   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequiredEnvVars []*string `json:"RequiredEnvVars,omitempty" xml:"RequiredEnvVars,omitempty" type:"Repeated"`
	RequiresApiKey  *bool     `json:"RequiresApiKey,omitempty" xml:"RequiresApiKey,omitempty"`
	// The name in the SKILL.md file.
	//
	// example:
	//
	// name****
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The skill slug identifier. This is user-defined and unique within the tenant dimension.
	//
	// example:
	//
	// admapix******
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The task status. Valid values:
	//
	// - PARSING_METADATA: parsing in progress.
	//
	// - COMPLETED: completed.
	//
	// - FAILED: failed.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task key for parsing the skill package.
	//
	// example:
	//
	// 2E7D8B71-2677-1B4C-9E25-A88B9C5******
	TaskKey *string `json:"TaskKey,omitempty" xml:"TaskKey,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetParseProgressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetParseProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetParseProgressResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetParseProgressResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetParseProgressResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetParseProgressResponseBodyData) GetRequiredEnvVars() []*string {
	return s.RequiredEnvVars
}

func (s *GetParseProgressResponseBodyData) GetRequiresApiKey() *bool {
	return s.RequiresApiKey
}

func (s *GetParseProgressResponseBodyData) GetSkillName() *string {
	return s.SkillName
}

func (s *GetParseProgressResponseBodyData) GetSlug() *string {
	return s.Slug
}

func (s *GetParseProgressResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetParseProgressResponseBodyData) GetTaskKey() *string {
	return s.TaskKey
}

func (s *GetParseProgressResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetParseProgressResponseBodyData) SetDescription(v string) *GetParseProgressResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetParseProgressResponseBodyData) SetErrorCode(v string) *GetParseProgressResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetParseProgressResponseBodyData) SetErrorMessage(v string) *GetParseProgressResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetParseProgressResponseBodyData) SetRequiredEnvVars(v []*string) *GetParseProgressResponseBodyData {
	s.RequiredEnvVars = v
	return s
}

func (s *GetParseProgressResponseBodyData) SetRequiresApiKey(v bool) *GetParseProgressResponseBodyData {
	s.RequiresApiKey = &v
	return s
}

func (s *GetParseProgressResponseBodyData) SetSkillName(v string) *GetParseProgressResponseBodyData {
	s.SkillName = &v
	return s
}

func (s *GetParseProgressResponseBodyData) SetSlug(v string) *GetParseProgressResponseBodyData {
	s.Slug = &v
	return s
}

func (s *GetParseProgressResponseBodyData) SetStatus(v string) *GetParseProgressResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetParseProgressResponseBodyData) SetTaskKey(v string) *GetParseProgressResponseBodyData {
	s.TaskKey = &v
	return s
}

func (s *GetParseProgressResponseBodyData) SetVersion(v string) *GetParseProgressResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetParseProgressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
