// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateAICoachTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchCreateAICoachTaskRequest
	GetRequestId() *string
	SetScriptRecordId(v string) *BatchCreateAICoachTaskRequest
	GetScriptRecordId() *string
	SetStudentIds(v []*string) *BatchCreateAICoachTaskRequest
	GetStudentIds() []*string
	SetStudentList(v []*BatchCreateAICoachTaskRequestStudentList) *BatchCreateAICoachTaskRequest
	GetStudentList() []*BatchCreateAICoachTaskRequestStudentList
}

type BatchCreateAICoachTaskRequest struct {
	// example:
	//
	// 7915125A-0D96-5A25-A54B-D3B739A86AFC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	ScriptRecordId *string                                     `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	StudentIds     []*string                                   `json:"studentIds,omitempty" xml:"studentIds,omitempty" type:"Repeated"`
	StudentList    []*BatchCreateAICoachTaskRequestStudentList `json:"studentList,omitempty" xml:"studentList,omitempty" type:"Repeated"`
}

func (s BatchCreateAICoachTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateAICoachTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateAICoachTaskRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateAICoachTaskRequest) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *BatchCreateAICoachTaskRequest) GetStudentIds() []*string {
	return s.StudentIds
}

func (s *BatchCreateAICoachTaskRequest) GetStudentList() []*BatchCreateAICoachTaskRequestStudentList {
	return s.StudentList
}

func (s *BatchCreateAICoachTaskRequest) SetRequestId(v string) *BatchCreateAICoachTaskRequest {
	s.RequestId = &v
	return s
}

func (s *BatchCreateAICoachTaskRequest) SetScriptRecordId(v string) *BatchCreateAICoachTaskRequest {
	s.ScriptRecordId = &v
	return s
}

func (s *BatchCreateAICoachTaskRequest) SetStudentIds(v []*string) *BatchCreateAICoachTaskRequest {
	s.StudentIds = v
	return s
}

func (s *BatchCreateAICoachTaskRequest) SetStudentList(v []*BatchCreateAICoachTaskRequestStudentList) *BatchCreateAICoachTaskRequest {
	s.StudentList = v
	return s
}

func (s *BatchCreateAICoachTaskRequest) Validate() error {
	if s.StudentList != nil {
		for _, item := range s.StudentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateAICoachTaskRequestStudentList struct {
	StudentAudioUrl *string `json:"studentAudioUrl,omitempty" xml:"studentAudioUrl,omitempty"`
	StudentId       *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
}

func (s BatchCreateAICoachTaskRequestStudentList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateAICoachTaskRequestStudentList) GoString() string {
	return s.String()
}

func (s *BatchCreateAICoachTaskRequestStudentList) GetStudentAudioUrl() *string {
	return s.StudentAudioUrl
}

func (s *BatchCreateAICoachTaskRequestStudentList) GetStudentId() *string {
	return s.StudentId
}

func (s *BatchCreateAICoachTaskRequestStudentList) SetStudentAudioUrl(v string) *BatchCreateAICoachTaskRequestStudentList {
	s.StudentAudioUrl = &v
	return s
}

func (s *BatchCreateAICoachTaskRequestStudentList) SetStudentId(v string) *BatchCreateAICoachTaskRequestStudentList {
	s.StudentId = &v
	return s
}

func (s *BatchCreateAICoachTaskRequestStudentList) Validate() error {
	return dara.Validate(s)
}
