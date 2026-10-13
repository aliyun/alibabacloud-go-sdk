// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitScriptPreReviewResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SubmitScriptPreReviewResultRequest
	GetInstanceId() *string
	SetRecordId(v string) *SubmitScriptPreReviewResultRequest
	GetRecordId() *string
	SetRejectReason(v string) *SubmitScriptPreReviewResultRequest
	GetRejectReason() *string
	SetReviewResult(v string) *SubmitScriptPreReviewResultRequest
	GetReviewResult() *string
}

type SubmitScriptPreReviewResultRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bdd49242-114c-4045-b1d1-25ccc1756c75
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the script review record.
	//
	// This parameter is required.
	//
	// example:
	//
	// b8e1a4d3-9c2f-4e6a-b7d5-3f1e2a4c6b8d
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The reason for pre-review rejection. This parameter is required when ReviewResult is set to REJECT.
	//
	// example:
	//
	// The script content involves sensitive information
	RejectReason *string `json:"RejectReason,omitempty" xml:"RejectReason,omitempty"`
	// The pre-review result. Valid values: PASS (approved) and REJECT (rejected). The value is not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// PASS
	ReviewResult *string `json:"ReviewResult,omitempty" xml:"ReviewResult,omitempty"`
}

func (s SubmitScriptPreReviewResultRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitScriptPreReviewResultRequest) GoString() string {
	return s.String()
}

func (s *SubmitScriptPreReviewResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SubmitScriptPreReviewResultRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *SubmitScriptPreReviewResultRequest) GetRejectReason() *string {
	return s.RejectReason
}

func (s *SubmitScriptPreReviewResultRequest) GetReviewResult() *string {
	return s.ReviewResult
}

func (s *SubmitScriptPreReviewResultRequest) SetInstanceId(v string) *SubmitScriptPreReviewResultRequest {
	s.InstanceId = &v
	return s
}

func (s *SubmitScriptPreReviewResultRequest) SetRecordId(v string) *SubmitScriptPreReviewResultRequest {
	s.RecordId = &v
	return s
}

func (s *SubmitScriptPreReviewResultRequest) SetRejectReason(v string) *SubmitScriptPreReviewResultRequest {
	s.RejectReason = &v
	return s
}

func (s *SubmitScriptPreReviewResultRequest) SetReviewResult(v string) *SubmitScriptPreReviewResultRequest {
	s.ReviewResult = &v
	return s
}

func (s *SubmitScriptPreReviewResultRequest) Validate() error {
	return dara.Validate(s)
}
