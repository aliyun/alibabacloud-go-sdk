// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIngressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteIngressResponseBody
	GetCode() *string
	SetData(v *DeleteIngressResponseBodyData) *DeleteIngressResponseBody
	GetData() *DeleteIngressResponseBodyData
	SetErrorCode(v string) *DeleteIngressResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteIngressResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIngressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteIngressResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteIngressResponseBody
	GetTraceId() *string
}

type DeleteIngressResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *DeleteIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the routing rule was deleted. Valid values:
	//
	// 	- **true**: The routing rule was deleted.
	//
	// 	- **false**: The routing rule failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteIngressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIngressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIngressResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteIngressResponseBody) GetData() *DeleteIngressResponseBodyData {
	return s.Data
}

func (s *DeleteIngressResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteIngressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIngressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIngressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIngressResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteIngressResponseBody) SetCode(v string) *DeleteIngressResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIngressResponseBody) SetData(v *DeleteIngressResponseBodyData) *DeleteIngressResponseBody {
	s.Data = v
	return s
}

func (s *DeleteIngressResponseBody) SetErrorCode(v string) *DeleteIngressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteIngressResponseBody) SetMessage(v string) *DeleteIngressResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIngressResponseBody) SetRequestId(v string) *DeleteIngressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIngressResponseBody) SetSuccess(v bool) *DeleteIngressResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIngressResponseBody) SetTraceId(v string) *DeleteIngressResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteIngressResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteIngressResponseBodyData struct {
	// The ID of the deleted routing rule.
	//
	// example:
	//
	// 87
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
}

func (s DeleteIngressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteIngressResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteIngressResponseBodyData) GetIngressId() *int64 {
	return s.IngressId
}

func (s *DeleteIngressResponseBodyData) SetIngressId(v int64) *DeleteIngressResponseBodyData {
	s.IngressId = &v
	return s
}

func (s *DeleteIngressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
