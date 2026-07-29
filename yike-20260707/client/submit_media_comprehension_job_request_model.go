// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaComprehensionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v string) *SubmitMediaComprehensionJobRequest
	GetInput() *string
	SetJobParams(v string) *SubmitMediaComprehensionJobRequest
	GetJobParams() *string
	SetJobType(v string) *SubmitMediaComprehensionJobRequest
	GetJobType() *string
	SetUserData(v string) *SubmitMediaComprehensionJobRequest
	GetUserData() *string
}

type SubmitMediaComprehensionJobRequest struct {
	// example:
	//
	// {"Medias":[{"Type":"video","Url":"https://xxx.mp4"}]}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// {"ProductName":"Quiet Blender Soymilk Maker","BrandName":"LiangChu","SellingPoints":["Low-noise blending","One-touch self-cleaning"]}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// example:
	//
	// VideoBreakdown
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// {"NotifyAddress": "http://xxx.callback.url"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaComprehensionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaComprehensionJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaComprehensionJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitMediaComprehensionJobRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitMediaComprehensionJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *SubmitMediaComprehensionJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaComprehensionJobRequest) SetInput(v string) *SubmitMediaComprehensionJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitMediaComprehensionJobRequest) SetJobParams(v string) *SubmitMediaComprehensionJobRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitMediaComprehensionJobRequest) SetJobType(v string) *SubmitMediaComprehensionJobRequest {
	s.JobType = &v
	return s
}

func (s *SubmitMediaComprehensionJobRequest) SetUserData(v string) *SubmitMediaComprehensionJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaComprehensionJobRequest) Validate() error {
	return dara.Validate(s)
}
