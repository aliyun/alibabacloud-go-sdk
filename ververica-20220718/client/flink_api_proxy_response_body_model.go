// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlinkApiProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *FlinkApiProxyResponseBody
	GetData() *string
	SetErrorCode(v string) *FlinkApiProxyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *FlinkApiProxyResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *FlinkApiProxyResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *FlinkApiProxyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlinkApiProxyResponseBody
	GetSuccess() *bool
}

type FlinkApiProxyResponseBody struct {
	// 	- If the value of success was true, the result of the proxy request was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	//
	// example:
	//
	// { "jobs": [ { "jid": "4df35f8e54554b23bf7dcd38a151****", "name": "69d001d5-419a-4bfc-9c2e-849cacd3****", "state": "RUNNING", "start-time": 1659154942068, "end-time": -1, "duration": 188161756, "last-modification": 1659154968305, "tasks": { "total": 2, "created": 0, "scheduled": 0, "deploying": 0, "running": 2, "finished": 0, "canceling": 0, "canceled": 0, "failed": 0, "reconciling": 0, "initializing": 0 } } ] }
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FlinkApiProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlinkApiProxyResponseBody) GoString() string {
	return s.String()
}

func (s *FlinkApiProxyResponseBody) GetData() *string {
	return s.Data
}

func (s *FlinkApiProxyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *FlinkApiProxyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *FlinkApiProxyResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *FlinkApiProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlinkApiProxyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlinkApiProxyResponseBody) SetData(v string) *FlinkApiProxyResponseBody {
	s.Data = &v
	return s
}

func (s *FlinkApiProxyResponseBody) SetErrorCode(v string) *FlinkApiProxyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FlinkApiProxyResponseBody) SetErrorMessage(v string) *FlinkApiProxyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *FlinkApiProxyResponseBody) SetHttpCode(v int32) *FlinkApiProxyResponseBody {
	s.HttpCode = &v
	return s
}

func (s *FlinkApiProxyResponseBody) SetRequestId(v string) *FlinkApiProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlinkApiProxyResponseBody) SetSuccess(v bool) *FlinkApiProxyResponseBody {
	s.Success = &v
	return s
}

func (s *FlinkApiProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
