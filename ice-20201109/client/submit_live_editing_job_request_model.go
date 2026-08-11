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
	// The list of clip segments. The output is produced by concatenating the segments in the list in order. JSON Array.
	//
	// Each segment contains a start time and an end time. If no live stream parameters are specified, the outer-level live stream configuration is used. Both start and end timestamps are in UTC. For parameter details, see the Clip data structure below.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"StartTime\\": \\" 2021-06-21T08:01:00Z\\",  \\"EndTime\\": \\" 2021-06-21T08:03:00Z\\" ,  "AppName": "app", "DomainName": "domain.com", "StreamName": "stream"},  {\\"StartTime\\": \\" 2021-06-21T08:05:00Z\\",  \\"EndTime\\": \\" 2021-06-21T08:09:00Z\\" }]
	Clips *string `json:"Clips,omitempty" xml:"Clips,omitempty"`
	// The live stream configuration. JSON Object. The following configuration items are required:
	//
	// - AppName: the name of the application to which the stream belongs.
	//
	// - DomainName: the domain name.
	//
	// - StreamName: the name of the live stream.
	//
	// example:
	//
	// { "AppName": "app", "DomainName": "domain.com", "StreamName": "stream"  }
	LiveStreamConfig *string `json:"LiveStreamConfig,omitempty" xml:"LiveStreamConfig,omitempty"`
	// The composition configuration for generating segments, in JSON format. Mode specifies the editing mode. Valid values:
	//
	// - **AccurateFast*	- (default): fast accurate editing. This mode is faster than the Accurate mode. The output file resolution is the same as the source stream resolution. Custom output width and height are not supported.
	//
	// - **Accurate**: accurate editing. You can specify the output width and height.
	//
	// - **Rough**: rough editing. The minimum precision is one TS segment. The output contains all segments within the specified start and end time. You can specify the output width and height.
	//
	// - **RoughFast**: fast rough editing. This mode is faster than the Accurate mode. The minimum precision is one TS segment. The output contains all segments within the specified start and end time. The output file resolution is the same as the source stream resolution. Custom output width and height are not supported.
	//
	// example:
	//
	// { "Mode": "AccurateFast"}
	MediaProduceConfig *string `json:"MediaProduceConfig,omitempty" xml:"MediaProduceConfig,omitempty"`
	// The destination configuration for the output. JSON Object. You can specify the URL of the output on OSS or the storage location in a VOD bucket.
	//
	// - When outputting to OSS, the MediaURL of the output destination is required.
	//
	// - When outputting to VOD, the StorageLocation and FileName parameters are required.
	//
	// example:
	//
	// { "MediaURL": "https://ice-auto-test.oss-cn-shanghai.aliyuncs.com/testfile.mp4" }, or { "StorageLocation": "bucket.oss-cn-shanghai.aliyuncs.com", "FileName": "output.mp4" }
	OutputMediaConfig *string `json:"OutputMediaConfig,omitempty" xml:"OutputMediaConfig,omitempty"`
	// The target type of the output. Valid values:
	//
	// - oss-object: an OSS object in an Alibaba Cloud OSS bucket.
	//
	// - vod-media: a media asset in Alibaba Cloud VOD.
	//
	// example:
	//
	// oss-object
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
	// The ID of the live editing project. If this parameter is not empty, the storage configuration associated with the project is used. If this parameter is empty, the storage configuration specified in the request parameters is used.
	//
	// example:
	//
	// ****fddd7748b58bf1d47e95****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The custom settings. JSON Object. Maximum length: 512 bytes.
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
