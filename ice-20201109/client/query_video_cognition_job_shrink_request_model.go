// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoCognitionJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeResultsShrink(v string) *QueryVideoCognitionJobShrinkRequest
	GetIncludeResultsShrink() *string
	SetJobId(v string) *QueryVideoCognitionJobShrinkRequest
	GetJobId() *string
	SetParams(v string) *QueryVideoCognitionJobShrinkRequest
	GetParams() *string
}

type QueryVideoCognitionJobShrinkRequest struct {
	// Specifies whether to include the full algorithm results in the response.
	IncludeResultsShrink *string `json:"IncludeResults,omitempty" xml:"IncludeResults,omitempty"`
	// The ID of the task to query. It is returned when you call the [SubmitSmarttagJob](https://help.aliyun.com/document_detail/478786.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Additional request parameters, provided as a JSON string.
	//
	// example:
	//
	// {}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s QueryVideoCognitionJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoCognitionJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryVideoCognitionJobShrinkRequest) GetIncludeResultsShrink() *string {
	return s.IncludeResultsShrink
}

func (s *QueryVideoCognitionJobShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryVideoCognitionJobShrinkRequest) GetParams() *string {
	return s.Params
}

func (s *QueryVideoCognitionJobShrinkRequest) SetIncludeResultsShrink(v string) *QueryVideoCognitionJobShrinkRequest {
	s.IncludeResultsShrink = &v
	return s
}

func (s *QueryVideoCognitionJobShrinkRequest) SetJobId(v string) *QueryVideoCognitionJobShrinkRequest {
	s.JobId = &v
	return s
}

func (s *QueryVideoCognitionJobShrinkRequest) SetParams(v string) *QueryVideoCognitionJobShrinkRequest {
	s.Params = &v
	return s
}

func (s *QueryVideoCognitionJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
