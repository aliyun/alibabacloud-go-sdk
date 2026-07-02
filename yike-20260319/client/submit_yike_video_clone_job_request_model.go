// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeVideoCloneJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobParams(v string) *SubmitYikeVideoCloneJobRequest
	GetJobParams() *string
	SetUserData(v string) *SubmitYikeVideoCloneJobRequest
	GetUserData() *string
}

type SubmitYikeVideoCloneJobRequest struct {
	// The job request content. JSON string that contains the following parameters:
	//
	// - SceneType: string. The replication scene type. Valid values: `variant-clone`: full replication, applicable to same-category content rewriting scenarios where only partial elements (person/voice/image/text) are replaced.
	//
	// - OriginalVideo: object type that contains the following field: MediaId - the media asset ID (video uploaded to the platform).
	//
	// - SceneConfig: JSON string type. The scene extension parameters. For the variant-clone type, the value is `{"OldProductName":"xxx","ProductName":"xxx"}`.
	//
	// - UserMaterials: Array<Object> type. The list of user materials that contains the following field: MediaId - the media asset ID (image or video uploaded to the platform).
	//
	// - AvatarData: object type. The digital human information. AvatarPortrait: required, string, the portrait image URL. AvatarVoice: optional, string, the audio URL (used as a reference for audio replication).
	//
	// - Resolution: string type. The video resolution. Valid values: `720P`, `1080P`.
	//
	// - WithSubtitles: bool type. Specifies whether to include subtitles. Valid values: true: includes subtitles (default). false: does not include subtitles.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "JobParams": "{\\"SceneType\\":\\"variant-clone\\",\\"OriginalVideo\\":{\\"MediaId\\":\\"1d342ee****c71f18000e7f6d45b6302\\"},\\"SceneConfig\\":\\xxxxxxxxx\\",\\"Resolution\\":\\"720P\\",\\"AvatarData\\":{\\"AvatarPortrait\\":\\"https://example-bucket.oss-cn-shanghai.aliyuncs.com/sucai/videoremake/0518/shuziren-005.png\\",\\"AvatarVoice\\":\\"xxxxxx\\"},\\"UserMaterials\\":[{\\"MediaId\\":\\"e3785e10****71f1bfc9e7f6c6586301\\"}],\\"WithSubtitles\\":true}"
	//
	// }
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// The custom user parameter. JSON string. The callback result carries this value as-is (for example, newsKey).
	//
	// System reserved field NotifyAddress: the callback URL. The system sends a callback to this URL after the task is completed. Example: {"NotifyAddress": "http://xxx.callback.url"}
	//
	// example:
	//
	// {\\"newsKey\\":\\"NEWS_20260420_001\\",\\"key1\\":\\"value1\\", \\"NotifyAddress\\":\\"https://cms.example.com/callback/video-task\\"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitYikeVideoCloneJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeVideoCloneJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitYikeVideoCloneJobRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitYikeVideoCloneJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitYikeVideoCloneJobRequest) SetJobParams(v string) *SubmitYikeVideoCloneJobRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitYikeVideoCloneJobRequest) SetUserData(v string) *SubmitYikeVideoCloneJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitYikeVideoCloneJobRequest) Validate() error {
	return dara.Validate(s)
}
