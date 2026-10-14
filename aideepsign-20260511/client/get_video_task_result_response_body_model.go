// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVideoTaskResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetVideoTaskResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVideoTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVideoTaskResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetVideoTaskResultResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetVideoTaskResultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetVideoTaskResultResponseBody
	GetTaskId() *string
	SetVideos(v []*GetVideoTaskResultResponseBodyVideos) *GetVideoTaskResultResponseBody
	GetVideos() []*GetVideoTaskResultResponseBodyVideos
}

type GetVideoTaskResultResponseBody struct {
	// **Business error code.*	- Returns "OK" if the request succeeds.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The cause of the failure. This field is returned only when the task status is failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task status. Valid values: running, succeeded, and failed. Set the polling interval to at least 5 seconds.
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request is successful.
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
	// The list of video information.
	Videos []*GetVideoTaskResultResponseBodyVideos `json:"Videos,omitempty" xml:"Videos,omitempty" type:"Repeated"`
}

func (s GetVideoTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoTaskResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVideoTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVideoTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoTaskResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetVideoTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVideoTaskResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoTaskResultResponseBody) GetVideos() []*GetVideoTaskResultResponseBodyVideos {
	return s.Videos
}

func (s *GetVideoTaskResultResponseBody) SetCode(v string) *GetVideoTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoTaskResultResponseBody) SetHttpStatusCode(v int32) *GetVideoTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVideoTaskResultResponseBody) SetMessage(v string) *GetVideoTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoTaskResultResponseBody) SetRequestId(v string) *GetVideoTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoTaskResultResponseBody) SetStatus(v string) *GetVideoTaskResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetVideoTaskResultResponseBody) SetSuccess(v bool) *GetVideoTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetVideoTaskResultResponseBody) SetTaskId(v string) *GetVideoTaskResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetVideoTaskResultResponseBody) SetVideos(v []*GetVideoTaskResultResponseBodyVideos) *GetVideoTaskResultResponseBody {
	s.Videos = v
	return s
}

func (s *GetVideoTaskResultResponseBody) Validate() error {
	if s.Videos != nil {
		for _, item := range s.Videos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoTaskResultResponseBodyVideos struct {
	// The OSS key of the output object.
	//
	// example:
	//
	// ai_deep_sign/1192853035118460/image-generation/2026/06/25/024d6974-1b14-4741-8919-9d02b7e3901b.mp4
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The presigned download URL of the video stored in OSS after C2PA signing.
	//
	// example:
	//
	// http://duguang-database-public.oss-cn-hangzhou.aliyuncs.com/cosmetic_plicense_init/024d6974-1b14-4741-8919-9d02b7e3901b.mp4
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetVideoTaskResultResponseBodyVideos) String() string {
	return dara.Prettify(s)
}

func (s GetVideoTaskResultResponseBodyVideos) GoString() string {
	return s.String()
}

func (s *GetVideoTaskResultResponseBodyVideos) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *GetVideoTaskResultResponseBodyVideos) GetUrl() *string {
	return s.Url
}

func (s *GetVideoTaskResultResponseBodyVideos) SetObjectKey(v string) *GetVideoTaskResultResponseBodyVideos {
	s.ObjectKey = &v
	return s
}

func (s *GetVideoTaskResultResponseBodyVideos) SetUrl(v string) *GetVideoTaskResultResponseBodyVideos {
	s.Url = &v
	return s
}

func (s *GetVideoTaskResultResponseBodyVideos) Validate() error {
	return dara.Validate(s)
}
