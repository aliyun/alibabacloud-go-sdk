// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateExperimentContentRequest
	GetContent() *string
	SetVersion(v int64) *UpdateExperimentContentRequest
	GetVersion() *int64
}

type UpdateExperimentContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {     "metadata": {       "name": "实验名称",       "id": "pai_exp_xxxdfafafasfa",       "desc": "实验描述",     },     "nodes": [     ],     "edges": [     ],     "globalParams": [     ],     "globalSettings":[     ]  }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateExperimentContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentContentRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateExperimentContentRequest) GetVersion() *int64 {
	return s.Version
}

func (s *UpdateExperimentContentRequest) SetContent(v string) *UpdateExperimentContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateExperimentContentRequest) SetVersion(v int64) *UpdateExperimentContentRequest {
	s.Version = &v
	return s
}

func (s *UpdateExperimentContentRequest) Validate() error {
	return dara.Validate(s)
}
