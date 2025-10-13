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
	// The response.
	Data *CreateIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information that is returned. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
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
	// Indicates whether the Secret is successfully deleted. Valid values:
	//
	// 	- **true**: The instance was deleted.
	//
	// 	- **false**: The instance failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. It is used to query the details of a request.
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
