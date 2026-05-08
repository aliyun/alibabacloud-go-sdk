// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICoachTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAICoachTaskRequest
	GetRequestId() *string
	SetScriptRecordId(v string) *CreateAICoachTaskRequest
	GetScriptRecordId() *string
	SetStudentAudioUrl(v string) *CreateAICoachTaskRequest
	GetStudentAudioUrl() *string
	SetStudentId(v string) *CreateAICoachTaskRequest
	GetStudentId() *string
}

type CreateAICoachTaskRequest struct {
	// example:
	//
	// 541E7123-2E8A-5BA2-AC38-665650C84129
	RequestId       *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScriptRecordId  *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	StudentAudioUrl *string `json:"studentAudioUrl,omitempty" xml:"studentAudioUrl,omitempty"`
	StudentId       *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
}

func (s CreateAICoachTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAICoachTaskRequest) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *CreateAICoachTaskRequest) GetStudentAudioUrl() *string {
	return s.StudentAudioUrl
}

func (s *CreateAICoachTaskRequest) GetStudentId() *string {
	return s.StudentId
}

func (s *CreateAICoachTaskRequest) SetRequestId(v string) *CreateAICoachTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateAICoachTaskRequest) SetScriptRecordId(v string) *CreateAICoachTaskRequest {
	s.ScriptRecordId = &v
	return s
}

func (s *CreateAICoachTaskRequest) SetStudentAudioUrl(v string) *CreateAICoachTaskRequest {
	s.StudentAudioUrl = &v
	return s
}

func (s *CreateAICoachTaskRequest) SetStudentId(v string) *CreateAICoachTaskRequest {
	s.StudentId = &v
	return s
}

func (s *CreateAICoachTaskRequest) Validate() error {
	return dara.Validate(s)
}
