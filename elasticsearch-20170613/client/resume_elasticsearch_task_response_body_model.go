// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeElasticsearchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeElasticsearchTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *ResumeElasticsearchTaskResponseBody
	GetResult() *bool
}

type ResumeElasticsearchTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return result:
	//
	// - true: The interrupted change was resumed successfully.
	//
	// - false: The interrupted change failed to be resumed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ResumeElasticsearchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeElasticsearchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeElasticsearchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeElasticsearchTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ResumeElasticsearchTaskResponseBody) SetRequestId(v string) *ResumeElasticsearchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeElasticsearchTaskResponseBody) SetResult(v bool) *ResumeElasticsearchTaskResponseBody {
	s.Result = &v
	return s
}

func (s *ResumeElasticsearchTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
