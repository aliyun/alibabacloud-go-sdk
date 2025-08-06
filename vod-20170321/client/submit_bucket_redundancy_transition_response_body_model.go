// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBucketRedundancyTransitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitBucketRedundancyTransitionResponseBody
	GetRequestId() *string
	SetTaskId(v string) *SubmitBucketRedundancyTransitionResponseBody
	GetTaskId() *string
}

type SubmitBucketRedundancyTransitionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitBucketRedundancyTransitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitBucketRedundancyTransitionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitBucketRedundancyTransitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitBucketRedundancyTransitionResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitBucketRedundancyTransitionResponseBody) SetRequestId(v string) *SubmitBucketRedundancyTransitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitBucketRedundancyTransitionResponseBody) SetTaskId(v string) *SubmitBucketRedundancyTransitionResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitBucketRedundancyTransitionResponseBody) Validate() error {
	return dara.Validate(s)
}
