// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageNewTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadMessageNewTotalResponseBody
	GetCode() *string
	SetData(v int64) *ReadMessageNewTotalResponseBody
	GetData() *int64
	SetMessage(v string) *ReadMessageNewTotalResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadMessageNewTotalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadMessageNewTotalResponseBody
	GetSuccess() *bool
}

type ReadMessageNewTotalResponseBody struct {
	// The error code returned when the call failed. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of new messages.
	//
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned when the call failed.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMessageNewTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageNewTotalResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageNewTotalResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadMessageNewTotalResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ReadMessageNewTotalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadMessageNewTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadMessageNewTotalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadMessageNewTotalResponseBody) SetCode(v string) *ReadMessageNewTotalResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetData(v int64) *ReadMessageNewTotalResponseBody {
	s.Data = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetMessage(v string) *ReadMessageNewTotalResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetRequestId(v string) *ReadMessageNewTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetSuccess(v bool) *ReadMessageNewTotalResponseBody {
	s.Success = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) Validate() error {
	return dara.Validate(s)
}
