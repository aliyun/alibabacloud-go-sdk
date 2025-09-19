// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerDefenseRuleSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyContainerDefenseRuleSwitchResponseBody
	GetCode() *string
	SetData(v int64) *ModifyContainerDefenseRuleSwitchResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *ModifyContainerDefenseRuleSwitchResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyContainerDefenseRuleSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyContainerDefenseRuleSwitchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyContainerDefenseRuleSwitchResponseBody
	GetSuccess() *bool
}

type ModifyContainerDefenseRuleSwitchResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 30
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C8A137FB-6E18-5741-9B47-D9A0EBA3495F
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

func (s ModifyContainerDefenseRuleSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerDefenseRuleSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) SetCode(v string) *ModifyContainerDefenseRuleSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) SetData(v int64) *ModifyContainerDefenseRuleSwitchResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) SetHttpStatusCode(v int32) *ModifyContainerDefenseRuleSwitchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) SetMessage(v string) *ModifyContainerDefenseRuleSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) SetRequestId(v string) *ModifyContainerDefenseRuleSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) SetSuccess(v bool) *ModifyContainerDefenseRuleSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
