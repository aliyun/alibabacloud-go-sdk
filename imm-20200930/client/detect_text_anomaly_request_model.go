// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectTextAnomalyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DetectTextAnomalyRequest
	GetContent() *string
	SetProjectName(v string) *DetectTextAnomalyRequest
	GetProjectName() *string
}

type DetectTextAnomalyRequest struct {
	// The text to be detected. It can contain up to 10,000 characters (including punctuation marks). Only Chinese text can be detected.
	//
	// This parameter is required.
	//
	// example:
	//
	// content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DetectTextAnomalyRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectTextAnomalyRequest) GoString() string {
	return s.String()
}

func (s *DetectTextAnomalyRequest) GetContent() *string {
	return s.Content
}

func (s *DetectTextAnomalyRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectTextAnomalyRequest) SetContent(v string) *DetectTextAnomalyRequest {
	s.Content = &v
	return s
}

func (s *DetectTextAnomalyRequest) SetProjectName(v string) *DetectTextAnomalyRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectTextAnomalyRequest) Validate() error {
	return dara.Validate(s)
}
