// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveEditingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClips(v string) *SubmitLiveEditingJobRequest
	GetClips() *string
	SetLiveStreamConfig(v string) *SubmitLiveEditingJobRequest
	GetLiveStreamConfig() *string
	SetMediaProduceConfig(v string) *SubmitLiveEditingJobRequest
	GetMediaProduceConfig() *string
	SetOutputMediaConfig(v string) *SubmitLiveEditingJobRequest
	GetOutputMediaConfig() *string
	SetOutputMediaTarget(v string) *SubmitLiveEditingJobRequest
	GetOutputMediaTarget() *string
	SetProjectId(v string) *SubmitLiveEditingJobRequest
	GetProjectId() *string
	SetUserData(v string) *SubmitLiveEditingJobRequest
	GetUserData() *string
}

type SubmitLiveEditingJobRequest struct {
	// The clips in the JSON array format. The output video is created by merging these clips sequentially.
	//
	// Each clip has a start time and an end time. If no live stream parameters are specified, the outer live stream configurations apply. The start and end timestamps are in UTC. For more information about the parameters, see the "Clip" section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"StartTime\\": \\" 2021-06-21T08:01:00Z\\",  \\"EndTime\\": \\" 2021-06-21T08:03:00Z\\" ,  "AppName": "app", "DomainName": "domain.com", "StreamName": "stream"},  {\\"StartTime\\": \\" 2021-06-21T08:05:00Z\\",  \\"EndTime\\": \\" 2021-06-21T08:09:00Z\\" }]
	Clips *string `json:"Clips,omitempty" xml:"Clips,omitempty"`
	// The live stream configurations, in the JSON format. The configurations must include the following parameters:
	//
	// 	- AppName: the name of the application to which the live stream belongs.
	//
	// 	- DomainName: the domain name of the application.
	//
	// 	- StreamName: the name of the live stream.
	//
	// example:
	//
	// { "AppName": "app", "DomainName": "domain.com", "StreamName": "stream"  }
	LiveStreamConfig *string `json:"LiveStreamConfig,omitempty" xml:"LiveStreamConfig,omitempty"`
	// The production configurations, in the JSON format. Mode specifies the editing mode. Valid values:
	//
	// 	- **AccurateFast*	- (default): fast editing. It is faster than the Accurate mode. The resolution of the output file is the same as that of the source stream. You cannot specify the width and height of the output file.
	//
	// 	- **Accurate**: accurate editing. In this mode, you can specify the width and height of the output file.
	//
	// 	- **Rough**: rough editing. The minimum precision is one TS segment. The output file comprises all segments within the specified time range. You can specify the width and height of the output file.
	//
	// 	- **RoughFast**: fast rough editing. It is faster than the Accurate mode. The minimum precision is one TS segment. The output file comprises all segments within the specified time range. The resolution of the output file is the same as that of the source stream. You cannot specify the width and height of the output file.
	//
	// example:
	//
	// { "Mode": "AccurateFast"}
	MediaProduceConfig *string `json:"MediaProduceConfig,omitempty" xml:"MediaProduceConfig,omitempty"`
	// The configurations of the output file, in the JSON format. You can specify an OSS URL or a storage location in a storage bucket of ApsaraVideo VOD.
	//
	// 	- To store the output file in OSS, you must specify MediaURL.
	//
	// 	- To store the output file in ApsaraVideo VOD, you must specify StorageLocation and FileName.
	OutputMediaConfig *string `json:"OutputMediaConfig,omitempty" xml:"OutputMediaConfig,omitempty"`
	// The type of the output file. Valid values:
	//
	// 	- oss-object: OSS object in an OSS bucket.
	//
	// 	- vod-media: media asset in Alibaba Cloud VOD.
	//
	// example:
	//
	// oss-object
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
	// The ID of the live editing project. If this parameter is specified, the system reads the storage configurations of the project. If this parameter is not specified, the specified storage configurations take precedence.
	//
	// example:
	//
	// ****fddd7748b58bf1d47e95****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The user-defined data in the JSON format, which can be up to 512 bytes in length.
	//
	// example:
	//
	// {"key": "value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitLiveEditingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveEditingJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitLiveEditingJobRequest) GetClips() *string {
	return s.Clips
}

func (s *SubmitLiveEditingJobRequest) GetLiveStreamConfig() *string {
	return s.LiveStreamConfig
}

func (s *SubmitLiveEditingJobRequest) GetMediaProduceConfig() *string {
	return s.MediaProduceConfig
}

func (s *SubmitLiveEditingJobRequest) GetOutputMediaConfig() *string {
	return s.OutputMediaConfig
}

func (s *SubmitLiveEditingJobRequest) GetOutputMediaTarget() *string {
	return s.OutputMediaTarget
}

func (s *SubmitLiveEditingJobRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *SubmitLiveEditingJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitLiveEditingJobRequest) SetClips(v string) *SubmitLiveEditingJobRequest {
	s.Clips = &v
	return s
}

func (s *SubmitLiveEditingJobRequest) SetLiveStreamConfig(v string) *SubmitLiveEditingJobRequest {
	s.LiveStreamConfig = &v
	return s
}

func (s *SubmitLiveEditingJobRequest) SetMediaProduceConfig(v string) *SubmitLiveEditingJobRequest {
	s.MediaProduceConfig = &v
	return s
}

func (s *SubmitLiveEditingJobRequest) SetOutputMediaConfig(v string) *SubmitLiveEditingJobRequest {
	s.OutputMediaConfig = &v
	return s
}

func (s *SubmitLiveEditingJobRequest) SetOutputMediaTarget(v string) *SubmitLiveEditingJobRequest {
	s.OutputMediaTarget = &v
	return s
}

func (s *SubmitLiveEditingJobRequest) SetProjectId(v string) *SubmitLiveEditingJobRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitLiveEditingJobRequest) SetUserData(v string) *SubmitLiveEditingJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitLiveEditingJobRequest) Validate() error {
	return dara.Validate(s)
}
