// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceUserPictureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFaceUserPictureResponseBody
	GetCode() *string
	SetErrorMessage(v string) *DeleteFaceUserPictureResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteFaceUserPictureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFaceUserPictureResponseBody
	GetSuccess() *bool
}

type DeleteFaceUserPictureResponseBody struct {
	// The return code of the operation:
	//
	// - **200**: The call is successful.
	//
	// - Other values: error codes. For more information, see [Error codes](https://help.aliyun.com/document_detail/288314.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// face user not exist
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 4de2c367-c1db-417c-aa15-8c585e595d92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
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

func (s DeleteFaceUserPictureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceUserPictureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserPictureResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFaceUserPictureResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFaceUserPictureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceUserPictureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFaceUserPictureResponseBody) SetCode(v string) *DeleteFaceUserPictureResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceUserPictureResponseBody) SetErrorMessage(v string) *DeleteFaceUserPictureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceUserPictureResponseBody) SetRequestId(v string) *DeleteFaceUserPictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceUserPictureResponseBody) SetSuccess(v bool) *DeleteFaceUserPictureResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFaceUserPictureResponseBody) Validate() error {
	return dara.Validate(s)
}
