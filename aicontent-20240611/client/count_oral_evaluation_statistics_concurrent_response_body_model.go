// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountOralEvaluationStatisticsConcurrentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*OralEvaluationStatisticsConcurrentCountResponse) *CountOralEvaluationStatisticsConcurrentResponseBody
	GetData() []*OralEvaluationStatisticsConcurrentCountResponse
	SetErrCode(v string) *CountOralEvaluationStatisticsConcurrentResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CountOralEvaluationStatisticsConcurrentResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CountOralEvaluationStatisticsConcurrentResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CountOralEvaluationStatisticsConcurrentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CountOralEvaluationStatisticsConcurrentResponseBody
	GetSuccess() *bool
}

type CountOralEvaluationStatisticsConcurrentResponseBody struct {
	// example:
	//
	// []
	Data []*OralEvaluationStatisticsConcurrentCountResponse `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CountOralEvaluationStatisticsConcurrentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CountOralEvaluationStatisticsConcurrentResponseBody) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) GetData() []*OralEvaluationStatisticsConcurrentCountResponse {
	return s.Data
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetData(v []*OralEvaluationStatisticsConcurrentCountResponse) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.Data = v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetErrCode(v string) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetErrMessage(v string) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetHttpStatusCode(v int32) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetRequestId(v string) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetSuccess(v bool) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.Success = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
