// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectSolutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *RejectSolutionRequest
	GetBizType() *string
	SetNote(v string) *RejectSolutionRequest
	GetNote() *string
	SetSolutionBizId(v string) *RejectSolutionRequest
	GetSolutionBizId() *string
}

type RejectSolutionRequest struct {
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// S20200512114050000001
	SolutionBizId *string `json:"SolutionBizId,omitempty" xml:"SolutionBizId,omitempty"`
}

func (s RejectSolutionRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectSolutionRequest) GoString() string {
	return s.String()
}

func (s *RejectSolutionRequest) GetBizType() *string {
	return s.BizType
}

func (s *RejectSolutionRequest) GetNote() *string {
	return s.Note
}

func (s *RejectSolutionRequest) GetSolutionBizId() *string {
	return s.SolutionBizId
}

func (s *RejectSolutionRequest) SetBizType(v string) *RejectSolutionRequest {
	s.BizType = &v
	return s
}

func (s *RejectSolutionRequest) SetNote(v string) *RejectSolutionRequest {
	s.Note = &v
	return s
}

func (s *RejectSolutionRequest) SetSolutionBizId(v string) *RejectSolutionRequest {
	s.SolutionBizId = &v
	return s
}

func (s *RejectSolutionRequest) Validate() error {
	return dara.Validate(s)
}
