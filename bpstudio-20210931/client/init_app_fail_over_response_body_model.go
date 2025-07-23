// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitAppFailOverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitAppFailOverResponseBody
	GetCode() *string
	SetData(v int32) *InitAppFailOverResponseBody
	GetData() *int32
	SetMessage(v string) *InitAppFailOverResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitAppFailOverResponseBody
	GetRequestId() *string
}

type InitAppFailOverResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The switchover task ID.
	//
	// example:
	//
	// 7250
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// Unsupported Operation PrepareEvent->FailOverPrepareSuccess FoApp_DDLJK2WM8ETU9JAC
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitAppFailOverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitAppFailOverResponseBody) GoString() string {
	return s.String()
}

func (s *InitAppFailOverResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitAppFailOverResponseBody) GetData() *int32 {
	return s.Data
}

func (s *InitAppFailOverResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitAppFailOverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitAppFailOverResponseBody) SetCode(v string) *InitAppFailOverResponseBody {
	s.Code = &v
	return s
}

func (s *InitAppFailOverResponseBody) SetData(v int32) *InitAppFailOverResponseBody {
	s.Data = &v
	return s
}

func (s *InitAppFailOverResponseBody) SetMessage(v string) *InitAppFailOverResponseBody {
	s.Message = &v
	return s
}

func (s *InitAppFailOverResponseBody) SetRequestId(v string) *InitAppFailOverResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitAppFailOverResponseBody) Validate() error {
	return dara.Validate(s)
}
