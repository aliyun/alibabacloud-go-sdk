// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWorkflowGlobalEnvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SaveWorkflowGlobalEnvRequest
	GetContent() *string
}

type SaveWorkflowGlobalEnvRequest struct {
	// The configuration content.
	//
	// This parameter is required.
	//
	// example:
	//
	// MODE=demo
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s SaveWorkflowGlobalEnvRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveWorkflowGlobalEnvRequest) GoString() string {
	return s.String()
}

func (s *SaveWorkflowGlobalEnvRequest) GetContent() *string {
	return s.Content
}

func (s *SaveWorkflowGlobalEnvRequest) SetContent(v string) *SaveWorkflowGlobalEnvRequest {
	s.Content = &v
	return s
}

func (s *SaveWorkflowGlobalEnvRequest) Validate() error {
	return dara.Validate(s)
}
