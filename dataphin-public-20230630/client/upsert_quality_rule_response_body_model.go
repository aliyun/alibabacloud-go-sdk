// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpsertQualityRuleResponseBody
	GetCode() *string
	SetData(v int64) *UpsertQualityRuleResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpsertQualityRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpsertQualityRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpsertQualityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpsertQualityRuleResponseBody
	GetSuccess() *bool
}

type UpsertQualityRuleResponseBody struct {
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

func (s UpsertQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpsertQualityRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpsertQualityRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpsertQualityRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpsertQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertQualityRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpsertQualityRuleResponseBody) SetCode(v string) *UpsertQualityRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpsertQualityRuleResponseBody) SetData(v int64) *UpsertQualityRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpsertQualityRuleResponseBody) SetHttpStatusCode(v int32) *UpsertQualityRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpsertQualityRuleResponseBody) SetMessage(v string) *UpsertQualityRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpsertQualityRuleResponseBody) SetRequestId(v string) *UpsertQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertQualityRuleResponseBody) SetSuccess(v bool) *UpsertQualityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpsertQualityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
