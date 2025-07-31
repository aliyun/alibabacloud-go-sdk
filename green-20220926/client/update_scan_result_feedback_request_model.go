// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScanResultFeedbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeedback(v string) *UpdateScanResultFeedbackRequest
	GetFeedback() *string
	SetLabels(v string) *UpdateScanResultFeedbackRequest
	GetLabels() *string
	SetQueryRequestId(v string) *UpdateScanResultFeedbackRequest
	GetQueryRequestId() *string
	SetRegionId(v string) *UpdateScanResultFeedbackRequest
	GetRegionId() *string
	SetResourceType(v string) *UpdateScanResultFeedbackRequest
	GetResourceType() *string
	SetRiskLevel(v string) *UpdateScanResultFeedbackRequest
	GetRiskLevel() *string
}

type UpdateScanResultFeedbackRequest struct {
	// example:
	//
	// missOut
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	Labels   *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 46232656-984E-****-A648-B1D0667B6C3E
	QueryRequestId *string `json:"QueryRequestId,omitempty" xml:"QueryRequestId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RiskLevel    *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s UpdateScanResultFeedbackRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScanResultFeedbackRequest) GoString() string {
	return s.String()
}

func (s *UpdateScanResultFeedbackRequest) GetFeedback() *string {
	return s.Feedback
}

func (s *UpdateScanResultFeedbackRequest) GetLabels() *string {
	return s.Labels
}

func (s *UpdateScanResultFeedbackRequest) GetQueryRequestId() *string {
	return s.QueryRequestId
}

func (s *UpdateScanResultFeedbackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateScanResultFeedbackRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateScanResultFeedbackRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *UpdateScanResultFeedbackRequest) SetFeedback(v string) *UpdateScanResultFeedbackRequest {
	s.Feedback = &v
	return s
}

func (s *UpdateScanResultFeedbackRequest) SetLabels(v string) *UpdateScanResultFeedbackRequest {
	s.Labels = &v
	return s
}

func (s *UpdateScanResultFeedbackRequest) SetQueryRequestId(v string) *UpdateScanResultFeedbackRequest {
	s.QueryRequestId = &v
	return s
}

func (s *UpdateScanResultFeedbackRequest) SetRegionId(v string) *UpdateScanResultFeedbackRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateScanResultFeedbackRequest) SetResourceType(v string) *UpdateScanResultFeedbackRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateScanResultFeedbackRequest) SetRiskLevel(v string) *UpdateScanResultFeedbackRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateScanResultFeedbackRequest) Validate() error {
	return dara.Validate(s)
}
