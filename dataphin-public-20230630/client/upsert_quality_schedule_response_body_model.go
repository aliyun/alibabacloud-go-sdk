// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpsertQualityScheduleResponseBody
	GetCode() *string
	SetData(v int64) *UpsertQualityScheduleResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpsertQualityScheduleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpsertQualityScheduleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpsertQualityScheduleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpsertQualityScheduleResponseBody
	GetSuccess() *bool
}

type UpsertQualityScheduleResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpsertQualityScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertQualityScheduleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpsertQualityScheduleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpsertQualityScheduleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpsertQualityScheduleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpsertQualityScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertQualityScheduleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpsertQualityScheduleResponseBody) SetCode(v string) *UpsertQualityScheduleResponseBody {
	s.Code = &v
	return s
}

func (s *UpsertQualityScheduleResponseBody) SetData(v int64) *UpsertQualityScheduleResponseBody {
	s.Data = &v
	return s
}

func (s *UpsertQualityScheduleResponseBody) SetHttpStatusCode(v int32) *UpsertQualityScheduleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpsertQualityScheduleResponseBody) SetMessage(v string) *UpsertQualityScheduleResponseBody {
	s.Message = &v
	return s
}

func (s *UpsertQualityScheduleResponseBody) SetRequestId(v string) *UpsertQualityScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertQualityScheduleResponseBody) SetSuccess(v bool) *UpsertQualityScheduleResponseBody {
	s.Success = &v
	return s
}

func (s *UpsertQualityScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
