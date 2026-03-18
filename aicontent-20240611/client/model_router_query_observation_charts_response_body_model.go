// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationChartsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ObservationChartsDTO) *ModelRouterQueryObservationChartsResponseBody
	GetData() *ObservationChartsDTO
	SetErrCode(v string) *ModelRouterQueryObservationChartsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryObservationChartsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryObservationChartsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryObservationChartsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryObservationChartsResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryObservationChartsResponseBody struct {
	// example:
	//
	// []
	Data *ObservationChartsDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterQueryObservationChartsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationChartsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationChartsResponseBody) GetData() *ObservationChartsDTO {
	return s.Data
}

func (s *ModelRouterQueryObservationChartsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryObservationChartsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryObservationChartsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryObservationChartsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryObservationChartsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryObservationChartsResponseBody) SetData(v *ObservationChartsDTO) *ModelRouterQueryObservationChartsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryObservationChartsResponseBody) SetErrCode(v string) *ModelRouterQueryObservationChartsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryObservationChartsResponseBody) SetErrMessage(v string) *ModelRouterQueryObservationChartsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryObservationChartsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryObservationChartsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryObservationChartsResponseBody) SetRequestId(v string) *ModelRouterQueryObservationChartsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryObservationChartsResponseBody) SetSuccess(v bool) *ModelRouterQueryObservationChartsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryObservationChartsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
