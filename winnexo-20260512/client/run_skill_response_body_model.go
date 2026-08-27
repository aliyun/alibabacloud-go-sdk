// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunSkillResponseBody
	GetCode() *string
	SetCreatedAt(v string) *RunSkillResponseBody
	GetCreatedAt() *string
	SetMessage(v string) *RunSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunSkillResponseBody
	GetRequestId() *string
	SetRunId(v string) *RunSkillResponseBody
	GetRunId() *string
	SetSkillCode(v string) *RunSkillResponseBody
	GetSkillCode() *string
	SetSkillName(v string) *RunSkillResponseBody
	GetSkillName() *string
	SetStatus(v string) *RunSkillResponseBody
	GetStatus() *string
}

type RunSkillResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The task creation time in ISO 8601 UTC format.
	//
	// example:
	//
	// string_value
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The asynchronous task ID, used for querying with getSkillRun.
	//
	// example:
	//
	// exampleRunId
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The skill code that was actually executed.
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// The skill name.
	//
	// example:
	//
	// string_value
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The task status. Returns Running immediately upon submission.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s RunSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunSkillResponseBody) GoString() string {
	return s.String()
}

func (s *RunSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunSkillResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *RunSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunSkillResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *RunSkillResponseBody) GetSkillCode() *string {
	return s.SkillCode
}

func (s *RunSkillResponseBody) GetSkillName() *string {
	return s.SkillName
}

func (s *RunSkillResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RunSkillResponseBody) SetCode(v string) *RunSkillResponseBody {
	s.Code = &v
	return s
}

func (s *RunSkillResponseBody) SetCreatedAt(v string) *RunSkillResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *RunSkillResponseBody) SetMessage(v string) *RunSkillResponseBody {
	s.Message = &v
	return s
}

func (s *RunSkillResponseBody) SetRequestId(v string) *RunSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSkillResponseBody) SetRunId(v string) *RunSkillResponseBody {
	s.RunId = &v
	return s
}

func (s *RunSkillResponseBody) SetSkillCode(v string) *RunSkillResponseBody {
	s.SkillCode = &v
	return s
}

func (s *RunSkillResponseBody) SetSkillName(v string) *RunSkillResponseBody {
	s.SkillName = &v
	return s
}

func (s *RunSkillResponseBody) SetStatus(v string) *RunSkillResponseBody {
	s.Status = &v
	return s
}

func (s *RunSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
