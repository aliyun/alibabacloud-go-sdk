// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelMetricsDTO) *ModelRouterQueryObservationMetricsResponseBody
	GetData() *ModelMetricsDTO
	SetErrCode(v string) *ModelRouterQueryObservationMetricsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryObservationMetricsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryObservationMetricsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryObservationMetricsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryObservationMetricsResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryObservationMetricsResponseBody struct {
	// example:
	//
	// []
	Data *ModelMetricsDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterQueryObservationMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetData() *ModelMetricsDTO {
	return s.Data
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetData(v *ModelMetricsDTO) *ModelRouterQueryObservationMetricsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetErrCode(v string) *ModelRouterQueryObservationMetricsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetErrMessage(v string) *ModelRouterQueryObservationMetricsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryObservationMetricsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetRequestId(v string) *ModelRouterQueryObservationMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetSuccess(v bool) *ModelRouterQueryObservationMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
