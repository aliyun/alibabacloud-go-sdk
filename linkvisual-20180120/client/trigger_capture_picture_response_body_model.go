// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerCapturePictureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TriggerCapturePictureResponseBody
	GetCode() *string
	SetData(v string) *TriggerCapturePictureResponseBody
	GetData() *string
	SetErrorMessage(v string) *TriggerCapturePictureResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *TriggerCapturePictureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TriggerCapturePictureResponseBody
	GetSuccess() *bool
}

type TriggerCapturePictureResponseBody struct {
	// The response code of the operation. Valid values:
	//
	// - **200**: The call was successful.
	//
	// - Other values: An error occurred. For more information about error codes, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the image that is generated after the IPC device captures the image.
	//
	// >The image ID takes effect after the IPC device successfully uploads the image.
	//
	// example:
	//
	// eventPicId_12****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// request forbidden
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06DC77A0-4622-42DB-9EE0-25FIOHS82JK1
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

func (s TriggerCapturePictureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerCapturePictureResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerCapturePictureResponseBody) GetCode() *string {
	return s.Code
}

func (s *TriggerCapturePictureResponseBody) GetData() *string {
	return s.Data
}

func (s *TriggerCapturePictureResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *TriggerCapturePictureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerCapturePictureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TriggerCapturePictureResponseBody) SetCode(v string) *TriggerCapturePictureResponseBody {
	s.Code = &v
	return s
}

func (s *TriggerCapturePictureResponseBody) SetData(v string) *TriggerCapturePictureResponseBody {
	s.Data = &v
	return s
}

func (s *TriggerCapturePictureResponseBody) SetErrorMessage(v string) *TriggerCapturePictureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TriggerCapturePictureResponseBody) SetRequestId(v string) *TriggerCapturePictureResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerCapturePictureResponseBody) SetSuccess(v bool) *TriggerCapturePictureResponseBody {
	s.Success = &v
	return s
}

func (s *TriggerCapturePictureResponseBody) Validate() error {
	return dara.Validate(s)
}
