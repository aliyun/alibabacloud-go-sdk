// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSolutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSolutionResponseBody
	GetRequestId() *string
	SetSolutionId(v int64) *CreateSolutionResponseBody
	GetSolutionId() *int64
}

type CreateSolutionResponseBody struct {
	// example:
	//
	// F55D90C1-31BE-4B2A-AA3F-25EFC36F9419
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100001089003
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s CreateSolutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSolutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSolutionResponseBody) GetSolutionId() *int64 {
	return s.SolutionId
}

func (s *CreateSolutionResponseBody) SetRequestId(v string) *CreateSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSolutionResponseBody) SetSolutionId(v int64) *CreateSolutionResponseBody {
	s.SolutionId = &v
	return s
}

func (s *CreateSolutionResponseBody) Validate() error {
	return dara.Validate(s)
}
