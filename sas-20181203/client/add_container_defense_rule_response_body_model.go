// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContainerDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddContainerDefenseRuleResponseBody
	GetCode() *string
	SetData(v int64) *AddContainerDefenseRuleResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *AddContainerDefenseRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddContainerDefenseRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddContainerDefenseRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddContainerDefenseRuleResponseBody
	GetSuccess() *bool
}

type AddContainerDefenseRuleResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the rule that is created.
	//
	// example:
	//
	// 182
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// There was an error with your request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8C376***AE74FB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddContainerDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddContainerDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddContainerDefenseRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddContainerDefenseRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddContainerDefenseRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddContainerDefenseRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddContainerDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddContainerDefenseRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddContainerDefenseRuleResponseBody) SetCode(v string) *AddContainerDefenseRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddContainerDefenseRuleResponseBody) SetData(v int64) *AddContainerDefenseRuleResponseBody {
	s.Data = &v
	return s
}

func (s *AddContainerDefenseRuleResponseBody) SetHttpStatusCode(v int32) *AddContainerDefenseRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddContainerDefenseRuleResponseBody) SetMessage(v string) *AddContainerDefenseRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddContainerDefenseRuleResponseBody) SetRequestId(v string) *AddContainerDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddContainerDefenseRuleResponseBody) SetSuccess(v bool) *AddContainerDefenseRuleResponseBody {
	s.Success = &v
	return s
}

func (s *AddContainerDefenseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
