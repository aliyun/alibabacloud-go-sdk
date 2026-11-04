// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDevicePictureLifeCycleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetDevicePictureLifeCycleResponseBody
	GetCode() *string
	SetErrorMessage(v string) *SetDevicePictureLifeCycleResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SetDevicePictureLifeCycleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetDevicePictureLifeCycleResponseBody
	GetSuccess() *bool
}

type SetDevicePictureLifeCycleResponseBody struct {
	// The response code. Valid values:
	//
	// - **200**: The call was successful.
	//
	// - Other values: The error code. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
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
	// 449ECE75-A36A-4ACE-B253-147C78C440C5
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

func (s SetDevicePictureLifeCycleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDevicePictureLifeCycleResponseBody) GoString() string {
	return s.String()
}

func (s *SetDevicePictureLifeCycleResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetDevicePictureLifeCycleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SetDevicePictureLifeCycleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDevicePictureLifeCycleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetDevicePictureLifeCycleResponseBody) SetCode(v string) *SetDevicePictureLifeCycleResponseBody {
	s.Code = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponseBody) SetErrorMessage(v string) *SetDevicePictureLifeCycleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponseBody) SetRequestId(v string) *SetDevicePictureLifeCycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponseBody) SetSuccess(v bool) *SetDevicePictureLifeCycleResponseBody {
	s.Success = &v
	return s
}

func (s *SetDevicePictureLifeCycleResponseBody) Validate() error {
	return dara.Validate(s)
}
