// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGenerateAICoachScriptTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssessmentPoint(v bool) *CreateGenerateAICoachScriptTaskRequest
	GetAssessmentPoint() *bool
	SetDescription(v string) *CreateGenerateAICoachScriptTaskRequest
	GetDescription() *string
	SetDialogueKey(v string) *CreateGenerateAICoachScriptTaskRequest
	GetDialogueKey() *string
	SetDialogueUrl(v string) *CreateGenerateAICoachScriptTaskRequest
	GetDialogueUrl() *string
	SetDocList(v []*CreateGenerateAICoachScriptTaskRequestDocList) *CreateGenerateAICoachScriptTaskRequest
	GetDocList() []*CreateGenerateAICoachScriptTaskRequestDocList
	SetDocUrlList(v []*string) *CreateGenerateAICoachScriptTaskRequest
	GetDocUrlList() []*string
	SetScriptName(v string) *CreateGenerateAICoachScriptTaskRequest
	GetScriptName() *string
}

type CreateGenerateAICoachScriptTaskRequest struct {
	AssessmentPoint *bool                                            `json:"assessmentPoint,omitempty" xml:"assessmentPoint,omitempty"`
	Description     *string                                          `json:"description,omitempty" xml:"description,omitempty"`
	DialogueKey     *string                                          `json:"dialogueKey,omitempty" xml:"dialogueKey,omitempty"`
	DialogueUrl     *string                                          `json:"dialogueUrl,omitempty" xml:"dialogueUrl,omitempty"`
	DocList         []*CreateGenerateAICoachScriptTaskRequestDocList `json:"docList,omitempty" xml:"docList,omitempty" type:"Repeated"`
	DocUrlList      []*string                                        `json:"docUrlList,omitempty" xml:"docUrlList,omitempty" type:"Repeated"`
	ScriptName      *string                                          `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
}

func (s CreateGenerateAICoachScriptTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGenerateAICoachScriptTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateGenerateAICoachScriptTaskRequest) GetAssessmentPoint() *bool {
	return s.AssessmentPoint
}

func (s *CreateGenerateAICoachScriptTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGenerateAICoachScriptTaskRequest) GetDialogueKey() *string {
	return s.DialogueKey
}

func (s *CreateGenerateAICoachScriptTaskRequest) GetDialogueUrl() *string {
	return s.DialogueUrl
}

func (s *CreateGenerateAICoachScriptTaskRequest) GetDocList() []*CreateGenerateAICoachScriptTaskRequestDocList {
	return s.DocList
}

func (s *CreateGenerateAICoachScriptTaskRequest) GetDocUrlList() []*string {
	return s.DocUrlList
}

func (s *CreateGenerateAICoachScriptTaskRequest) GetScriptName() *string {
	return s.ScriptName
}

func (s *CreateGenerateAICoachScriptTaskRequest) SetAssessmentPoint(v bool) *CreateGenerateAICoachScriptTaskRequest {
	s.AssessmentPoint = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskRequest) SetDescription(v string) *CreateGenerateAICoachScriptTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskRequest) SetDialogueKey(v string) *CreateGenerateAICoachScriptTaskRequest {
	s.DialogueKey = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskRequest) SetDialogueUrl(v string) *CreateGenerateAICoachScriptTaskRequest {
	s.DialogueUrl = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskRequest) SetDocList(v []*CreateGenerateAICoachScriptTaskRequestDocList) *CreateGenerateAICoachScriptTaskRequest {
	s.DocList = v
	return s
}

func (s *CreateGenerateAICoachScriptTaskRequest) SetDocUrlList(v []*string) *CreateGenerateAICoachScriptTaskRequest {
	s.DocUrlList = v
	return s
}

func (s *CreateGenerateAICoachScriptTaskRequest) SetScriptName(v string) *CreateGenerateAICoachScriptTaskRequest {
	s.ScriptName = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskRequest) Validate() error {
	if s.DocList != nil {
		for _, item := range s.DocList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateGenerateAICoachScriptTaskRequestDocList struct {
	DocId   *string `json:"docId,omitempty" xml:"docId,omitempty"`
	DocName *string `json:"docName,omitempty" xml:"docName,omitempty"`
	KbId    *string `json:"kbId,omitempty" xml:"kbId,omitempty"`
}

func (s CreateGenerateAICoachScriptTaskRequestDocList) String() string {
	return dara.Prettify(s)
}

func (s CreateGenerateAICoachScriptTaskRequestDocList) GoString() string {
	return s.String()
}

func (s *CreateGenerateAICoachScriptTaskRequestDocList) GetDocId() *string {
	return s.DocId
}

func (s *CreateGenerateAICoachScriptTaskRequestDocList) GetDocName() *string {
	return s.DocName
}

func (s *CreateGenerateAICoachScriptTaskRequestDocList) GetKbId() *string {
	return s.KbId
}

func (s *CreateGenerateAICoachScriptTaskRequestDocList) SetDocId(v string) *CreateGenerateAICoachScriptTaskRequestDocList {
	s.DocId = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskRequestDocList) SetDocName(v string) *CreateGenerateAICoachScriptTaskRequestDocList {
	s.DocName = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskRequestDocList) SetKbId(v string) *CreateGenerateAICoachScriptTaskRequestDocList {
	s.KbId = &v
	return s
}

func (s *CreateGenerateAICoachScriptTaskRequestDocList) Validate() error {
	return dara.Validate(s)
}
