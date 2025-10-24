// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightExtractJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallBack(v string) *SubmitCopyrightExtractJobRequest
	GetCallBack() *string
	SetInput(v string) *SubmitCopyrightExtractJobRequest
	GetInput() *string
	SetParams(v string) *SubmitCopyrightExtractJobRequest
	GetParams() *string
	SetUrl(v string) *SubmitCopyrightExtractJobRequest
	GetUrl() *string
	SetUserData(v string) *SubmitCopyrightExtractJobRequest
	GetUserData() *string
}

type SubmitCopyrightExtractJobRequest struct {
	// example:
	//
	// http://example.com/callback
	CallBack *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	// example:
	//
	// {"Bucket":"example","Location":"oss-cn-shanghai","Object":"example.mp4"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// {"algoType":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// http://www.example.com/video/test.mp4
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitCopyrightExtractJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractJobRequest) GetCallBack() *string {
	return s.CallBack
}

func (s *SubmitCopyrightExtractJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitCopyrightExtractJobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitCopyrightExtractJobRequest) GetUrl() *string {
	return s.Url
}

func (s *SubmitCopyrightExtractJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitCopyrightExtractJobRequest) SetCallBack(v string) *SubmitCopyrightExtractJobRequest {
	s.CallBack = &v
	return s
}

func (s *SubmitCopyrightExtractJobRequest) SetInput(v string) *SubmitCopyrightExtractJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitCopyrightExtractJobRequest) SetParams(v string) *SubmitCopyrightExtractJobRequest {
	s.Params = &v
	return s
}

func (s *SubmitCopyrightExtractJobRequest) SetUrl(v string) *SubmitCopyrightExtractJobRequest {
	s.Url = &v
	return s
}

func (s *SubmitCopyrightExtractJobRequest) SetUserData(v string) *SubmitCopyrightExtractJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitCopyrightExtractJobRequest) Validate() error {
	return dara.Validate(s)
}
