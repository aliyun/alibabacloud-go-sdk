// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHoneypotResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteHoneypotResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteHoneypotResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHoneypotResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHoneypotResponseBody
	GetSuccess() *bool
}

type DeleteHoneypotResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// E14721CB-B32F-56F2-8490-CDA18E4F9268
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHoneypotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHoneypotResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteHoneypotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHoneypotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHoneypotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHoneypotResponseBody) SetCode(v string) *DeleteHoneypotResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHoneypotResponseBody) SetHttpStatusCode(v int32) *DeleteHoneypotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteHoneypotResponseBody) SetMessage(v string) *DeleteHoneypotResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHoneypotResponseBody) SetRequestId(v string) *DeleteHoneypotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHoneypotResponseBody) SetSuccess(v bool) *DeleteHoneypotResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHoneypotResponseBody) Validate() error {
	return dara.Validate(s)
}
