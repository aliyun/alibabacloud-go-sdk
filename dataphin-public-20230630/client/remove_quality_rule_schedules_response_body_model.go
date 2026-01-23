// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveQualityRuleSchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveQualityRuleSchedulesResponseBody
	GetCode() *string
	SetData(v int32) *RemoveQualityRuleSchedulesResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *RemoveQualityRuleSchedulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveQualityRuleSchedulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveQualityRuleSchedulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveQualityRuleSchedulesResponseBody
	GetSuccess() *bool
}

type RemoveQualityRuleSchedulesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s RemoveQualityRuleSchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveQualityRuleSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveQualityRuleSchedulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveQualityRuleSchedulesResponseBody) GetData() *int32 {
	return s.Data
}

func (s *RemoveQualityRuleSchedulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveQualityRuleSchedulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveQualityRuleSchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveQualityRuleSchedulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveQualityRuleSchedulesResponseBody) SetCode(v string) *RemoveQualityRuleSchedulesResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveQualityRuleSchedulesResponseBody) SetData(v int32) *RemoveQualityRuleSchedulesResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveQualityRuleSchedulesResponseBody) SetHttpStatusCode(v int32) *RemoveQualityRuleSchedulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveQualityRuleSchedulesResponseBody) SetMessage(v string) *RemoveQualityRuleSchedulesResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveQualityRuleSchedulesResponseBody) SetRequestId(v string) *RemoveQualityRuleSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveQualityRuleSchedulesResponseBody) SetSuccess(v bool) *RemoveQualityRuleSchedulesResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveQualityRuleSchedulesResponseBody) Validate() error {
	return dara.Validate(s)
}
