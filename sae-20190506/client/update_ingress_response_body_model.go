// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIngressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateIngressResponseBody
	GetCode() *string
	SetData(v *UpdateIngressResponseBodyData) *UpdateIngressResponseBody
	GetData() *UpdateIngressResponseBodyData
	SetErrorCode(v string) *UpdateIngressResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateIngressResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateIngressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateIngressResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateIngressResponseBody
	GetTraceId() *string
}

type UpdateIngressResponseBody struct {
	// The HTTP status code or the error code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Responses.
	Data *UpdateIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code. Value values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see the **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information. Valid values:
	//
	// 	- The error message returned because the request is normal and **success*	- is returned.
	//
	// 	- If the request is abnormal, the specific exception error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the configuration of the Ingress instance is updated. Valid values:
	//
	// 	- **true**: The update was successful.
	//
	// 	- **false**: Update failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateIngressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIngressResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIngressResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateIngressResponseBody) GetData() *UpdateIngressResponseBodyData {
	return s.Data
}

func (s *UpdateIngressResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateIngressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateIngressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIngressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateIngressResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateIngressResponseBody) SetCode(v string) *UpdateIngressResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateIngressResponseBody) SetData(v *UpdateIngressResponseBodyData) *UpdateIngressResponseBody {
	s.Data = v
	return s
}

func (s *UpdateIngressResponseBody) SetErrorCode(v string) *UpdateIngressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateIngressResponseBody) SetMessage(v string) *UpdateIngressResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateIngressResponseBody) SetRequestId(v string) *UpdateIngressResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIngressResponseBody) SetSuccess(v bool) *UpdateIngressResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateIngressResponseBody) SetTraceId(v string) *UpdateIngressResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateIngressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIngressResponseBodyData struct {
	// The ID of the routing rule.
	//
	// example:
	//
	// 87
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
}

func (s UpdateIngressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateIngressResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateIngressResponseBodyData) GetIngressId() *int64 {
	return s.IngressId
}

func (s *UpdateIngressResponseBodyData) SetIngressId(v int64) *UpdateIngressResponseBodyData {
	s.IngressId = &v
	return s
}

func (s *UpdateIngressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
