// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGenerationSource(v string) *CountTextRequest
	GetGenerationSource() *string
	SetIndustry(v string) *CountTextRequest
	GetIndustry() *string
	SetPublishStatus(v string) *CountTextRequest
	GetPublishStatus() *string
	SetStyle(v string) *CountTextRequest
	GetStyle() *string
}

type CountTextRequest struct {
	GenerationSource *string `json:"generationSource,omitempty" xml:"generationSource,omitempty"`
	Industry         *string `json:"industry,omitempty" xml:"industry,omitempty"`
	PublishStatus    *string `json:"publishStatus,omitempty" xml:"publishStatus,omitempty"`
	Style            *string `json:"style,omitempty" xml:"style,omitempty"`
}

func (s CountTextRequest) String() string {
	return dara.Prettify(s)
}

func (s CountTextRequest) GoString() string {
	return s.String()
}

func (s *CountTextRequest) GetGenerationSource() *string {
	return s.GenerationSource
}

func (s *CountTextRequest) GetIndustry() *string {
	return s.Industry
}

func (s *CountTextRequest) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *CountTextRequest) GetStyle() *string {
	return s.Style
}

func (s *CountTextRequest) SetGenerationSource(v string) *CountTextRequest {
	s.GenerationSource = &v
	return s
}

func (s *CountTextRequest) SetIndustry(v string) *CountTextRequest {
	s.Industry = &v
	return s
}

func (s *CountTextRequest) SetPublishStatus(v string) *CountTextRequest {
	s.PublishStatus = &v
	return s
}

func (s *CountTextRequest) SetStyle(v string) *CountTextRequest {
	s.Style = &v
	return s
}

func (s *CountTextRequest) Validate() error {
	return dara.Validate(s)
}
