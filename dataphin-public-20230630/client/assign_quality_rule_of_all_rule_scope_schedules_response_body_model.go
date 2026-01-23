// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignQualityRuleOfAllRuleScopeSchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody
	GetCode() *string
	SetData(v bool) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody
	GetSuccess() *bool
}

type AssignQualityRuleOfAllRuleScopeSchedulesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) GetData() *bool {
	return s.Data
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) SetCode(v string) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody {
	s.Code = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) SetData(v bool) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody {
	s.Data = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) SetHttpStatusCode(v int32) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) SetMessage(v string) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody {
	s.Message = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) SetRequestId(v string) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) SetSuccess(v bool) *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody {
	s.Success = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) Validate() error {
	return dara.Validate(s)
}
