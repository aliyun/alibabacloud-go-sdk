// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaComprehensionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobParams(v string) *SubmitMediaComprehensionJobRequest
	GetJobParams() *string
	SetUserData(v string) *SubmitMediaComprehensionJobRequest
	GetUserData() *string
}

type SubmitMediaComprehensionJobRequest struct {
	// example:
	//
	// {"ProductName":"Quiet Blender Soymilk Maker","BrandName":"LiangChu","SellingPoints":["Low-noise blending","One-touch self-cleaning"]}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
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

func (s *SubmitMediaComprehensionJobRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitMediaComprehensionJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaComprehensionJobRequest) SetJobParams(v string) *SubmitMediaComprehensionJobRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitMediaComprehensionJobRequest) SetUserData(v string) *SubmitMediaComprehensionJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaComprehensionJobRequest) Validate() error {
	return dara.Validate(s)
}
