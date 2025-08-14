// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoCognitionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeResults(v *QueryVideoCognitionJobRequestIncludeResults) *QueryVideoCognitionJobRequest
	GetIncludeResults() *QueryVideoCognitionJobRequestIncludeResults
	SetJobId(v string) *QueryVideoCognitionJobRequest
	GetJobId() *string
	SetParams(v string) *QueryVideoCognitionJobRequest
	GetParams() *string
}

type QueryVideoCognitionJobRequest struct {
	IncludeResults *QueryVideoCognitionJobRequestIncludeResults `json:"IncludeResults,omitempty" xml:"IncludeResults,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s QueryVideoCognitionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoCognitionJobRequest) GoString() string {
	return s.String()
}

func (s *QueryVideoCognitionJobRequest) GetIncludeResults() *QueryVideoCognitionJobRequestIncludeResults {
	return s.IncludeResults
}

func (s *QueryVideoCognitionJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryVideoCognitionJobRequest) GetParams() *string {
	return s.Params
}

func (s *QueryVideoCognitionJobRequest) SetIncludeResults(v *QueryVideoCognitionJobRequestIncludeResults) *QueryVideoCognitionJobRequest {
	s.IncludeResults = v
	return s
}

func (s *QueryVideoCognitionJobRequest) SetJobId(v string) *QueryVideoCognitionJobRequest {
	s.JobId = &v
	return s
}

func (s *QueryVideoCognitionJobRequest) SetParams(v string) *QueryVideoCognitionJobRequest {
	s.Params = &v
	return s
}

func (s *QueryVideoCognitionJobRequest) Validate() error {
	return dara.Validate(s)
}

type QueryVideoCognitionJobRequestIncludeResults struct {
	// example:
	//
	// true
	NeedAsr *bool `json:"NeedAsr,omitempty" xml:"NeedAsr,omitempty"`
	// example:
	//
	// true
	NeedOcr *bool `json:"NeedOcr,omitempty" xml:"NeedOcr,omitempty"`
	// example:
	//
	// true
	NeedProcess *bool `json:"NeedProcess,omitempty" xml:"NeedProcess,omitempty"`
}

func (s QueryVideoCognitionJobRequestIncludeResults) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoCognitionJobRequestIncludeResults) GoString() string {
	return s.String()
}

func (s *QueryVideoCognitionJobRequestIncludeResults) GetNeedAsr() *bool {
	return s.NeedAsr
}

func (s *QueryVideoCognitionJobRequestIncludeResults) GetNeedOcr() *bool {
	return s.NeedOcr
}

func (s *QueryVideoCognitionJobRequestIncludeResults) GetNeedProcess() *bool {
	return s.NeedProcess
}

func (s *QueryVideoCognitionJobRequestIncludeResults) SetNeedAsr(v bool) *QueryVideoCognitionJobRequestIncludeResults {
	s.NeedAsr = &v
	return s
}

func (s *QueryVideoCognitionJobRequestIncludeResults) SetNeedOcr(v bool) *QueryVideoCognitionJobRequestIncludeResults {
	s.NeedOcr = &v
	return s
}

func (s *QueryVideoCognitionJobRequestIncludeResults) SetNeedProcess(v bool) *QueryVideoCognitionJobRequestIncludeResults {
	s.NeedProcess = &v
	return s
}

func (s *QueryVideoCognitionJobRequestIncludeResults) Validate() error {
	return dara.Validate(s)
}
