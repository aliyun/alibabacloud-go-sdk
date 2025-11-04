// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSegmentationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitSegmentationJobRequest
	GetClientToken() *string
	SetInputConfig(v string) *SubmitSegmentationJobRequest
	GetInputConfig() *string
	SetJobParams(v string) *SubmitSegmentationJobRequest
	GetJobParams() *string
	SetOutputConfig(v string) *SubmitSegmentationJobRequest
	GetOutputConfig() *string
	SetUserData(v string) *SubmitSegmentationJobRequest
	GetUserData() *string
}

type SubmitSegmentationJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The input configuration. For detailed parameters, see [InputConfig](~~2874121#cc59ad3082jbx~~).
	//
	// example:
	//
	// {
	//
	// 	"Type": "OSS",
	//
	// 	"Media": "http://test-bucket.oss-cn-shanghai.aliyuncs.com/test.mp4"
	//
	// }
	//
	// or {
	//
	// 	"Type": "Media",
	//
	// 	"Media": "ce49a020e****1ef81c1e6f6d5686302"
	//
	// }
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The task parameters. For details, see [JobParams](~~2874121#a60357f2d5iix~~).
	//
	// example:
	//
	// {
	//
	// 	"Mode": "UserDefined",
	//
	// 	"Ranges": [{
	//
	// 		"In": 10,
	//
	// 		"Out": 20
	//
	// 	}, {
	//
	// 		"In": 35,
	//
	// 		"Out": 50
	//
	// 	}]
	//
	// }
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// The output configuration. For detailed parameters, see [OutputConfig](~~2874121#cef23186a8d6w~~).
	//
	// example:
	//
	// {
	//
	// 	"OutputMediaTarget": "oss-object",
	//
	// 	"Bucket": "test-bucket",
	//
	// 	"ObjectKey": "path/to/test_{index}.mp4",
	//
	// 	"Width": 1920,
	//
	// 	"Height": 1080,
	//
	// 	"ExportAsNewMedia": false
	//
	// }
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// The user-defined data in the JSON format, which can be up to 512 bytes in length.
	//
	// example:
	//
	// {"test": "22"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSegmentationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSegmentationJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSegmentationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitSegmentationJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitSegmentationJobRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitSegmentationJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitSegmentationJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSegmentationJobRequest) SetClientToken(v string) *SubmitSegmentationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitSegmentationJobRequest) SetInputConfig(v string) *SubmitSegmentationJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitSegmentationJobRequest) SetJobParams(v string) *SubmitSegmentationJobRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitSegmentationJobRequest) SetOutputConfig(v string) *SubmitSegmentationJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitSegmentationJobRequest) SetUserData(v string) *SubmitSegmentationJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSegmentationJobRequest) Validate() error {
	return dara.Validate(s)
}
