// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTraceAppConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *SaveTraceAppConfigResponseBody
	GetCode() *int64
	SetData(v string) *SaveTraceAppConfigResponseBody
	GetData() *string
	SetMessage(v string) *SaveTraceAppConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveTraceAppConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveTraceAppConfigResponseBody
	GetSuccess() *bool
}

type SaveTraceAppConfigResponseBody struct {
	// The HTTP status code. 2XX indicates that the request was successful. 3XX indicates that the request was redirected. 4XX indicates that a request error occurred. 5XX indicates that a server error occurred.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveTraceAppConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTraceAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SaveTraceAppConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *SaveTraceAppConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveTraceAppConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTraceAppConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveTraceAppConfigResponseBody) SetCode(v int64) *SaveTraceAppConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SaveTraceAppConfigResponseBody) SetData(v string) *SaveTraceAppConfigResponseBody {
	s.Data = &v
	return s
}

func (s *SaveTraceAppConfigResponseBody) SetMessage(v string) *SaveTraceAppConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SaveTraceAppConfigResponseBody) SetRequestId(v string) *SaveTraceAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTraceAppConfigResponseBody) SetSuccess(v bool) *SaveTraceAppConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SaveTraceAppConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
