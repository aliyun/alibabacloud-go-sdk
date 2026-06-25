// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGreyTagRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteGreyTagRouteResponseBody
	GetCode() *string
	SetData(v *DeleteGreyTagRouteResponseBodyData) *DeleteGreyTagRouteResponseBody
	GetData() *DeleteGreyTagRouteResponseBodyData
	SetErrorCode(v string) *DeleteGreyTagRouteResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteGreyTagRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGreyTagRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGreyTagRouteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteGreyTagRouteResponseBody
	GetTraceId() *string
}

type DeleteGreyTagRouteResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The request is successful.
	//
	// - **3xx**: The request is redirected.
	//
	// - **4xx**: A request error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the grey tag route.
	Data *DeleteGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. The following describes when this field is returned:
	//
	// - This parameter is not returned if the request is successful.
	//
	// - This parameter is returned if the request fails. For more information, see the **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The response message. Valid values:
	//
	// - **success*	- is returned if the request is successful.
	//
	// - An error message is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9D29CBD0-45D3-410B-9826-52F86F90****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID. You can use this ID to query call details.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteGreyTagRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGreyTagRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGreyTagRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGreyTagRouteResponseBody) GetData() *DeleteGreyTagRouteResponseBodyData {
	return s.Data
}

func (s *DeleteGreyTagRouteResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteGreyTagRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGreyTagRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGreyTagRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGreyTagRouteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteGreyTagRouteResponseBody) SetCode(v string) *DeleteGreyTagRouteResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetData(v *DeleteGreyTagRouteResponseBodyData) *DeleteGreyTagRouteResponseBody {
	s.Data = v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetErrorCode(v string) *DeleteGreyTagRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetMessage(v string) *DeleteGreyTagRouteResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetRequestId(v string) *DeleteGreyTagRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetSuccess(v bool) *DeleteGreyTagRouteResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetTraceId(v string) *DeleteGreyTagRouteResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteGreyTagRouteResponseBodyData struct {
	// The globally unique ID of the grey tag route.
	//
	// example:
	//
	// 16
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
}

func (s DeleteGreyTagRouteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteGreyTagRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteGreyTagRouteResponseBodyData) GetGreyTagRouteId() *int64 {
	return s.GreyTagRouteId
}

func (s *DeleteGreyTagRouteResponseBodyData) SetGreyTagRouteId(v int64) *DeleteGreyTagRouteResponseBodyData {
	s.GreyTagRouteId = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
