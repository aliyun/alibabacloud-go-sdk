// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePictureSearchAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePictureSearchAppResponseBody
	GetCode() *string
	SetErrorMessage(v string) *UpdatePictureSearchAppResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdatePictureSearchAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePictureSearchAppResponseBody
	GetSuccess() *bool
}

type UpdatePictureSearchAppResponseBody struct {
	// The error code returned when the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// Description is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CC6B9D69-627C-4C9D-AEAA-0BC08AB6DBED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePictureSearchAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePictureSearchAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePictureSearchAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePictureSearchAppResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdatePictureSearchAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePictureSearchAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePictureSearchAppResponseBody) SetCode(v string) *UpdatePictureSearchAppResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePictureSearchAppResponseBody) SetErrorMessage(v string) *UpdatePictureSearchAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePictureSearchAppResponseBody) SetRequestId(v string) *UpdatePictureSearchAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePictureSearchAppResponseBody) SetSuccess(v bool) *UpdatePictureSearchAppResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePictureSearchAppResponseBody) Validate() error {
	return dara.Validate(s)
}
