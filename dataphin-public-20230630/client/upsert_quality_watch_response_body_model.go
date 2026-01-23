// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityWatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpsertQualityWatchResponseBody
	GetCode() *string
	SetData(v int64) *UpsertQualityWatchResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpsertQualityWatchResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpsertQualityWatchResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpsertQualityWatchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpsertQualityWatchResponseBody
	GetSuccess() *bool
}

type UpsertQualityWatchResponseBody struct {
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

func (s UpsertQualityWatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpsertQualityWatchResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpsertQualityWatchResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpsertQualityWatchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpsertQualityWatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertQualityWatchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpsertQualityWatchResponseBody) SetCode(v string) *UpsertQualityWatchResponseBody {
	s.Code = &v
	return s
}

func (s *UpsertQualityWatchResponseBody) SetData(v int64) *UpsertQualityWatchResponseBody {
	s.Data = &v
	return s
}

func (s *UpsertQualityWatchResponseBody) SetHttpStatusCode(v int32) *UpsertQualityWatchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpsertQualityWatchResponseBody) SetMessage(v string) *UpsertQualityWatchResponseBody {
	s.Message = &v
	return s
}

func (s *UpsertQualityWatchResponseBody) SetRequestId(v string) *UpsertQualityWatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertQualityWatchResponseBody) SetSuccess(v bool) *UpsertQualityWatchResponseBody {
	s.Success = &v
	return s
}

func (s *UpsertQualityWatchResponseBody) Validate() error {
	return dara.Validate(s)
}
