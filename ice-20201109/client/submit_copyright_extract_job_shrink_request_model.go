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
	// The video file for watermark extraction.
	//
	// > - The region of the Object Storage Service (OSS) file or media asset must match the region of the current Intelligent Media Service (IMS) instance.
	//
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The watermark job parameters, specified as a JSON string.
	//
	// - algoType: The algorithm type. Default value: v1. The extraction algorithm type must match the algorithm type used for embedding the watermark.
	//
	//   - v1: The copyright extraction algorithm for long-form videos.
	//
	//   - v2: The copyright extraction algorithm for short-form videos.
	//
	// example:
	//
	// {"algoType":"v2"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The user-defined data. The maximum length is 1,024 bytes.
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
