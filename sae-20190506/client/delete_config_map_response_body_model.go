// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteConfigMapResponseBody
	GetCode() *string
	SetData(v *DeleteConfigMapResponseBodyData) *DeleteConfigMapResponseBody
	GetData() *DeleteConfigMapResponseBodyData
	SetErrorCode(v string) *DeleteConfigMapResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteConfigMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteConfigMapResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteConfigMapResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteConfigMapResponseBody
	GetTraceId() *string
}

type DeleteConfigMapResponseBody struct {
	// The HTTP status code returned for the request.
	//
	// - **2xx**: success
	//
	// - **3xx**: redirection
	//
	// - **4xx**: client error
	//
	// - **5xx**: server error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The operation result.
	Data *DeleteConfigMapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - Not returned if the request is successful.
	//
	// - Returned if the request fails. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The response message.
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, an error message is returned.
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
	// Indicates whether the deletion succeeded. Valid values:
	//
	// - **true**: The deletion was successful.
	//
	// - **false**: The deletion failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID for querying request details.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteConfigMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigMapResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteConfigMapResponseBody) GetData() *DeleteConfigMapResponseBodyData {
	return s.Data
}

func (s *DeleteConfigMapResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteConfigMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConfigMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConfigMapResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteConfigMapResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteConfigMapResponseBody) SetCode(v string) *DeleteConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConfigMapResponseBody) SetData(v *DeleteConfigMapResponseBodyData) *DeleteConfigMapResponseBody {
	s.Data = v
	return s
}

func (s *DeleteConfigMapResponseBody) SetErrorCode(v string) *DeleteConfigMapResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteConfigMapResponseBody) SetMessage(v string) *DeleteConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConfigMapResponseBody) SetRequestId(v string) *DeleteConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigMapResponseBody) SetSuccess(v bool) *DeleteConfigMapResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConfigMapResponseBody) SetTraceId(v string) *DeleteConfigMapResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteConfigMapResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteConfigMapResponseBodyData struct {
	// The ID of the deleted ConfigMap instance.
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
}

func (s DeleteConfigMapResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigMapResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteConfigMapResponseBodyData) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *DeleteConfigMapResponseBodyData) SetConfigMapId(v int64) *DeleteConfigMapResponseBodyData {
	s.ConfigMapId = &v
	return s
}

func (s *DeleteConfigMapResponseBodyData) Validate() error {
	return dara.Validate(s)
}
