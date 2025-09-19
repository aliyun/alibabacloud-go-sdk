// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpaStrategyNewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateOpaStrategyNewResponseBody
	GetCode() *string
	SetData(v []*string) *UpdateOpaStrategyNewResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *UpdateOpaStrategyNewResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateOpaStrategyNewResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateOpaStrategyNewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateOpaStrategyNewResponseBody
	GetSuccess() *bool
}

type UpdateOpaStrategyNewResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of the clusters that failed to be updated.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F75B5FF5-DCB2-59CE-8978-08510707A9E6
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

func (s UpdateOpaStrategyNewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateOpaStrategyNewResponseBody) GetData() []*string {
	return s.Data
}

func (s *UpdateOpaStrategyNewResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateOpaStrategyNewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateOpaStrategyNewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOpaStrategyNewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateOpaStrategyNewResponseBody) SetCode(v string) *UpdateOpaStrategyNewResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateOpaStrategyNewResponseBody) SetData(v []*string) *UpdateOpaStrategyNewResponseBody {
	s.Data = v
	return s
}

func (s *UpdateOpaStrategyNewResponseBody) SetHttpStatusCode(v int32) *UpdateOpaStrategyNewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateOpaStrategyNewResponseBody) SetMessage(v string) *UpdateOpaStrategyNewResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateOpaStrategyNewResponseBody) SetRequestId(v string) *UpdateOpaStrategyNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOpaStrategyNewResponseBody) SetSuccess(v bool) *UpdateOpaStrategyNewResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateOpaStrategyNewResponseBody) Validate() error {
	return dara.Validate(s)
}
