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
	// The request reason.
	//
	// - **AR01**: Rectified. Request to unblock.
	//
	// - **AR02**: No violation found after investigation.
	//
	// - **AR03**: The instance or service has been shut down and cannot be operated. Request to unblock and then clear the violation information.
	//
	// - **AR04**: Files deleted. Request to unblock.
	//
	// - **AR05**: The instance has been released.
	//
	// - **AR00**: Other. Provide a description.
	//
	// This parameter is required.
	//
	// example:
	//
	// AR01
	ApplyRequest *string `json:"ApplyRequest,omitempty" xml:"ApplyRequest,omitempty"`
	// The commitment letter.
	//
	// example:
	//
	// [{\\"fileName\\":\\"5a4b4xxxxd0b6.png\\",\\"filePath\\":\\"xxx/1cxxx7d0202.png\\",\\"name\\":\\"5axxxc1d0b6.png\\"}]
	CommitmentLetter *string `json:"CommitmentLetter,omitempty" xml:"CommitmentLetter,omitempty"`
	// The description of the situation.
	//
	// example:
	//
	// Rectification completed. Related websites have been shut down.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of specified event IDs.
	//
	// example:
	//
	// 43029423
	EventIdListShrink *string `json:"EventIdList,omitempty" xml:"EventIdList,omitempty"`
	// The qualification proof.
	//
	// example:
	//
	// [{\\"fileName\\":\\"5a4b4xxxxd0b6.png\\",\\"filePath\\":\\"xxx/1cxxx7d0202.png\\",\\"name\\":\\"5axxxc1d0b6.png\\"}]
	QualificationProof *string `json:"QualificationProof,omitempty" xml:"QualificationProof,omitempty"`
	// Specifies whether manual review is required.
	//
	// - **true**: Manual review is required.
	//
	// - **false**: Manual review is not required.
	//
	// > Default value: manual review is not required.
	//
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
