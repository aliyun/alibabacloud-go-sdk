// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleTaskLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityRuleTaskLogResponseBody
	GetCode() *string
	SetData(v string) *GetQualityRuleTaskLogResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetQualityRuleTaskLogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityRuleTaskLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityRuleTaskLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityRuleTaskLogResponseBody
	GetSuccess() *bool
}

type GetQualityRuleTaskLogResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// test
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s GetQualityRuleTaskLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTaskLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTaskLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityRuleTaskLogResponseBody) GetData() *string {
	return s.Data
}

func (s *GetQualityRuleTaskLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityRuleTaskLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityRuleTaskLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityRuleTaskLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityRuleTaskLogResponseBody) SetCode(v string) *GetQualityRuleTaskLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityRuleTaskLogResponseBody) SetData(v string) *GetQualityRuleTaskLogResponseBody {
	s.Data = &v
	return s
}

func (s *GetQualityRuleTaskLogResponseBody) SetHttpStatusCode(v int32) *GetQualityRuleTaskLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityRuleTaskLogResponseBody) SetMessage(v string) *GetQualityRuleTaskLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityRuleTaskLogResponseBody) SetRequestId(v string) *GetQualityRuleTaskLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityRuleTaskLogResponseBody) SetSuccess(v bool) *GetQualityRuleTaskLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityRuleTaskLogResponseBody) Validate() error {
	return dara.Validate(s)
}
