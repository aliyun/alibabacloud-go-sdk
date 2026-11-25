// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIAppJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIAppJob(v *GetAIAppJobResponseBodyAIAppJob) *GetAIAppJobResponseBody
	GetAIAppJob() *GetAIAppJobResponseBodyAIAppJob
	SetArtifacts(v []*GetAIAppJobResponseBodyArtifacts) *GetAIAppJobResponseBody
	GetArtifacts() []*GetAIAppJobResponseBodyArtifacts
	SetRequestId(v string) *GetAIAppJobResponseBody
	GetRequestId() *string
}

type GetAIAppJobResponseBody struct {
	// The task information.
	AIAppJob *GetAIAppJobResponseBodyAIAppJob `json:"AIAppJob,omitempty" xml:"AIAppJob,omitempty" type:"Struct"`
	// The output assets.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "Status": "completed",
	//
	//     "Type": "image",
	//
	//     "Url": "xxx",
	//
	//     "Name": "blue.png"
	//
	//   },
	//
	//   {
	//
	//     "Status": "completed",
	//
	//     "Type": "image",
	//
	//     "Url": "xxx",
	//
	//     "Name": "red.png"
	//
	//   }
	//
	// ]
	Artifacts []*GetAIAppJobResponseBodyArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAIAppJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIAppJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIAppJobResponseBody) GetAIAppJob() *GetAIAppJobResponseBodyAIAppJob {
	return s.AIAppJob
}

func (s *GetAIAppJobResponseBody) GetArtifacts() []*GetAIAppJobResponseBodyArtifacts {
	return s.Artifacts
}

func (s *GetAIAppJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIAppJobResponseBody) SetAIAppJob(v *GetAIAppJobResponseBodyAIAppJob) *GetAIAppJobResponseBody {
	s.AIAppJob = v
	return s
}

func (s *GetAIAppJobResponseBody) SetArtifacts(v []*GetAIAppJobResponseBodyArtifacts) *GetAIAppJobResponseBody {
	s.Artifacts = v
	return s
}

