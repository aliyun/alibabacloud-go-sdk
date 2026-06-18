// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartHotlineServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartHotlineServiceResponseBody
	GetCode() *string
	SetData(v string) *StartHotlineServiceResponseBody
	GetData() *string
	SetHttpStatusCode(v int64) *StartHotlineServiceResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *StartHotlineServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartHotlineServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartHotlineServiceResponseBody
	GetSuccess() *bool
}

type StartHotlineServiceResponseBody struct {
	// Status code. A value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Token required to initiate a heartbeat, returned after a successful request.
	//
	// example:
	//
	// 0079e7a845e373****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartHotlineServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartHotlineServiceResponseBody) GetData() *string {
	return s.Data
}

func (s *StartHotlineServiceResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *StartHotlineServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartHotlineServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartHotlineServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartHotlineServiceResponseBody) SetCode(v string) *StartHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetData(v string) *StartHotlineServiceResponseBody {
	s.Data = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetHttpStatusCode(v int64) *StartHotlineServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetMessage(v string) *StartHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetRequestId(v string) *StartHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetSuccess(v bool) *StartHotlineServiceResponseBody {
	s.Success = &v
	return s
}

func (s *StartHotlineServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
