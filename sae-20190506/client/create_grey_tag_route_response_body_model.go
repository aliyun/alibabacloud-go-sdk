// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGreyTagRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGreyTagRouteResponseBody
	GetCode() *string
	SetData(v *CreateGreyTagRouteResponseBodyData) *CreateGreyTagRouteResponseBody
	GetData() *CreateGreyTagRouteResponseBodyData
	SetErrorCode(v string) *CreateGreyTagRouteResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateGreyTagRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGreyTagRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGreyTagRouteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateGreyTagRouteResponseBody
	GetTraceId() *string
}

type CreateGreyTagRouteResponseBody struct {
	// The HTTP status code. Valid values:
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
	// The information about the canary release rule.
	Data *CreateGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- **ErrorCode*	- is not returned if a request is successful.
	//
	// 	- **ErrorCode*	- is returned if a request failed. For more information, see **Error code*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned for the operation.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9D29CBD0-45D3-410B-9826-52F86F90****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the information of the change order was queried. Valid values:
	//
	// 	- **true**: The information was queried.
	//
	// 	- **false**: The information failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateGreyTagRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGreyTagRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGreyTagRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGreyTagRouteResponseBody) GetData() *CreateGreyTagRouteResponseBodyData {
	return s.Data
}

func (s *CreateGreyTagRouteResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateGreyTagRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGreyTagRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGreyTagRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGreyTagRouteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateGreyTagRouteResponseBody) SetCode(v string) *CreateGreyTagRouteResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetData(v *CreateGreyTagRouteResponseBodyData) *CreateGreyTagRouteResponseBody {
	s.Data = v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetErrorCode(v string) *CreateGreyTagRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetMessage(v string) *CreateGreyTagRouteResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetRequestId(v string) *CreateGreyTagRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetSuccess(v bool) *CreateGreyTagRouteResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetTraceId(v string) *CreateGreyTagRouteResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGreyTagRouteResponseBodyData struct {
	// The ID of the canary release rule. The ID is globally unique.
	//
	// example:
	//
	// 16
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
}

func (s CreateGreyTagRouteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateGreyTagRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGreyTagRouteResponseBodyData) GetGreyTagRouteId() *int64 {
	return s.GreyTagRouteId
}

func (s *CreateGreyTagRouteResponseBodyData) SetGreyTagRouteId(v int64) *CreateGreyTagRouteResponseBodyData {
	s.GreyTagRouteId = &v
	return s
}

func (s *CreateGreyTagRouteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
