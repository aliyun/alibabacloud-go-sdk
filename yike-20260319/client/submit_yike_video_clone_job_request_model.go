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
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "JobParams": "{\\"SceneType\\":\\"variant-clone\\",\\"OriginalVideo\\":{\\"MediaId\\":\\"1d342ee****c71f18000e7f6d45b6302\\"},\\"SceneConfig\\":\\xxxxxxxxx\\",\\"Resolution\\":\\"720P\\",\\"AvatarData\\":{\\"AvatarPortrait\\":\\"https://example-bucket.oss-cn-shanghai.aliyuncs.com/sucai/videoremake/0518/shuziren-005.png\\",\\"AvatarVoice\\":\\"xxxxxx\\"},\\"UserMaterials\\":[{\\"MediaId\\":\\"e3785e10****71f1bfc9e7f6c6586301\\"},{\\"MediaId\\":\\"e34196a05****1f1bfb0f6f7d44b6301\\"},{\\"MediaId\\":\\"e3628c205****1f1bfc5f7f6d4496301\\"},{\\"MediaId\\":\\"e35e1f505****1f1bfc9e7f6c6586301\\"},{\\"MediaId\\":\\"e37bb9705****1f18000e7f6d45b6301\\"}],\\"WithSubtitles\\":true,\\"VoiceDuration\\":15}",
	//
	//   "UserData": "{\\"newsKey\\":\\"NEWS_20260420_001\\",\\"key1\\":\\"value1\\", \\"NotifyAddress\\":\\"https://example.com/callback/video-task\\"}"
	//
	// }
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
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
