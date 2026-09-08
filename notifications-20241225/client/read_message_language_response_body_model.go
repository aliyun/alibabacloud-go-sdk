// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageLanguageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadMessageLanguageResponseBody
	GetCode() *string
	SetData(v string) *ReadMessageLanguageResponseBody
	GetData() *string
	SetMessage(v string) *ReadMessageLanguageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadMessageLanguageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadMessageLanguageResponseBody
	GetSuccess() *bool
}

type ReadMessageLanguageResponseBody struct {
	// The error code returned if the call failed. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result.
	//
	// example:
	//
	// zh-CN
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned if the call failed.
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

func (s ReadMessageLanguageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageLanguageResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageLanguageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadMessageLanguageResponseBody) GetData() *string {
	return s.Data
}

func (s *ReadMessageLanguageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadMessageLanguageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadMessageLanguageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadMessageLanguageResponseBody) SetCode(v string) *ReadMessageLanguageResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageLanguageResponseBody) SetData(v string) *ReadMessageLanguageResponseBody {
	s.Data = &v
	return s
}

func (s *ReadMessageLanguageResponseBody) SetMessage(v string) *ReadMessageLanguageResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageLanguageResponseBody) SetRequestId(v string) *ReadMessageLanguageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageLanguageResponseBody) SetSuccess(v bool) *ReadMessageLanguageResponseBody {
	s.Success = &v
	return s
}

func (s *ReadMessageLanguageResponseBody) Validate() error {
	return dara.Validate(s)
}
