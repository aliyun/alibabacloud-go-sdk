// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountOralEvaluationStatisticsErrorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*OralEvaluationStatisticsErrorCountResponse) *CountOralEvaluationStatisticsErrorResponseBody
	GetData() []*OralEvaluationStatisticsErrorCountResponse
	SetErrCode(v string) *CountOralEvaluationStatisticsErrorResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CountOralEvaluationStatisticsErrorResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CountOralEvaluationStatisticsErrorResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CountOralEvaluationStatisticsErrorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CountOralEvaluationStatisticsErrorResponseBody
	GetSuccess() *bool
}

type CountOralEvaluationStatisticsErrorResponseBody struct {
	// example:
	//
	// []
	Data []*OralEvaluationStatisticsErrorCountResponse `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s CountOralEvaluationStatisticsErrorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CountOralEvaluationStatisticsErrorResponseBody) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) GetData() []*OralEvaluationStatisticsErrorCountResponse {
	return s.Data
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetData(v []*OralEvaluationStatisticsErrorCountResponse) *CountOralEvaluationStatisticsErrorResponseBody {
	s.Data = v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetErrCode(v string) *CountOralEvaluationStatisticsErrorResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetErrMessage(v string) *CountOralEvaluationStatisticsErrorResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetHttpStatusCode(v int32) *CountOralEvaluationStatisticsErrorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetRequestId(v string) *CountOralEvaluationStatisticsErrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetSuccess(v bool) *CountOralEvaluationStatisticsErrorResponseBody {
	s.Success = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) Validate() error {
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
