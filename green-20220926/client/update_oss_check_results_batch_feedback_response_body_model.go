// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsBatchFeedbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvalidCount(v int32) *UpdateOssCheckResultsBatchFeedbackResponseBody
	GetInvalidCount() *int32
	SetRepeatCount(v int32) *UpdateOssCheckResultsBatchFeedbackResponseBody
	GetRepeatCount() *int32
	SetRequestId(v string) *UpdateOssCheckResultsBatchFeedbackResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *UpdateOssCheckResultsBatchFeedbackResponseBody
	GetSuccessCount() *int32
	SetTips(v string) *UpdateOssCheckResultsBatchFeedbackResponseBody
	GetTips() *string
	SetTotalCount(v int32) *UpdateOssCheckResultsBatchFeedbackResponseBody
	GetTotalCount() *int32
}

type UpdateOssCheckResultsBatchFeedbackResponseBody struct {
	// example:
	//
	// 1
	InvalidCount *int32 `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	// example:
	//
	// 1
	RepeatCount *int32 `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// xxxxx
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s UpdateOssCheckResultsBatchFeedbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsBatchFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) GetInvalidCount() *int32 {
	return s.InvalidCount
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) GetTips() *string {
	return s.Tips
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) SetInvalidCount(v int32) *UpdateOssCheckResultsBatchFeedbackResponseBody {
	s.InvalidCount = &v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) SetRepeatCount(v int32) *UpdateOssCheckResultsBatchFeedbackResponseBody {
	s.RepeatCount = &v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) SetRequestId(v string) *UpdateOssCheckResultsBatchFeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) SetSuccessCount(v int32) *UpdateOssCheckResultsBatchFeedbackResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) SetTips(v string) *UpdateOssCheckResultsBatchFeedbackResponseBody {
	s.Tips = &v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) SetTotalCount(v int32) *UpdateOssCheckResultsBatchFeedbackResponseBody {
	s.TotalCount = &v
	return s
}

func (s *UpdateOssCheckResultsBatchFeedbackResponseBody) Validate() error {
	return dara.Validate(s)
}
