// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindNlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindNlbResponseBody
	GetCode() *string
	SetData(v *BindNlbResponseBodyData) *BindNlbResponseBody
	GetData() *BindNlbResponseBodyData
	SetErrorCode(v string) *BindNlbResponseBody
	GetErrorCode() *string
	SetMessage(v string) *BindNlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindNlbResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindNlbResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BindNlbResponseBody
	GetTraceId() *string
}

type BindNlbResponseBody struct {
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
	// The returned data.
	Data *BindNlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code. Value values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
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
	// Id of the request
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the applications were stopped. Valid values:
	//
	// 	- **true**: The applications were stopped.
	//
	// 	- **false**: The applications failed to be stopped.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// 0a06dfe717389800573793090e0589
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s BindNlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindNlbResponseBody) GoString() string {
	return s.String()
}

func (s *BindNlbResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindNlbResponseBody) GetData() *BindNlbResponseBodyData {
	return s.Data
}

func (s *BindNlbResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BindNlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindNlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindNlbResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindNlbResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BindNlbResponseBody) SetCode(v string) *BindNlbResponseBody {
	s.Code = &v
	return s
}

func (s *BindNlbResponseBody) SetData(v *BindNlbResponseBodyData) *BindNlbResponseBody {
	s.Data = v
	return s
}

func (s *BindNlbResponseBody) SetErrorCode(v string) *BindNlbResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BindNlbResponseBody) SetMessage(v string) *BindNlbResponseBody {
	s.Message = &v
	return s
}

func (s *BindNlbResponseBody) SetRequestId(v string) *BindNlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindNlbResponseBody) SetSuccess(v bool) *BindNlbResponseBody {
	s.Success = &v
	return s
}

func (s *BindNlbResponseBody) SetTraceId(v string) *BindNlbResponseBody {
	s.TraceId = &v
	return s
}

func (s *BindNlbResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindNlbResponseBodyData struct {
	// The ID of the change order. The ID can be used to query the status of the change task.
	//
	// example:
	//
	// ba386059-69b1-4e65-b1e5-0682d9fa****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s BindNlbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindNlbResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindNlbResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *BindNlbResponseBodyData) SetChangeOrderId(v string) *BindNlbResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *BindNlbResponseBodyData) Validate() error {
	return dara.Validate(s)
}
