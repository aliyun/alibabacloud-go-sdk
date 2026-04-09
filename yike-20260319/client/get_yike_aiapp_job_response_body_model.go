// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAIAppJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetYikeAIAppJobResponseBody
	GetAppId() *string
	SetAppParams(v string) *GetYikeAIAppJobResponseBody
	GetAppParams() *string
	SetExecutionFinishTime(v string) *GetYikeAIAppJobResponseBody
	GetExecutionFinishTime() *string
	SetExecutionStartTime(v string) *GetYikeAIAppJobResponseBody
	GetExecutionStartTime() *string
	SetFolderId(v string) *GetYikeAIAppJobResponseBody
	GetFolderId() *string
	SetJobId(v string) *GetYikeAIAppJobResponseBody
	GetJobId() *string
	SetProductionId(v string) *GetYikeAIAppJobResponseBody
	GetProductionId() *string
	SetRequestId(v string) *GetYikeAIAppJobResponseBody
	GetRequestId() *string
	SetResult(v *GetYikeAIAppJobResponseBodyResult) *GetYikeAIAppJobResponseBody
	GetResult() *GetYikeAIAppJobResponseBodyResult
	SetStatus(v string) *GetYikeAIAppJobResponseBody
	GetStatus() *string
}

type GetYikeAIAppJobResponseBody struct {
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// {\\"LoadImage.1.TargetImage\\":\\"MediaId1\\"}
	AppParams *string `json:"AppParams,omitempty" xml:"AppParams,omitempty"`
	// example:
	//
	// 2026-02-06T18:53:18.809+08:00
	ExecutionFinishTime *string `json:"ExecutionFinishTime,omitempty" xml:"ExecutionFinishTime,omitempty"`
	// example:
	//
	// 2026-02-06T18:53:34.001+08:00
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	// example:
	//
	// folder-u3ilwhoc36ux9a****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ****cdb3e74639973036bc84****
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// example:
	//
	// 0622C702-41BE-467E-AF2E-883D4517962E
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetYikeAIAppJobResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetYikeAIAppJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAIAppJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeAIAppJobResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetYikeAIAppJobResponseBody) GetAppParams() *string {
	return s.AppParams
}

func (s *GetYikeAIAppJobResponseBody) GetExecutionFinishTime() *string {
	return s.ExecutionFinishTime
}

func (s *GetYikeAIAppJobResponseBody) GetExecutionStartTime() *string {
	return s.ExecutionStartTime
}

func (s *GetYikeAIAppJobResponseBody) GetFolderId() *string {
	return s.FolderId
}

func (s *GetYikeAIAppJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeAIAppJobResponseBody) GetProductionId() *string {
	return s.ProductionId
}

func (s *GetYikeAIAppJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeAIAppJobResponseBody) GetResult() *GetYikeAIAppJobResponseBodyResult {
	return s.Result
}

func (s *GetYikeAIAppJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetYikeAIAppJobResponseBody) SetAppId(v string) *GetYikeAIAppJobResponseBody {
	s.AppId = &v
	return s
}

func (s *GetYikeAIAppJobResponseBody) SetAppParams(v string) *GetYikeAIAppJobResponseBody {
	s.AppParams = &v
	return s
}

func (s *GetYikeAIAppJobResponseBody) SetExecutionFinishTime(v string) *GetYikeAIAppJobResponseBody {
	s.ExecutionFinishTime = &v
	return s
}

func (s *GetYikeAIAppJobResponseBody) SetExecutionStartTime(v string) *GetYikeAIAppJobResponseBody {
	s.ExecutionStartTime = &v
	return s
}

func (s *GetYikeAIAppJobResponseBody) SetFolderId(v string) *GetYikeAIAppJobResponseBody {
	s.FolderId = &v
	return s
}

func (s *GetYikeAIAppJobResponseBody) SetJobId(v string) *GetYikeAIAppJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetYikeAIAppJobResponseBody) SetProductionId(v string) *GetYikeAIAppJobResponseBody {
	s.ProductionId = &v
	return s
}

func (s *GetYikeAIAppJobResponseBody) SetRequestId(v string) *GetYikeAIAppJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeAIAppJobResponseBody) SetResult(v *GetYikeAIAppJobResponseBodyResult) *GetYikeAIAppJobResponseBody {
	s.Result = v
	return s
}

func (s *GetYikeAIAppJobResponseBody) SetStatus(v string) *GetYikeAIAppJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetYikeAIAppJobResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetYikeAIAppJobResponseBodyResult struct {
	AudioResult []*GetYikeAIAppJobResponseBodyResultAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Repeated"`
	ImageResult []*GetYikeAIAppJobResponseBodyResultImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	VideoResult []*GetYikeAIAppJobResponseBodyResultVideoResult `json:"VideoResult,omitempty" xml:"VideoResult,omitempty" type:"Repeated"`
}

func (s GetYikeAIAppJobResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAIAppJobResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetYikeAIAppJobResponseBodyResult) GetAudioResult() []*GetYikeAIAppJobResponseBodyResultAudioResult {
	return s.AudioResult
}

