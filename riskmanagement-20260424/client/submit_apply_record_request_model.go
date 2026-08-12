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
	EventIdList []*string `json:"EventIdList,omitempty" xml:"EventIdList,omitempty" type:"Repeated"`
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
