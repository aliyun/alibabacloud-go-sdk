// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitApplyRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyRequest(v string) *SubmitApplyRecordRequest
	GetApplyRequest() *string
	SetCommitmentLetter(v string) *SubmitApplyRecordRequest
	GetCommitmentLetter() *string
	SetDescription(v string) *SubmitApplyRecordRequest
	GetDescription() *string
	SetEventIdList(v []*string) *SubmitApplyRecordRequest
	GetEventIdList() []*string
	SetQualificationProof(v string) *SubmitApplyRecordRequest
	GetQualificationProof() *string
	SetTrial(v bool) *SubmitApplyRecordRequest
	GetTrial() *bool
}

type SubmitApplyRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// AR01
	ApplyRequest *string `json:"ApplyRequest,omitempty" xml:"ApplyRequest,omitempty"`
	// example:
	//
	// [{\\"fileName\\":\\"5a4b4xxxxd0b6.png\\",\\"filePath\\":\\"xxx/1cxxx7d0202.png\\",\\"name\\":\\"5axxxc1d0b6.png\\"}]
	CommitmentLetter *string `json:"CommitmentLetter,omitempty" xml:"CommitmentLetter,omitempty"`
	// example:
	//
	// 已经整改，关掉相关网站。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 43029423
	EventIdList []*string `json:"EventIdList,omitempty" xml:"EventIdList,omitempty" type:"Repeated"`
	// example:
	//
	// [{\\"fileName\\":\\"5a4b4xxxxd0b6.png\\",\\"filePath\\":\\"xxx/1cxxx7d0202.png\\",\\"name\\":\\"5axxxc1d0b6.png\\"}]
	QualificationProof *string `json:"QualificationProof,omitempty" xml:"QualificationProof,omitempty"`
	// example:
	//
	// false
	Trial *bool `json:"Trial,omitempty" xml:"Trial,omitempty"`
}

func (s SubmitApplyRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitApplyRecordRequest) GoString() string {
	return s.String()
}

func (s *SubmitApplyRecordRequest) GetApplyRequest() *string {
	return s.ApplyRequest
}

func (s *SubmitApplyRecordRequest) GetCommitmentLetter() *string {
	return s.CommitmentLetter
}

func (s *SubmitApplyRecordRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitApplyRecordRequest) GetEventIdList() []*string {
	return s.EventIdList
}

func (s *SubmitApplyRecordRequest) GetQualificationProof() *string {
	return s.QualificationProof
}

func (s *SubmitApplyRecordRequest) GetTrial() *bool {
	return s.Trial
}

func (s *SubmitApplyRecordRequest) SetApplyRequest(v string) *SubmitApplyRecordRequest {
	s.ApplyRequest = &v
	return s
}

func (s *SubmitApplyRecordRequest) SetCommitmentLetter(v string) *SubmitApplyRecordRequest {
	s.CommitmentLetter = &v
	return s
}

func (s *SubmitApplyRecordRequest) SetDescription(v string) *SubmitApplyRecordRequest {
	s.Description = &v
	return s
}

func (s *SubmitApplyRecordRequest) SetEventIdList(v []*string) *SubmitApplyRecordRequest {
	s.EventIdList = v
	return s
}

func (s *SubmitApplyRecordRequest) SetQualificationProof(v string) *SubmitApplyRecordRequest {
	s.QualificationProof = &v
	return s
}

func (s *SubmitApplyRecordRequest) SetTrial(v bool) *SubmitApplyRecordRequest {
	s.Trial = &v
	return s
}

func (s *SubmitApplyRecordRequest) Validate() error {
	return dara.Validate(s)
}
