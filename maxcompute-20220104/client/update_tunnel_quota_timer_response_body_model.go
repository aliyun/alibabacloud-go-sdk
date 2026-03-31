// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTunnelQuotaTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateTunnelQuotaTimerResponseBody
	GetData() *string
	SetErrorCode(v string) *UpdateTunnelQuotaTimerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateTunnelQuotaTimerResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *UpdateTunnelQuotaTimerResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateTunnelQuotaTimerResponseBody
	GetRequestId() *string
}

type UpdateTunnelQuotaTimerResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc12e4316675560945192024e1044
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateTunnelQuotaTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTunnelQuotaTimerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTunnelQuotaTimerResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateTunnelQuotaTimerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTunnelQuotaTimerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateTunnelQuotaTimerResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateTunnelQuotaTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTunnelQuotaTimerResponseBody) SetData(v string) *UpdateTunnelQuotaTimerResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponseBody) SetErrorCode(v string) *UpdateTunnelQuotaTimerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponseBody) SetErrorMsg(v string) *UpdateTunnelQuotaTimerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponseBody) SetHttpCode(v int32) *UpdateTunnelQuotaTimerResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponseBody) SetRequestId(v string) *UpdateTunnelQuotaTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponseBody) Validate() error {
	return dara.Validate(s)
}
