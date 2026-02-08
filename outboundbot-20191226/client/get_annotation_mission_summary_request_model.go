// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnnotationMissionSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotationMissionId(v string) *GetAnnotationMissionSummaryRequest
	GetAnnotationMissionId() *string
}

type GetAnnotationMissionSummaryRequest struct {
	// Annotation job ID
	//
	// example:
	//
	// 50e53ac8-24a4-46d5-b174-ee88867f4780
	AnnotationMissionId *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
}

func (s GetAnnotationMissionSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionSummaryRequest) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *GetAnnotationMissionSummaryRequest) SetAnnotationMissionId(v string) *GetAnnotationMissionSummaryRequest {
	s.AnnotationMissionId = &v
	return s
}

func (s *GetAnnotationMissionSummaryRequest) Validate() error {
	return dara.Validate(s)
}
