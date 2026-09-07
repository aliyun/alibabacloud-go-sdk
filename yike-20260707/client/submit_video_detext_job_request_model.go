// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoDetextJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitVideoDetextJobRequest
	GetClientToken() *string
	SetInput(v string) *SubmitVideoDetextJobRequest
	GetInput() *string
	SetJobParameters(v string) *SubmitVideoDetextJobRequest
	GetJobParameters() *string
	SetOutput(v string) *SubmitVideoDetextJobRequest
	GetOutput() *string
	SetUserData(v string) *SubmitVideoDetextJobRequest
	GetUserData() *string
}

type SubmitVideoDetextJobRequest struct {
	// The user-level idempotency token. The token can be up to 40 characters in length. If the same user submits a request with the same token, the original task is returned.
	//
	// example:
	//
	// detext-client-20260820-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The input configuration JSON string. You must specify either VideoUrl or VideoMediaId, but not both.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"VideoMediaId":"media-video-001"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The text erasure parameter JSON string. This parameter can contain EraseAllText, TimeRanges, TextTargets, FullEraseTargets, and Config.
	//
	// example:
	//
	// {"EraseAllText":false,"TextTargets":[{"Box":[0.1,0.8,0.8,0.15],"TimeRanges":[[0,30]]}]}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// The output configuration JSON string. You can use OssUri to specify the customer\\"s OSS bucket. If a directory is specified, the output file is named detext.mp4.
	//
	// example:
	//
	// {"OssUri":"oss://example-bucket/video-detext/job-001/"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The custom user data JSON string. This parameter can contain the asynchronous notification address NotifyAddress.
	//
	// example:
	//
	// {"NotifyAddress":"mns://account.mns.cn-shanghai.aliyuncs.com/queues/detext-result"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitVideoDetextJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoDetextJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoDetextJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitVideoDetextJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitVideoDetextJobRequest) GetJobParameters() *string {
	return s.JobParameters
}

func (s *SubmitVideoDetextJobRequest) GetOutput() *string {
	return s.Output
}

func (s *SubmitVideoDetextJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitVideoDetextJobRequest) SetClientToken(v string) *SubmitVideoDetextJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitVideoDetextJobRequest) SetInput(v string) *SubmitVideoDetextJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitVideoDetextJobRequest) SetJobParameters(v string) *SubmitVideoDetextJobRequest {
	s.JobParameters = &v
	return s
}

func (s *SubmitVideoDetextJobRequest) SetOutput(v string) *SubmitVideoDetextJobRequest {
	s.Output = &v
	return s
}

func (s *SubmitVideoDetextJobRequest) SetUserData(v string) *SubmitVideoDetextJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitVideoDetextJobRequest) Validate() error {
	return dara.Validate(s)
}
