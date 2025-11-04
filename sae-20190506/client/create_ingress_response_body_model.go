// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIngressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateIngressResponseBody
	GetCode() *string
	SetData(v *CreateIngressResponseBodyData) *CreateIngressResponseBody
	GetData() *CreateIngressResponseBodyData
	SetErrorCode(v string) *CreateIngressResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateIngressResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIngressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateIngressResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateIngressResponseBody
	GetTraceId() *string
}

type CreateIngressResponseBody struct {
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
	Data *CreateIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error code. Valid values:
	//
	// 	- If the request is successful, no **ErrorCode*	- fields are returned.
	//
	// 	- Request failed: **ErrorCode*	- fields are returned. For more information, see **Error codes**.
	//
	// example:
	//
	// success
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
	// Indicates whether the routing rule was created. Valid values:
	//
	// 	- **true**: The ConfigMap was created.
	//
	// 	- **false**: The ConfigMap failed to be created.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. This parameter is used to query the exact call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateIngressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIngressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIngressResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateIngressResponseBody) GetData() *CreateIngressResponseBodyData {
	return s.Data
}

func (s *CreateIngressResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateIngressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIngressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIngressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateIngressResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateIngressResponseBody) SetCode(v string) *CreateIngressResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIngressResponseBody) SetData(v *CreateIngressResponseBodyData) *CreateIngressResponseBody {
	s.Data = v
	return s
}

func (s *CreateIngressResponseBody) SetErrorCode(v string) *CreateIngressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateIngressResponseBody) SetMessage(v string) *CreateIngressResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIngressResponseBody) SetRequestId(v string) *CreateIngressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIngressResponseBody) SetSuccess(v bool) *CreateIngressResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIngressResponseBody) SetTraceId(v string) *CreateIngressResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateIngressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIngressResponseBodyData struct {
	// The ID of the routing rule.
	//
	// example:
	//
	// 87
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
}

func (s CreateIngressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateIngressResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIngressResponseBodyData) GetIngressId() *int64 {
	return s.IngressId
}

func (s *CreateIngressResponseBodyData) SetIngressId(v int64) *CreateIngressResponseBodyData {
	s.IngressId = &v
	return s
}

func (s *CreateIngressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
