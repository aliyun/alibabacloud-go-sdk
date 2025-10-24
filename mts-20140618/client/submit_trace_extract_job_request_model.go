// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceExtractJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallBack(v string) *SubmitTraceExtractJobRequest
	GetCallBack() *string
	SetInput(v string) *SubmitTraceExtractJobRequest
	GetInput() *string
	SetParams(v string) *SubmitTraceExtractJobRequest
	GetParams() *string
	SetUrl(v string) *SubmitTraceExtractJobRequest
	GetUrl() *string
	SetUserData(v string) *SubmitTraceExtractJobRequest
	GetUserData() *string
}

type SubmitTraceExtractJobRequest struct {
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
	// {"m3u8Type":"v1"}
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

func (s SubmitTraceExtractJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractJobRequest) GetCallBack() *string {
	return s.CallBack
}

func (s *SubmitTraceExtractJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitTraceExtractJobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitTraceExtractJobRequest) GetUrl() *string {
	return s.Url
}

func (s *SubmitTraceExtractJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTraceExtractJobRequest) SetCallBack(v string) *SubmitTraceExtractJobRequest {
	s.CallBack = &v
	return s
}

func (s *SubmitTraceExtractJobRequest) SetInput(v string) *SubmitTraceExtractJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitTraceExtractJobRequest) SetParams(v string) *SubmitTraceExtractJobRequest {
	s.Params = &v
	return s
}

func (s *SubmitTraceExtractJobRequest) SetUrl(v string) *SubmitTraceExtractJobRequest {
	s.Url = &v
	return s
}

func (s *SubmitTraceExtractJobRequest) SetUserData(v string) *SubmitTraceExtractJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTraceExtractJobRequest) Validate() error {
	return dara.Validate(s)
}
