// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetYikeAIAppJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobList(v []*BatchGetYikeAIAppJobResponseBodyJobList) *BatchGetYikeAIAppJobResponseBody
	GetJobList() []*BatchGetYikeAIAppJobResponseBodyJobList
	SetRequestId(v string) *BatchGetYikeAIAppJobResponseBody
	GetRequestId() *string
}

type BatchGetYikeAIAppJobResponseBody struct {
	JobList []*BatchGetYikeAIAppJobResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetYikeAIAppJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAIAppJobResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAIAppJobResponseBody) GetJobList() []*BatchGetYikeAIAppJobResponseBodyJobList {
	return s.JobList
}

func (s *BatchGetYikeAIAppJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetYikeAIAppJobResponseBody) SetJobList(v []*BatchGetYikeAIAppJobResponseBodyJobList) *BatchGetYikeAIAppJobResponseBody {
	s.JobList = v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBody) SetRequestId(v string) *BatchGetYikeAIAppJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBody) Validate() error {
	if s.JobList != nil {
		for _, item := range s.JobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetYikeAIAppJobResponseBodyJobList struct {
	// example:
	//
	// 9e09955d662a
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// {\\"LoadImage.1.TargetImage\\":\\"MediaId1\\"}
	AppInputConfig *string `json:"AppInputConfig,omitempty" xml:"AppInputConfig,omitempty"`
	// example:
	//
	// 2026-02-06T18:53:34.001+08:00
	ExecutionFinishTime *string `json:"ExecutionFinishTime,omitempty" xml:"ExecutionFinishTime,omitempty"`
	// example:
	//
	// 2026-02-06T18:53:18.809+08:00
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	// example:
	//
	// pd_0617169475
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// pd_0617169475
	ProductionId *string                                        `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	Result       *BatchGetYikeAIAppJobResponseBodyJobListResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s BatchGetYikeAIAppJobResponseBodyJobList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAIAppJobResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) GetAppId() *string {
	return s.AppId
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) GetAppInputConfig() *string {
	return s.AppInputConfig
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) GetExecutionFinishTime() *string {
	return s.ExecutionFinishTime
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) GetExecutionStartTime() *string {
	return s.ExecutionStartTime
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) GetFolderId() *string {
	return s.FolderId
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) GetJobId() *string {
	return s.JobId
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) GetProductionId() *string {
	return s.ProductionId
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) GetResult() *BatchGetYikeAIAppJobResponseBodyJobListResult {
	return s.Result
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) GetStatus() *string {
	return s.Status
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) SetAppId(v string) *BatchGetYikeAIAppJobResponseBodyJobList {
	s.AppId = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) SetAppInputConfig(v string) *BatchGetYikeAIAppJobResponseBodyJobList {
	s.AppInputConfig = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) SetExecutionFinishTime(v string) *BatchGetYikeAIAppJobResponseBodyJobList {
	s.ExecutionFinishTime = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) SetExecutionStartTime(v string) *BatchGetYikeAIAppJobResponseBodyJobList {
	s.ExecutionStartTime = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) SetFolderId(v string) *BatchGetYikeAIAppJobResponseBodyJobList {
	s.FolderId = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) SetJobId(v string) *BatchGetYikeAIAppJobResponseBodyJobList {
	s.JobId = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) SetProductionId(v string) *BatchGetYikeAIAppJobResponseBodyJobList {
	s.ProductionId = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) SetResult(v *BatchGetYikeAIAppJobResponseBodyJobListResult) *BatchGetYikeAIAppJobResponseBodyJobList {
	s.Result = v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) SetStatus(v string) *BatchGetYikeAIAppJobResponseBodyJobList {
	s.Status = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobList) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetYikeAIAppJobResponseBodyJobListResult struct {
	AudioResult []*BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Repeated"`
	ImageResult []*BatchGetYikeAIAppJobResponseBodyJobListResultImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	VideoResult []*BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult `json:"VideoResult,omitempty" xml:"VideoResult,omitempty" type:"Repeated"`
}

func (s BatchGetYikeAIAppJobResponseBodyJobListResult) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAIAppJobResponseBodyJobListResult) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResult) GetAudioResult() []*BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult {
	return s.AudioResult
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResult) GetImageResult() []*BatchGetYikeAIAppJobResponseBodyJobListResultImageResult {
	return s.ImageResult
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResult) GetVideoResult() []*BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult {
	return s.VideoResult
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResult) SetAudioResult(v []*BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult) *BatchGetYikeAIAppJobResponseBodyJobListResult {
	s.AudioResult = v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResult) SetImageResult(v []*BatchGetYikeAIAppJobResponseBodyJobListResultImageResult) *BatchGetYikeAIAppJobResponseBodyJobListResult {
	s.ImageResult = v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResult) SetVideoResult(v []*BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult) *BatchGetYikeAIAppJobResponseBodyJobListResult {
	s.VideoResult = v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResult) Validate() error {
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

type BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult struct {
	// example:
	//
	// bb1dfa20a0c971f08c94e7f6d6496202
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// https://dtlive-bj.oss-cn-beijing.aliyuncs.com/cover/01e1271d-ff4f-4689-9c20-e1df81486859_open_live_cover.mp3
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult) SetMediaId(v string) *BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult {
	s.MediaId = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult) SetOutputUrl(v string) *BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult {
	s.OutputUrl = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultAudioResult) Validate() error {
	return dara.Validate(s)
}

type BatchGetYikeAIAppJobResponseBodyJobListResultImageResult struct {
	// example:
	//
	// 318d6350a57d71f08c9ae7f7d4496302
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// https://dtlive-bj.oss-cn-beijing.aliyuncs.com/cover/01e1271d-ff4f-4689-9c20-e1df81486859_open_live_cover.jpg
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s BatchGetYikeAIAppJobResponseBodyJobListResultImageResult) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAIAppJobResponseBodyJobListResultImageResult) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultImageResult) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultImageResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultImageResult) SetMediaId(v string) *BatchGetYikeAIAppJobResponseBodyJobListResultImageResult {
	s.MediaId = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultImageResult) SetOutputUrl(v string) *BatchGetYikeAIAppJobResponseBodyJobListResultImageResult {
	s.OutputUrl = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultImageResult) Validate() error {
	return dara.Validate(s)
}

type BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult struct {
	// example:
	//
	// ec2de25068fd71f0a48cf7e6c4596302
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// https://dtlive-bj.oss-cn-beijing.aliyuncs.com/cover/01e1271d-ff4f-4689-9c20-e1df81486859_open_live_cover.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult) SetMediaId(v string) *BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult {
	s.MediaId = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult) SetOutputUrl(v string) *BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult {
	s.OutputUrl = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponseBodyJobListResultVideoResult) Validate() error {
	return dara.Validate(s)
}
