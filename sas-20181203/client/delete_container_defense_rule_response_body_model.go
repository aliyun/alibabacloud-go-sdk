// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContainerDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteContainerDefenseRuleResponseBody
	GetCode() *string
	SetData(v int64) *DeleteContainerDefenseRuleResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *DeleteContainerDefenseRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteContainerDefenseRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContainerDefenseRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteContainerDefenseRuleResponseBody
	GetSuccess() *bool
}

type DeleteContainerDefenseRuleResponseBody struct {
	// The status code that is returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of deleted rules.
	//
	// example:
	//
	// 1
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
	// 47365EC5-**-6DEA1788EB11
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

func (s DeleteContainerDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContainerDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContainerDefenseRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteContainerDefenseRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *DeleteContainerDefenseRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteContainerDefenseRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContainerDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContainerDefenseRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteContainerDefenseRuleResponseBody) SetCode(v string) *DeleteContainerDefenseRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContainerDefenseRuleResponseBody) SetData(v int64) *DeleteContainerDefenseRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteContainerDefenseRuleResponseBody) SetHttpStatusCode(v int32) *DeleteContainerDefenseRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteContainerDefenseRuleResponseBody) SetMessage(v string) *DeleteContainerDefenseRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContainerDefenseRuleResponseBody) SetRequestId(v string) *DeleteContainerDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContainerDefenseRuleResponseBody) SetSuccess(v bool) *DeleteContainerDefenseRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteContainerDefenseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
