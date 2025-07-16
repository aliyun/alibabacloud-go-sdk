// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFeatureViewTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishFeatureViewTableResponseBody
	GetRequestId() *string
	SetTaskId(v string) *PublishFeatureViewTableResponseBody
	GetTaskId() *string
}

type PublishFeatureViewTableResponseBody struct {
	// example:
	//
	// 627B5776-4D06-5A49-8A04-508AA39653F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PublishFeatureViewTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishFeatureViewTableResponseBody) GoString() string {
	return s.String()
}

func (s *PublishFeatureViewTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishFeatureViewTableResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *PublishFeatureViewTableResponseBody) SetRequestId(v string) *PublishFeatureViewTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishFeatureViewTableResponseBody) SetTaskId(v string) *PublishFeatureViewTableResponseBody {
	s.TaskId = &v
	return s
}

func (s *PublishFeatureViewTableResponseBody) Validate() error {
	return dara.Validate(s)
}
