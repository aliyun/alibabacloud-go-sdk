// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitApplyRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyRequest(v string) *SubmitApplyRecordShrinkRequest
	GetApplyRequest() *string
	SetCommitmentLetter(v string) *SubmitApplyRecordShrinkRequest
	GetCommitmentLetter() *string
	SetDescription(v string) *SubmitApplyRecordShrinkRequest
	GetDescription() *string
	SetEventIdListShrink(v string) *SubmitApplyRecordShrinkRequest
	GetEventIdListShrink() *string
	SetQualificationProof(v string) *SubmitApplyRecordShrinkRequest
	GetQualificationProof() *string
	SetTrial(v bool) *SubmitApplyRecordShrinkRequest
	GetTrial() *bool
}

type SubmitApplyRecordShrinkRequest struct {
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
	EventIdListShrink *string `json:"EventIdList,omitempty" xml:"EventIdList,omitempty"`
	// example:
	//
	// [{\\"fileName\\":\\"5a4b4xxxxd0b6.png\\",\\"filePath\\":\\"xxx/1cxxx7d0202.png\\",\\"name\\":\\"5axxxc1d0b6.png\\"}]
	QualificationProof *string `json:"QualificationProof,omitempty" xml:"QualificationProof,omitempty"`
	// example:
	//
	// false
	Trial *bool `json:"Trial,omitempty" xml:"Trial,omitempty"`
}

func (s SubmitApplyRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitApplyRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitApplyRecordShrinkRequest) GetApplyRequest() *string {
	return s.ApplyRequest
}

func (s *SubmitApplyRecordShrinkRequest) GetCommitmentLetter() *string {
	return s.CommitmentLetter
}

func (s *SubmitApplyRecordShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitApplyRecordShrinkRequest) GetEventIdListShrink() *string {
	return s.EventIdListShrink
}

func (s *SubmitApplyRecordShrinkRequest) GetQualificationProof() *string {
	return s.QualificationProof
}

func (s *SubmitApplyRecordShrinkRequest) GetTrial() *bool {
	return s.Trial
}

func (s *SubmitApplyRecordShrinkRequest) SetApplyRequest(v string) *SubmitApplyRecordShrinkRequest {
	s.ApplyRequest = &v
	return s
}

func (s *SubmitApplyRecordShrinkRequest) SetCommitmentLetter(v string) *SubmitApplyRecordShrinkRequest {
	s.CommitmentLetter = &v
	return s
}

func (s *SubmitApplyRecordShrinkRequest) SetDescription(v string) *SubmitApplyRecordShrinkRequest {
	s.Description = &v
	return s
}

func (s *SubmitApplyRecordShrinkRequest) SetEventIdListShrink(v string) *SubmitApplyRecordShrinkRequest {
	s.EventIdListShrink = &v
	return s
}

func (s *SubmitApplyRecordShrinkRequest) SetQualificationProof(v string) *SubmitApplyRecordShrinkRequest {
	s.QualificationProof = &v
	return s
}

func (s *SubmitApplyRecordShrinkRequest) SetTrial(v bool) *SubmitApplyRecordShrinkRequest {
	s.Trial = &v
	return s
}

func (s *SubmitApplyRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