func (s *GetAIAppJobResponseBody) SetRequestId(v string) *GetAIAppJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIAppJobResponseBody) Validate() error {
	if s.AIAppJob != nil {
		if err := s.AIAppJob.Validate(); err != nil {
			return err
		}
	}
	if s.Artifacts != nil {
		for _, item := range s.Artifacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAIAppJobResponseBodyAIAppJob struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2026-10-06T02:06:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the task was completed.
	//
	// example:
	//
	// 2026-10-06T16:59:25Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// a2880a70cec1414cac476bce0bd39ad7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The node results of the workflow task. The results are in JSON format and vary based on the workflow configuration.
	//
	// example:
	//
	// {\\"Key\\":\\"Value\\"}
	NodeResults *string `json:"NodeResults,omitempty" xml:"NodeResults,omitempty"`
	// The task outputs.
	//
	// example:
	//
	// {\\"Status\\":\\"Success\\"}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The task status. Valid values:
	//
	// - init: The task is created.
	//
	// - running: The task is running.
	//
	// - succeeded: The task succeeded.
	//
	// - failed: The task failed.
	//
	// - stopped: The task is stopped.
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The custom information.
	//
	// example:
	//
	// {\\"Key\\":\\"Value\\"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The application version.
	//
	// example:
	//
	// xxxxxx
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAIAppJobResponseBodyAIAppJob) String() string {
	return dara.Prettify(s)
}

func (s GetAIAppJobResponseBodyAIAppJob) GoString() string {
	return s.String()
}

func (s *GetAIAppJobResponseBodyAIAppJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAIAppJobResponseBodyAIAppJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetAIAppJobResponseBodyAIAppJob) GetJobId() *string {
	return s.JobId
}

func (s *GetAIAppJobResponseBodyAIAppJob) GetNodeResults() *string {
	return s.NodeResults
}

func (s *GetAIAppJobResponseBodyAIAppJob) GetOutputs() *string {
	return s.Outputs
}

func (s *GetAIAppJobResponseBodyAIAppJob) GetStatus() *string {
	return s.Status
}

func (s *GetAIAppJobResponseBodyAIAppJob) GetUserData() *string {
	return s.UserData
}

func (s *GetAIAppJobResponseBodyAIAppJob) GetVersion() *string {
	return s.Version
}

func (s *GetAIAppJobResponseBodyAIAppJob) SetCreateTime(v string) *GetAIAppJobResponseBodyAIAppJob {
	s.CreateTime = &v
	return s
}

func (s *GetAIAppJobResponseBodyAIAppJob) SetFinishTime(v string) *GetAIAppJobResponseBodyAIAppJob {
	s.FinishTime = &v
	return s
}

func (s *GetAIAppJobResponseBodyAIAppJob) SetJobId(v string) *GetAIAppJobResponseBodyAIAppJob {
	s.JobId = &v
	return s
}

func (s *GetAIAppJobResponseBodyAIAppJob) SetNodeResults(v string) *GetAIAppJobResponseBodyAIAppJob {
	s.NodeResults = &v
	return s
}

func (s *GetAIAppJobResponseBodyAIAppJob) SetOutputs(v string) *GetAIAppJobResponseBodyAIAppJob {
	s.Outputs = &v
	return s
}

func (s *GetAIAppJobResponseBodyAIAppJob) SetStatus(v string) *GetAIAppJobResponseBodyAIAppJob {
	s.Status = &v
	return s
}

func (s *GetAIAppJobResponseBodyAIAppJob) SetUserData(v string) *GetAIAppJobResponseBodyAIAppJob {
	s.UserData = &v
	return s
}

func (s *GetAIAppJobResponseBodyAIAppJob) SetVersion(v string) *GetAIAppJobResponseBodyAIAppJob {
	s.Version = &v
	return s
}

func (s *GetAIAppJobResponseBodyAIAppJob) Validate() error {
	return dara.Validate(s)
}

type GetAIAppJobResponseBodyArtifacts struct {
	// The asset ID, which is the media asset ID obtained after registration.
	//
	// example:
	//
	// xxxx77i16crgqtn96ip914bmixxxx
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The asset name.
	//
	// example:
	//
	// test_00x
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The asset status.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The asset type. Valid values: image, audio, video, document, and text.
	//
	// example:
	//
	// image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The asset URL.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1CN0157aHKQ1LlJ45vCQ8w_!!6000000001339-0-tps-1080-773.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIAppJobResponseBodyArtifacts) String() string {
	return dara.Prettify(s)
}

func (s GetAIAppJobResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *GetAIAppJobResponseBodyArtifacts) GetAssetId() *string {
	return s.AssetId
}

func (s *GetAIAppJobResponseBodyArtifacts) GetName() *string {
	return s.Name
}

func (s *GetAIAppJobResponseBodyArtifacts) GetStatus() *string {
	return s.Status
}

func (s *GetAIAppJobResponseBodyArtifacts) GetType() *string {
	return s.Type
}

func (s *GetAIAppJobResponseBodyArtifacts) GetUrl() *string {
	return s.Url
}

func (s *GetAIAppJobResponseBodyArtifacts) SetAssetId(v string) *GetAIAppJobResponseBodyArtifacts {
	s.AssetId = &v
	return s
}

func (s *GetAIAppJobResponseBodyArtifacts) SetName(v string) *GetAIAppJobResponseBodyArtifacts {
	s.Name = &v
	return s
}

func (s *GetAIAppJobResponseBodyArtifacts) SetStatus(v string) *GetAIAppJobResponseBodyArtifacts {
	s.Status = &v
	return s
}

func (s *GetAIAppJobResponseBodyArtifacts) SetType(v string) *GetAIAppJobResponseBodyArtifacts {
	s.Type = &v
	return s
}

func (s *GetAIAppJobResponseBodyArtifacts) SetUrl(v string) *GetAIAppJobResponseBodyArtifacts {
	s.Url = &v
	return s
}

func (s *GetAIAppJobResponseBodyArtifacts) Validate() error {
	return dara.Validate(s)
}
