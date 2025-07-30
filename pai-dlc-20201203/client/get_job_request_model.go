// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNeedDetail(v bool) *GetJobRequest
	GetNeedDetail() *bool
}

type GetJobRequest struct {
	// Specifies whether to return the job details. Default value: true.
	//
	// example:
	//
	// true
	NeedDetail *bool `json:"NeedDetail,omitempty" xml:"NeedDetail,omitempty"`
}

func (s GetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) GetNeedDetail() *bool {
	return s.NeedDetail
}

func (s *GetJobRequest) SetNeedDetail(v bool) *GetJobRequest {
	s.NeedDetail = &v
	return s
}

func (s *GetJobRequest) Validate() error {
	return dara.Validate(s)
}
