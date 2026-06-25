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
	// API status or POP error code. Details are as follows:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned result.
	Data *UpdateIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error code. Details are as follows:
	//
	// - If the request is successful, the **ErrorCode*	- field is not returned.
	//
	// - If the request failed, the **ErrorCode*	- field is returned. For more information, see the **Error Codes*	- list in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information. Details are as follows:
	//
	// - If the request is normal, **success*	- is returned.
	//
	// - If the request is abnormal, a specific abnormal error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the Ingress instance configuration was successfully updated. Details are as follows:
	//
	// - **true**: The update was successful.
	//
	// - **false**: The update failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Call chain ID.
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
	// Routing rule ID.
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
