// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachAssessmentPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPointId(v string) *GetAICoachAssessmentPointRequest
	GetPointId() *string
}

type GetAICoachAssessmentPointRequest struct {
	// example:
	//
	// 1
	PointId *string `json:"pointId,omitempty" xml:"pointId,omitempty"`
}

func (s GetAICoachAssessmentPointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachAssessmentPointRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointRequest) GetPointId() *string {
	return s.PointId
}

func (s *GetAICoachAssessmentPointRequest) SetPointId(v string) *GetAICoachAssessmentPointRequest {
	s.PointId = &v
	return s
}

func (s *GetAICoachAssessmentPointRequest) Validate() error {
	return dara.Validate(s)
}
