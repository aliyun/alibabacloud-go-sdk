// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteApplicationResponseBody
	GetCode() *string
	SetData(v *DeleteApplicationResponseBodyData) *DeleteApplicationResponseBody
	GetData() *DeleteApplicationResponseBodyData
	SetErrorCode(v string) *DeleteApplicationResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteApplicationResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteApplicationResponseBody
	GetTraceId() *string
}

type DeleteApplicationResponseBody struct {
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
	// The result returned.
	Data *DeleteApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned. Valid values:
	//
	// 	- If the request was successful, **success*	- is returned.
	//
	// 	- If the request failed, an error code is returned.
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
	// Indicates whether the application is deleted. Valid values:
	//
	// 	- **true**: The applications were deleted
	//
	// 	- **false**: The applications failed to be deleted.
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

func (s DeleteApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteApplicationResponseBody) GetData() *DeleteApplicationResponseBodyData {
	return s.Data
}

func (s *DeleteApplicationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteApplicationResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteApplicationResponseBody) SetCode(v string) *DeleteApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetData(v *DeleteApplicationResponseBodyData) *DeleteApplicationResponseBody {
	s.Data = v
	return s
}

func (s *DeleteApplicationResponseBody) SetErrorCode(v string) *DeleteApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetMessage(v string) *DeleteApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetSuccess(v bool) *DeleteApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetTraceId(v string) *DeleteApplicationResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteApplicationResponseBodyData struct {
	// The ID of the change order. The ID can be used to query the status of the change task.
	//
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s DeleteApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DeleteApplicationResponseBodyData) SetChangeOrderId(v string) *DeleteApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *DeleteApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
