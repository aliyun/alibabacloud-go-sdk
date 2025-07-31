// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsBatchFeedbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeedback(v string) *UpdateOssCheckResultsBatchFeedbackRequest
	GetFeedback() *string
	SetItems(v string) *UpdateOssCheckResultsBatchFeedbackRequest
	GetItems() *string
	SetParentTaskId(v string) *UpdateOssCheckResultsBatchFeedbackRequest
	GetParentTaskId() *string
}

type UpdateOssCheckResultsBatchFeedbackRequest struct {
	// example:
	//
	// misreport
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// example:
	//
	// []
	Items *string `json:"Items,omitempty" xml:"Items,omitempty"`
	// example:
	//
	// P_XHDUS
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
}

func (s UpdateOssCheckResultsBatchFeedbackRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsBatchFeedbackRequest) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsBatchFeedbackRequest) GetFeedback() *string {
	return s.Feedback
}

func (s *UpdateOssCheckResultsBatchFeedbackRequest) GetItems() *string {
	return s.Items
}

func (s *UpdateOssCheckResultsBatchFeedbackRequest) GetParentTaskId() *string {
	return s.ParentTaskId
}

func (s *UpdateOssCheckResultsBatchFeedbackRequest) SetFeedback(v string) *UpdateOssCheckResultsBatchFeedbackRequest {
	s.Feedback = &v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackRequest) SetItems(v string) *UpdateOssCheckResultsBatchFeedbackRequest {
	s.Items = &v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackRequest) SetParentTaskId(v string) *UpdateOssCheckResultsBatchFeedbackRequest {
	s.ParentTaskId = &v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackRequest) Validate() error {
	return dara.Validate(s)
}
