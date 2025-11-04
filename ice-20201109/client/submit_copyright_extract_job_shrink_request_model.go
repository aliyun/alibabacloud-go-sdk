// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightExtractJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputShrink(v string) *SubmitCopyrightExtractJobShrinkRequest
	GetInputShrink() *string
	SetParams(v string) *SubmitCopyrightExtractJobShrinkRequest
	GetParams() *string
	SetUserData(v string) *SubmitCopyrightExtractJobShrinkRequest
	GetUserData() *string
}

type SubmitCopyrightExtractJobShrinkRequest struct {
	// The source video file from which to extract the watermark.
	//
	// > The OSS object or media asset must reside in the same region as the IMS service region.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// Additional parameters for the watermark job, provided as a JSON string. Supported parameter:
	//
	// 	- algoType: The algorithm type. Defaults to v1. The extraction algorithm must match the one used for embedding.
	//
	//     	- v1: Copyright watermark extraction algorithm for long videos.
	//
	//     	- v2: Copyright watermark extraction algorithm for short videos.
	//
	// example:
	//
	// {"algoType":"v2"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The custom data, which can be up to 1,024 bytes in size.
	//
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitCopyrightExtractJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightExtractJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitCopyrightExtractJobShrinkRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitCopyrightExtractJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitCopyrightExtractJobShrinkRequest) SetInputShrink(v string) *SubmitCopyrightExtractJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitCopyrightExtractJobShrinkRequest) SetParams(v string) *SubmitCopyrightExtractJobShrinkRequest {
	s.Params = &v
	return s
}

func (s *SubmitCopyrightExtractJobShrinkRequest) SetUserData(v string) *SubmitCopyrightExtractJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitCopyrightExtractJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
