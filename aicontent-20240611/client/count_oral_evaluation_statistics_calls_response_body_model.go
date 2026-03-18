// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountOralEvaluationStatisticsCallsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*OralEvaluationStatisticsCallsCountResponse) *CountOralEvaluationStatisticsCallsResponseBody
	GetData() []*OralEvaluationStatisticsCallsCountResponse
	SetErrCode(v string) *CountOralEvaluationStatisticsCallsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CountOralEvaluationStatisticsCallsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CountOralEvaluationStatisticsCallsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CountOralEvaluationStatisticsCallsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CountOralEvaluationStatisticsCallsResponseBody
	GetSuccess() *bool
}

type CountOralEvaluationStatisticsCallsResponseBody struct {
	// example:
	//
	// []
	Data []*OralEvaluationStatisticsCallsCountResponse `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s CountOralEvaluationStatisticsCallsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CountOralEvaluationStatisticsCallsResponseBody) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) GetData() []*OralEvaluationStatisticsCallsCountResponse {
	return s.Data
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetData(v []*OralEvaluationStatisticsCallsCountResponse) *CountOralEvaluationStatisticsCallsResponseBody {
	s.Data = v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetErrCode(v string) *CountOralEvaluationStatisticsCallsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetErrMessage(v string) *CountOralEvaluationStatisticsCallsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetHttpStatusCode(v int32) *CountOralEvaluationStatisticsCallsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetRequestId(v string) *CountOralEvaluationStatisticsCallsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetSuccess(v bool) *CountOralEvaluationStatisticsCallsResponseBody {
	s.Success = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) Validate() error {
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
