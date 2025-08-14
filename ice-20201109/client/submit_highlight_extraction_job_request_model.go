// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHighlightExtractionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitHighlightExtractionJobRequest
	GetClientToken() *string
	SetInputConfig(v string) *SubmitHighlightExtractionJobRequest
	GetInputConfig() *string
	SetOutputConfig(v string) *SubmitHighlightExtractionJobRequest
	GetOutputConfig() *string
	SetUserData(v string) *SubmitHighlightExtractionJobRequest
	GetUserData() *string
}

type SubmitHighlightExtractionJobRequest struct {
	// The client token used to ensure the idempotency of the request.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The input configuration.
	//
	// example:
	//
	// {
	//
	// 	"MediaArray": [{
	//
	// 		"MediaId": "ceb72f00e****1ef8216e7e6c64a6302"
	//
	// 	}, {
	//
	// 		"MediaId": "ce450c40e****1ef8216e7e6c64a6302"
	//
	// 	}, {
	//
	// 		"MediaId": "ce49a020e****1ef81c1e6f6d5686302"
	//
	// 	}, {
	//
	// 		"MediaId": "d047e120e****1ef81c1e6f6d5686302"
	//
	// 	}, {
	//
	// 		"MediaId": "cfe2ddc0e****1ef81c1e6f6d5686302"
	//
	// 	}],
	//
	// 	"Strategy": {
	//
	// 		"Count": 5,
	//
	// 		"ClipDuration": 15
	//
	// 	}
	//
	// }
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The output configuration.
	//
	// example:
	//
	// {
	//
	// 	"NeedExport": true,
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
	// The user-defined data, including the business and callback configurations. For more information, see [UserData](~~357745#section-urj-v3f-0s1~~).
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitHighlightExtractionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitHighlightExtractionJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitHighlightExtractionJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitHighlightExtractionJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitHighlightExtractionJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitHighlightExtractionJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitHighlightExtractionJobRequest) SetClientToken(v string) *SubmitHighlightExtractionJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitHighlightExtractionJobRequest) SetInputConfig(v string) *SubmitHighlightExtractionJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitHighlightExtractionJobRequest) SetOutputConfig(v string) *SubmitHighlightExtractionJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitHighlightExtractionJobRequest) SetUserData(v string) *SubmitHighlightExtractionJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitHighlightExtractionJobRequest) Validate() error {
	return dara.Validate(s)
}
