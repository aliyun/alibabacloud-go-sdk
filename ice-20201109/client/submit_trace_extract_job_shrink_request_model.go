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
	// The source video file from which to extract the watermark.
	//
	// > The OSS object or media asset must reside in the same region as the IMS service region.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// Additional parameters for the watermark job, provided as a JSON string. Supported parameter:
	//
	// 	- m3u8Type: The extraction algorithm type. Defaults to v1.
	//
	//     	- v1: Extracts from an M3U8 with absolute paths.
	//
	//     	- v2: Extracts from an M3U8 with relative paths.
	//
	// example:
	//
	// {"m3u8Type":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The custom data, which can be up to 1,024 bytes in size.
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
