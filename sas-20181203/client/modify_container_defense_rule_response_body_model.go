// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyContainerDefenseRuleResponseBody
	GetCode() *string
	SetData(v int64) *ModifyContainerDefenseRuleResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *ModifyContainerDefenseRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyContainerDefenseRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyContainerDefenseRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyContainerDefenseRuleResponseBody
	GetSuccess() *bool
}

type ModifyContainerDefenseRuleResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the rule that was modified.
	//
	// example:
	//
	// 100
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code. The value 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 768DFBE5-*A5DC35**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyContainerDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyContainerDefenseRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyContainerDefenseRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifyContainerDefenseRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyContainerDefenseRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyContainerDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyContainerDefenseRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyContainerDefenseRuleResponseBody) SetCode(v string) *ModifyContainerDefenseRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyContainerDefenseRuleResponseBody) SetData(v int64) *ModifyContainerDefenseRuleResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyContainerDefenseRuleResponseBody) SetHttpStatusCode(v int32) *ModifyContainerDefenseRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyContainerDefenseRuleResponseBody) SetMessage(v string) *ModifyContainerDefenseRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyContainerDefenseRuleResponseBody) SetRequestId(v string) *ModifyContainerDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyContainerDefenseRuleResponseBody) SetSuccess(v bool) *ModifyContainerDefenseRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyContainerDefenseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
