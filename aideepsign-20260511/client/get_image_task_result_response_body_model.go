// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageTaskResultResponseBody
	GetCode() *string
	SetErrorMessage(v string) *GetImageTaskResultResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetImageTaskResultResponseBody
	GetHttpStatusCode() *int32
	SetImages(v []*GetImageTaskResultResponseBodyImages) *GetImageTaskResultResponseBody
	GetImages() []*GetImageTaskResultResponseBodyImages
	SetMessage(v string) *GetImageTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageTaskResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetImageTaskResultResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetImageTaskResultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetImageTaskResultResponseBody
	GetTaskId() *string
}

type GetImageTaskResultResponseBody struct {
	// The business error code. The value `OK` is returned if the request succeeds.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is returned only when the task status is `failed`.
	//
	// example:
	//
	// Instance access forbidden.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code. The value `200` is returned if the request succeeds.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The list of generated images. This parameter is returned only when `Status` is `succeeded`.
	Images []*GetImageTaskResultResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The additional information. The value `success` is returned if the request succeeds. An error message is returned if the task fails. This parameter is returned only when `Status` is `failed`.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task status. Valid values: `pending` (waiting), `running` (in progress), `succeeded` (completed), `failed` (failed).
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImageTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageTaskResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetImageTaskResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetImageTaskResultResponseBody) GetImages() []*GetImageTaskResultResponseBodyImages {
	return s.Images
}

func (s *GetImageTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageTaskResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetImageTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageTaskResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageTaskResultResponseBody) SetCode(v string) *GetImageTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageTaskResultResponseBody) SetErrorMessage(v string) *GetImageTaskResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetImageTaskResultResponseBody) SetHttpStatusCode(v int32) *GetImageTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImageTaskResultResponseBody) SetImages(v []*GetImageTaskResultResponseBodyImages) *GetImageTaskResultResponseBody {
	s.Images = v
	return s
}

func (s *GetImageTaskResultResponseBody) SetMessage(v string) *GetImageTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageTaskResultResponseBody) SetRequestId(v string) *GetImageTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageTaskResultResponseBody) SetStatus(v string) *GetImageTaskResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetImageTaskResultResponseBody) SetSuccess(v bool) *GetImageTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageTaskResultResponseBody) SetTaskId(v string) *GetImageTaskResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetImageTaskResultResponseBody) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetImageTaskResultResponseBodyImages struct {
	// The `ObjectKey` of the image in OSS. You can use this value in subsequent API calls.
	//
	// example:
	//
	// deepsign/123456789/image-generation/abc12345-def6-7890-abcd-ef1234567890.png
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The pre-signed download URL of the image. The URL is valid for 1 hour.
	//
	// example:
	//
	// https://bucket.oss-cn-hangzhou.aliyuncs.com/deepsign/123456789/image-generation/abc12345.png?Expires=1718700000&OSSAccessKeyId=...
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetImageTaskResultResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s GetImageTaskResultResponseBodyImages) GoString() string {
	return s.String()
}

func (s *GetImageTaskResultResponseBodyImages) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *GetImageTaskResultResponseBodyImages) GetUrl() *string {
	return s.Url
}

func (s *GetImageTaskResultResponseBodyImages) SetObjectKey(v string) *GetImageTaskResultResponseBodyImages {
	s.ObjectKey = &v
	return s
}

func (s *GetImageTaskResultResponseBodyImages) SetUrl(v string) *GetImageTaskResultResponseBodyImages {
	s.Url = &v
	return s
}

func (s *GetImageTaskResultResponseBodyImages) Validate() error {
	return dara.Validate(s)
}
