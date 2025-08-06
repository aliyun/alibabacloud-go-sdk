// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectUserSolutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *RejectUserSolutionRequest
	GetBizType() *string
	SetNote(v string) *RejectUserSolutionRequest
	GetNote() *string
	SetSolutionBizId(v string) *RejectUserSolutionRequest
	GetSolutionBizId() *string
}

type RejectUserSolutionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.companyreg
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// S20211227151912000001
	SolutionBizId *string `json:"SolutionBizId,omitempty" xml:"SolutionBizId,omitempty"`
}

func (s RejectUserSolutionRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectUserSolutionRequest) GoString() string {
	return s.String()
}

func (s *RejectUserSolutionRequest) GetBizType() *string {
	return s.BizType
}

func (s *RejectUserSolutionRequest) GetNote() *string {
	return s.Note
}

func (s *RejectUserSolutionRequest) GetSolutionBizId() *string {
	return s.SolutionBizId
}

func (s *RejectUserSolutionRequest) SetBizType(v string) *RejectUserSolutionRequest {
	s.BizType = &v
	return s
}

func (s *RejectUserSolutionRequest) SetNote(v string) *RejectUserSolutionRequest {
	s.Note = &v
	return s
}

func (s *RejectUserSolutionRequest) SetSolutionBizId(v string) *RejectUserSolutionRequest {
	s.SolutionBizId = &v
	return s
}

func (s *RejectUserSolutionRequest) Validate() error {
	return dara.Validate(s)
}
