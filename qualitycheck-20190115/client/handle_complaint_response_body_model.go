// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleComplaintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HandleComplaintResponseBody
	GetCode() *string
	SetData(v string) *HandleComplaintResponseBody
	GetData() *string
	SetMessage(v string) *HandleComplaintResponseBody
	GetMessage() *string
	SetRequestId(v string) *HandleComplaintResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HandleComplaintResponseBody
	GetSuccess() *bool
}

type HandleComplaintResponseBody struct {
	// The result code. A value of **200*	- indicates that the request was successful. Other values indicate that the request failed. You can use this code to identify the cause of the error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	//
	// example:
	//
	// 无
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message. If the request is successful, **successful*	- is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. A value of **true*	- indicates success, and **false or null*	- indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HandleComplaintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HandleComplaintResponseBody) GoString() string {
	return s.String()
}

func (s *HandleComplaintResponseBody) GetCode() *string {
	return s.Code
}

func (s *HandleComplaintResponseBody) GetData() *string {
	return s.Data
}

func (s *HandleComplaintResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HandleComplaintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HandleComplaintResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HandleComplaintResponseBody) SetCode(v string) *HandleComplaintResponseBody {
	s.Code = &v
	return s
}

func (s *HandleComplaintResponseBody) SetData(v string) *HandleComplaintResponseBody {
	s.Data = &v
	return s
}

func (s *HandleComplaintResponseBody) SetMessage(v string) *HandleComplaintResponseBody {
	s.Message = &v
	return s
}

func (s *HandleComplaintResponseBody) SetRequestId(v string) *HandleComplaintResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleComplaintResponseBody) SetSuccess(v bool) *HandleComplaintResponseBody {
	s.Success = &v
	return s
}

func (s *HandleComplaintResponseBody) Validate() error {
	return dara.Validate(s)
}
