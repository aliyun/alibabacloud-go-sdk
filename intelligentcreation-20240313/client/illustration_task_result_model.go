// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIllustrationTaskResult interface {
	dara.Model
	String() string
	GoString() string
	SetIllustrationTask(v *IllustrationTask) *IllustrationTaskResult
	GetIllustrationTask() *IllustrationTask
	SetRequestId(v string) *IllustrationTaskResult
	GetRequestId() *string
}

type IllustrationTaskResult struct {
	IllustrationTask *IllustrationTask `json:"illustrationTask,omitempty" xml:"illustrationTask,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s IllustrationTaskResult) String() string {
	return dara.Prettify(s)
}

func (s IllustrationTaskResult) GoString() string {
	return s.String()
}

func (s *IllustrationTaskResult) GetIllustrationTask() *IllustrationTask {
	return s.IllustrationTask
}

func (s *IllustrationTaskResult) GetRequestId() *string {
	return s.RequestId
}

func (s *IllustrationTaskResult) SetIllustrationTask(v *IllustrationTask) *IllustrationTaskResult {
	s.IllustrationTask = v
	return s
}

func (s *IllustrationTaskResult) SetRequestId(v string) *IllustrationTaskResult {
	s.RequestId = &v
	return s
}

func (s *IllustrationTaskResult) Validate() error {
	if s.IllustrationTask != nil {
		if err := s.IllustrationTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}
