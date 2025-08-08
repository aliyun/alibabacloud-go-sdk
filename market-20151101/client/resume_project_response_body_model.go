// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeProjectResponseBody
	GetRequestId() *string
	SetResult(v bool) *ResumeProjectResponseBody
	GetResult() *bool
	SetSuccess(v bool) *ResumeProjectResponseBody
	GetSuccess() *bool
}

type ResumeProjectResponseBody struct {
	// example:
	//
	// ee69a00f-189b-400f-9fd2-af89749fb50f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeProjectResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ResumeProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeProjectResponseBody) SetRequestId(v string) *ResumeProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeProjectResponseBody) SetResult(v bool) *ResumeProjectResponseBody {
	s.Result = &v
	return s
}

func (s *ResumeProjectResponseBody) SetSuccess(v bool) *ResumeProjectResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
