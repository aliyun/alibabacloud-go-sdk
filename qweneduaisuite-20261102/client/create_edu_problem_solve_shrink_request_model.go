// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEduProblemSolveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueryShrink(v string) *CreateEduProblemSolveShrinkRequest
	GetQueryShrink() *string
	SetRequestId(v string) *CreateEduProblemSolveShrinkRequest
	GetRequestId() *string
}

type CreateEduProblemSolveShrinkRequest struct {
	// This parameter is required.
	QueryShrink *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// 019F593F-D6EE-5F44-819A-6CDA127D1637
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateEduProblemSolveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEduProblemSolveShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEduProblemSolveShrinkRequest) GetQueryShrink() *string {
	return s.QueryShrink
}

func (s *CreateEduProblemSolveShrinkRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEduProblemSolveShrinkRequest) SetQueryShrink(v string) *CreateEduProblemSolveShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *CreateEduProblemSolveShrinkRequest) SetRequestId(v string) *CreateEduProblemSolveShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *CreateEduProblemSolveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
