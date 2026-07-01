// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceExtractJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputShrink(v string) *SubmitTraceExtractJobShrinkRequest
	GetInputShrink() *string
	SetParams(v string) *SubmitTraceExtractJobShrinkRequest
	GetParams() *string
	SetUserData(v string) *SubmitTraceExtractJobShrinkRequest
	GetUserData() *string
}

type SubmitTraceExtractJobShrinkRequest struct {
	// The input video from which to extract the watermark.
	//
	// > - The OSS object or media asset must be in the same region as your IMS service.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// Extraction job parameters, specified as a JSON string. The following parameters are supported:
	//
	// - `m3u8Type`: The algorithm type. The default value is `v1`.
	//
	//   - `v1`: Extracts an m3u8 playlist with absolute paths.
	//
	//   - `v2`: Extracts an m3u8 playlist with relative paths.
	//
	// example:
	//
	// {"m3u8Type":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The user-defined data. Maximum length: 1,024 bytes.
	//
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitTraceExtractJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceExtractJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitTraceExtractJobShrinkRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitTraceExtractJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTraceExtractJobShrinkRequest) SetInputShrink(v string) *SubmitTraceExtractJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitTraceExtractJobShrinkRequest) SetParams(v string) *SubmitTraceExtractJobShrinkRequest {
	s.Params = &v
	return s
}

func (s *SubmitTraceExtractJobShrinkRequest) SetUserData(v string) *SubmitTraceExtractJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTraceExtractJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
