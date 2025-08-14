// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaConvertJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitMediaConvertJobRequest
	GetClientToken() *string
	SetConfig(v string) *SubmitMediaConvertJobRequest
	GetConfig() *string
	SetPipelineId(v string) *SubmitMediaConvertJobRequest
	GetPipelineId() *string
	SetUserData(v string) *SubmitMediaConvertJobRequest
	GetUserData() *string
}

type SubmitMediaConvertJobRequest struct {
	// The idempotency key that is used to ensure repeated requests have the same effect as a single request.
	//
	// example:
	//
	// 86f8e525-9d73-4dac-88aa-7aa4e950c00a
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configurations of the transcoding task.
	//
	// This parameter is required.
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the queue.
	//
	// example:
	//
	// e197ecfb103e4849922b054d3032f954
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The user data.
	//
	// example:
	//
	// {"videoId":"abcd"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaConvertJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaConvertJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaConvertJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitMediaConvertJobRequest) GetConfig() *string {
	return s.Config
}

func (s *SubmitMediaConvertJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitMediaConvertJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaConvertJobRequest) SetClientToken(v string) *SubmitMediaConvertJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitMediaConvertJobRequest) SetConfig(v string) *SubmitMediaConvertJobRequest {
	s.Config = &v
	return s
}

func (s *SubmitMediaConvertJobRequest) SetPipelineId(v string) *SubmitMediaConvertJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitMediaConvertJobRequest) SetUserData(v string) *SubmitMediaConvertJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaConvertJobRequest) Validate() error {
	return dara.Validate(s)
}
