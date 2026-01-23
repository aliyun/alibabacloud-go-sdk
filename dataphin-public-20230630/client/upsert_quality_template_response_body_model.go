// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpsertQualityTemplateResponseBody
	GetCode() *string
	SetData(v int64) *UpsertQualityTemplateResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpsertQualityTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpsertQualityTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpsertQualityTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpsertQualityTemplateResponseBody
	GetSuccess() *bool
}

type UpsertQualityTemplateResponseBody struct {
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

func (s UpsertQualityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertQualityTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpsertQualityTemplateResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpsertQualityTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpsertQualityTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpsertQualityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertQualityTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpsertQualityTemplateResponseBody) SetCode(v string) *UpsertQualityTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpsertQualityTemplateResponseBody) SetData(v int64) *UpsertQualityTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *UpsertQualityTemplateResponseBody) SetHttpStatusCode(v int32) *UpsertQualityTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpsertQualityTemplateResponseBody) SetMessage(v string) *UpsertQualityTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpsertQualityTemplateResponseBody) SetRequestId(v string) *UpsertQualityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertQualityTemplateResponseBody) SetSuccess(v bool) *UpsertQualityTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpsertQualityTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
