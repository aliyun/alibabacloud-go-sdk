// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishHotlineServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FinishHotlineServiceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *FinishHotlineServiceResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *FinishHotlineServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *FinishHotlineServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FinishHotlineServiceResponseBody
	GetSuccess() *bool
}

type FinishHotlineServiceResponseBody struct {
	// The status code. A value of Success indicates that the request was successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FinishHotlineServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FinishHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *FinishHotlineServiceResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *FinishHotlineServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FinishHotlineServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FinishHotlineServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FinishHotlineServiceResponseBody) SetCode(v string) *FinishHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetHttpStatusCode(v int64) *FinishHotlineServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetMessage(v string) *FinishHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetRequestId(v string) *FinishHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetSuccess(v bool) *FinishHotlineServiceResponseBody {
	s.Success = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
