// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTimeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTimeTemplateResponseBody
	GetCode() *string
	SetErrorMessage(v string) *UpdateTimeTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTimeTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTimeTemplateResponseBody
	GetSuccess() *bool
}

type UpdateTimeTemplateResponseBody struct {
	// The return code.
	//
	// Valid values:
	//
	// - **200**: The call is successful.
	//
	// - Other values: An error code is returned. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// tenant auth error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06DC77A0-4622-42DB-9EE0-25FIOHS82JK1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// Valid values:
	//
	// - **true**: The call is successful.
	//
	// - **false**: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTimeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTimeTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTimeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTimeTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTimeTemplateResponseBody) SetCode(v string) *UpdateTimeTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTimeTemplateResponseBody) SetErrorMessage(v string) *UpdateTimeTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTimeTemplateResponseBody) SetRequestId(v string) *UpdateTimeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTimeTemplateResponseBody) SetSuccess(v bool) *UpdateTimeTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTimeTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
