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
	// A JSON array that specifies the clips to edit. The job creates the output file by concatenating these clips in the specified order.
	//
	// Each clip includes a start and end time. If live stream parameters are not specified for a clip, the system uses the global `LiveStreamConfig` settings. The start and end timestamps must be in UTC. For more details, see the Clip data structure below.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"StartTime\\": \\" 2021-06-21T08:01:00Z\\",  \\"EndTime\\": \\" 2021-06-21T08:03:00Z\\" ,  "AppName": "app", "DomainName": "domain.com", "StreamName": "stream"},  {\\"StartTime\\": \\" 2021-06-21T08:05:00Z\\",  \\"EndTime\\": \\" 2021-06-21T08:09:00Z\\" }]
	Clips *string `json:"Clips,omitempty" xml:"Clips,omitempty"`
	// The configuration of the source live stream, specified as a JSON object. It includes the following parameters:
	//
	// - `AppName`: The name of the application to which the stream belongs.
	//
	// - `DomainName`: The domain name of the stream.
	//
	// - `StreamName`: The name of the live stream.
	//
	// example:
	//
	// { "AppName": "app", "DomainName": "domain.com", "StreamName": "stream"  }
	LiveStreamConfig *string `json:"LiveStreamConfig,omitempty" xml:"LiveStreamConfig,omitempty"`
	// The production configuration for the output file, specified as a JSON object. The `Mode` parameter specifies the editing mode. Valid values are:
	//
	// - **AccurateFast*	- (Default): Fast and precise editing. It offers faster processing compared to the `Accurate` mode. The output file has the same resolution as the source stream. You cannot specify a custom width and height for the output file.
	//
	// - **Accurate**: Precise editing. This mode lets you specify a custom width and height for the output file.
	//
	// - **Rough**: Rough editing with a precision of a single TS segment. The output file includes all segments between the specified start and end times. You can specify a custom width and height for the output file.
	//
	// - **RoughFast**: Fast rough-cut editing, which is faster than the `Accurate` mode. It has a precision of a single TS segment, and the output file includes all segments between the specified start and end times. The output file has the same resolution as the source stream. You cannot specify a custom width and height for the output file.
	//
	// example:
	//
	// { "Mode": "AccurateFast"}
	MediaProduceConfig *string `json:"MediaProduceConfig,omitempty" xml:"MediaProduceConfig,omitempty"`
	// The destination configuration for the output file, specified as a JSON object. You can specify either a URL on OSS or a storage location in a VOD bucket.
	//
	// - To output to OSS, the `MediaURL` parameter is required.
	//
	// - To output to VOD, the `StorageLocation` and `FileName` parameters are required.
	//
	// example:
	//
	// { "MediaURL": "https://ice-auto-test.oss-cn-shanghai.aliyuncs.com/testfile.mp4" }, or { "StorageLocation": "bucket.oss-cn-shanghai.aliyuncs.com", "FileName": "output.mp4" }
	OutputMediaConfig *string `json:"OutputMediaConfig,omitempty" xml:"OutputMediaConfig,omitempty"`
	// The destination type for the output file. Valid values:
	//
	// - `oss-object`: An object in an Alibaba Cloud OSS bucket.
	//
	// - `vod-media`: A media asset in Alibaba Cloud VOD.
	//
	// example:
	//
	// oss-object
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
	// The ID of the live editing project. If you specify this parameter, the system uses the storage settings from the project. If left empty, the system uses the storage settings provided in the request instead.
	//
	// example:
	//
	// ****fddd7748b58bf1d47e95****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Custom user data, provided as a JSON object. The maximum length is 512 bytes.
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
