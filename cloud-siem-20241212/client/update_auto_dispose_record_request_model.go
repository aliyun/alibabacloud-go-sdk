// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoDisposeRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDecisionConclusion(v string) *UpdateAutoDisposeRecordRequest
	GetAutoDecisionConclusion() *string
	SetAutoDecisionEntityList(v string) *UpdateAutoDisposeRecordRequest
	GetAutoDecisionEntityList() *string
	SetAutoDecisionResult(v string) *UpdateAutoDisposeRecordRequest
	GetAutoDecisionResult() *string
	SetAutoDisposeRecordId(v string) *UpdateAutoDisposeRecordRequest
	GetAutoDisposeRecordId() *string
	SetLang(v string) *UpdateAutoDisposeRecordRequest
	GetLang() *string
}

type UpdateAutoDisposeRecordRequest struct {
	AutoDecisionConclusion *string `json:"AutoDecisionConclusion,omitempty" xml:"AutoDecisionConclusion,omitempty"`
	// example:
	//
	// [{"entityType":"file","entityName":"/path/file.file","entityUuid":"b7efb45ce7ff09758****","disposalMethod":"delete","playbookUuid":"9213bhdjagdja****"}]
	AutoDecisionEntityList *string `json:"AutoDecisionEntityList,omitempty" xml:"AutoDecisionEntityList,omitempty"`
	// example:
	//
	// success
	AutoDecisionResult *string `json:"AutoDecisionResult,omitempty" xml:"AutoDecisionResult,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4227e0cdc4b7efb45ce7ff09758****
	AutoDisposeRecordId *string `json:"AutoDisposeRecordId,omitempty" xml:"AutoDisposeRecordId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s UpdateAutoDisposeRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoDisposeRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoDisposeRecordRequest) GetAutoDecisionConclusion() *string {
	return s.AutoDecisionConclusion
}

func (s *UpdateAutoDisposeRecordRequest) GetAutoDecisionEntityList() *string {
	return s.AutoDecisionEntityList
}

func (s *UpdateAutoDisposeRecordRequest) GetAutoDecisionResult() *string {
	return s.AutoDecisionResult
}

func (s *UpdateAutoDisposeRecordRequest) GetAutoDisposeRecordId() *string {
	return s.AutoDisposeRecordId
}

func (s *UpdateAutoDisposeRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAutoDisposeRecordRequest) SetAutoDecisionConclusion(v string) *UpdateAutoDisposeRecordRequest {
	s.AutoDecisionConclusion = &v
	return s
}

func (s *UpdateAutoDisposeRecordRequest) SetAutoDecisionEntityList(v string) *UpdateAutoDisposeRecordRequest {
	s.AutoDecisionEntityList = &v
	return s
}

func (s *UpdateAutoDisposeRecordRequest) SetAutoDecisionResult(v string) *UpdateAutoDisposeRecordRequest {
	s.AutoDecisionResult = &v
	return s
}

func (s *UpdateAutoDisposeRecordRequest) SetAutoDisposeRecordId(v string) *UpdateAutoDisposeRecordRequest {
	s.AutoDisposeRecordId = &v
	return s
}

func (s *UpdateAutoDisposeRecordRequest) SetLang(v string) *UpdateAutoDisposeRecordRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAutoDisposeRecordRequest) Validate() error {
	return dara.Validate(s)
}
