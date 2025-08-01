// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateClusterSpecResponseBody
	GetCode() *int32
	SetData(v string) *UpdateClusterSpecResponseBody
	GetData() *string
	SetErrorCode(v string) *UpdateClusterSpecResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *UpdateClusterSpecResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateClusterSpecResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateClusterSpecResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateClusterSpecResponseBody
	GetSuccess() *bool
}

type UpdateClusterSpecResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- The **ErrorCode*	- parameter is returned if the request fails. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5B170A0D-2C5D-4CF8-B808-03966B86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClusterSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterSpecResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateClusterSpecResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateClusterSpecResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateClusterSpecResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateClusterSpecResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateClusterSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterSpecResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateClusterSpecResponseBody) SetCode(v int32) *UpdateClusterSpecResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateClusterSpecResponseBody) SetData(v string) *UpdateClusterSpecResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateClusterSpecResponseBody) SetErrorCode(v string) *UpdateClusterSpecResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateClusterSpecResponseBody) SetHttpStatusCode(v int32) *UpdateClusterSpecResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateClusterSpecResponseBody) SetMessage(v string) *UpdateClusterSpecResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateClusterSpecResponseBody) SetRequestId(v string) *UpdateClusterSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterSpecResponseBody) SetSuccess(v bool) *UpdateClusterSpecResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateClusterSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