func (s *GetYikeAIAppJobResponseBodyResult) GetImageResult() []*GetYikeAIAppJobResponseBodyResultImageResult {
	return s.ImageResult
}

func (s *GetYikeAIAppJobResponseBodyResult) GetVideoResult() []*GetYikeAIAppJobResponseBodyResultVideoResult {
	return s.VideoResult
}

func (s *GetYikeAIAppJobResponseBodyResult) SetAudioResult(v []*GetYikeAIAppJobResponseBodyResultAudioResult) *GetYikeAIAppJobResponseBodyResult {
	s.AudioResult = v
	return s
}

func (s *GetYikeAIAppJobResponseBodyResult) SetImageResult(v []*GetYikeAIAppJobResponseBodyResultImageResult) *GetYikeAIAppJobResponseBodyResult {
	s.ImageResult = v
	return s
}

func (s *GetYikeAIAppJobResponseBodyResult) SetVideoResult(v []*GetYikeAIAppJobResponseBodyResultVideoResult) *GetYikeAIAppJobResponseBodyResult {
	s.VideoResult = v
	return s
}

func (s *GetYikeAIAppJobResponseBodyResult) Validate() error {
	if s.AudioResult != nil {
		for _, item := range s.AudioResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageResult != nil {
		for _, item := range s.ImageResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoResult != nil {
		for _, item := range s.VideoResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetYikeAIAppJobResponseBodyResultAudioResult struct {
	// example:
	//
	// 1a7852b032c371f0b64fe6f6c74b****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// http://outin-***.oss-cn-shanghai.aliyuncs.com/stream/48555e8b-181dd5a8c07/48555e8b-181dd5a8c07.mp3
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s GetYikeAIAppJobResponseBodyResultAudioResult) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAIAppJobResponseBodyResultAudioResult) GoString() string {
	return s.String()
}

func (s *GetYikeAIAppJobResponseBodyResultAudioResult) GetMediaId() *string {
	return s.MediaId
}

func (s *GetYikeAIAppJobResponseBodyResultAudioResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetYikeAIAppJobResponseBodyResultAudioResult) SetMediaId(v string) *GetYikeAIAppJobResponseBodyResultAudioResult {
	s.MediaId = &v
	return s
}

func (s *GetYikeAIAppJobResponseBodyResultAudioResult) SetOutputUrl(v string) *GetYikeAIAppJobResponseBodyResultAudioResult {
	s.OutputUrl = &v
	return s
}

func (s *GetYikeAIAppJobResponseBodyResultAudioResult) Validate() error {
	return dara.Validate(s)
}

type GetYikeAIAppJobResponseBodyResultImageResult struct {
	// example:
	//
	// 1a7852b032c371f0b64fe6f6c74b****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// http://outin-***.oss-cn-shanghai.aliyuncs.com/stream/48555e8b-181dd5a8c07/48555e8b-181dd5a8c07.png
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s GetYikeAIAppJobResponseBodyResultImageResult) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAIAppJobResponseBodyResultImageResult) GoString() string {
	return s.String()
}

func (s *GetYikeAIAppJobResponseBodyResultImageResult) GetMediaId() *string {
	return s.MediaId
}

func (s *GetYikeAIAppJobResponseBodyResultImageResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetYikeAIAppJobResponseBodyResultImageResult) SetMediaId(v string) *GetYikeAIAppJobResponseBodyResultImageResult {
	s.MediaId = &v
	return s
}

func (s *GetYikeAIAppJobResponseBodyResultImageResult) SetOutputUrl(v string) *GetYikeAIAppJobResponseBodyResultImageResult {
	s.OutputUrl = &v
	return s
}

func (s *GetYikeAIAppJobResponseBodyResultImageResult) Validate() error {
	return dara.Validate(s)
}

type GetYikeAIAppJobResponseBodyResultVideoResult struct {
	// example:
	//
	// 1a7852b032c371f0b64fe6f6c74b****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// http://outin-***.oss-cn-shanghai.aliyuncs.com/stream/48555e8b-181dd5a8c07/48555e8b-181dd5a8c07.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s GetYikeAIAppJobResponseBodyResultVideoResult) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAIAppJobResponseBodyResultVideoResult) GoString() string {
	return s.String()
}

func (s *GetYikeAIAppJobResponseBodyResultVideoResult) GetMediaId() *string {
	return s.MediaId
}

func (s *GetYikeAIAppJobResponseBodyResultVideoResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetYikeAIAppJobResponseBodyResultVideoResult) SetMediaId(v string) *GetYikeAIAppJobResponseBodyResultVideoResult {
	s.MediaId = &v
	return s
}

func (s *GetYikeAIAppJobResponseBodyResultVideoResult) SetOutputUrl(v string) *GetYikeAIAppJobResponseBodyResultVideoResult {
	s.OutputUrl = &v
	return s
}

func (s *GetYikeAIAppJobResponseBodyResultVideoResult) Validate() error {
	return dara.Validate(s)
}
