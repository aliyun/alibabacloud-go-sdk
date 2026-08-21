// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIImageJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIImageJobList(v []*GetAIImageJobsResponseBodyAIImageJobList) *GetAIImageJobsResponseBody
	GetAIImageJobList() []*GetAIImageJobsResponseBodyAIImageJobList
	SetRequestId(v string) *GetAIImageJobsResponseBody
	GetRequestId() *string
}

type GetAIImageJobsResponseBody struct {
	// The list of AI image processing tasks.
	AIImageJobList []*GetAIImageJobsResponseBodyAIImageJobList `json:"AIImageJobList,omitempty" xml:"AIImageJobList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7721B494-1F78-4E*****E8-A7CEE7315BFA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAIImageJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIImageJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIImageJobsResponseBody) GetAIImageJobList() []*GetAIImageJobsResponseBodyAIImageJobList {
	return s.AIImageJobList
}

func (s *GetAIImageJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIImageJobsResponseBody) SetAIImageJobList(v []*GetAIImageJobsResponseBodyAIImageJobList) *GetAIImageJobsResponseBody {
	s.AIImageJobList = v
	return s
}

func (s *GetAIImageJobsResponseBody) SetRequestId(v string) *GetAIImageJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIImageJobsResponseBody) Validate() error {
	if s.AIImageJobList != nil {
		for _, item := range s.AIImageJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAIImageJobsResponseBodyAIImageJobList struct {
	// The OSS URL of the AI image.
	//
	// >This is the task result. The URL does not contain complete authentication information. To obtain authentication information, generate it yourself or call the [ListAIImage](https://help.aliyun.com/document_detail/186924.html) operation to retrieve the media asset result.
	//
	// example:
	//
	// [{"Score":5.035636554444242,"Url":"http://outin-*****.oss-cn-shanghai.aliyuncs.com/357a8748c577*****789d2726e6436aa/image/ai/b0a7612554d*****5cbe3-00001.gif"}]
	AIImageResult *string `json:"AIImageResult,omitempty" xml:"AIImageResult,omitempty"`
	// The error code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the AI image processing task was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2020-10-15T03:30:03Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the AI image processing task.
	//
	// example:
	//
	// cf08a2c6e11e*****de1711b738b9067
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The task status. Valid values:
	//
	// - **success**: The task is successful.
	//
	// - **fail**: The task failed.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The snapshot of the configuration information of the specified template when the task was submitted.
	//
	// example:
	//
	// {"Format":"gif","SetDefaultCover":"true"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The AI template ID.
	//
	// example:
	//
	// 5a86a00f15194*****d7fe7de1b4a173
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The custom settings.
	//
	// - The value must be a JSON string.
	//
	// - The value must contain the MessageCallback or Extend parameter.
	//
	// - The maximum length is 512 bytes.
	//
	// For more information about the parameter structure, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// example:
	//
	// {"Extend":{"localId":"****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The video ID.
	//
	// example:
	//
	// 357a8748c577*****789d2726e6436aa
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetAIImageJobsResponseBodyAIImageJobList) String() string {
	return dara.Prettify(s)
}

func (s GetAIImageJobsResponseBodyAIImageJobList) GoString() string {
	return s.String()
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) GetAIImageResult() *string {
	return s.AIImageResult
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) GetCode() *string {
	return s.Code
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) GetJobId() *string {
	return s.JobId
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) GetMessage() *string {
	return s.Message
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) GetStatus() *string {
	return s.Status
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) GetUserData() *string {
	return s.UserData
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) GetVideoId() *string {
	return s.VideoId
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetAIImageResult(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.AIImageResult = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetCode(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.Code = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetCreationTime(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.CreationTime = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetJobId(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.JobId = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetMessage(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.Message = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetStatus(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.Status = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetTemplateConfig(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.TemplateConfig = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetTemplateId(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.TemplateId = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetUserData(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.UserData = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetVideoId(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.VideoId = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) Validate() error {
	return dara.Validate(s)
}
