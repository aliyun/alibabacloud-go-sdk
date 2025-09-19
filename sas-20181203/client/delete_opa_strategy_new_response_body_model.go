// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpaStrategyNewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteOpaStrategyNewResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteOpaStrategyNewResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteOpaStrategyNewResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteOpaStrategyNewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteOpaStrategyNewResponseBody
	GetSuccess() *bool
}

type DeleteOpaStrategyNewResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 8FD58F11-0F4D-5C7F-B9B2-CFD76108F9A4
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

func (s DeleteOpaStrategyNewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpaStrategyNewResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOpaStrategyNewResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteOpaStrategyNewResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteOpaStrategyNewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteOpaStrategyNewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOpaStrategyNewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteOpaStrategyNewResponseBody) SetCode(v string) *DeleteOpaStrategyNewResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOpaStrategyNewResponseBody) SetHttpStatusCode(v int32) *DeleteOpaStrategyNewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteOpaStrategyNewResponseBody) SetMessage(v string) *DeleteOpaStrategyNewResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOpaStrategyNewResponseBody) SetRequestId(v string) *DeleteOpaStrategyNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOpaStrategyNewResponseBody) SetSuccess(v bool) *DeleteOpaStrategyNewResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOpaStrategyNewResponseBody) Validate() error {
	return dara.Validate(s)
}
