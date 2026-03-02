// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppInfoContent interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *NeuronAppInfoContent
	GetContent() *string
	SetTitle(v string) *NeuronAppInfoContent
	GetTitle() *string
}

type NeuronAppInfoContent struct {
	// This parameter is required.
	//
	// example:
	//
	// order
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1343424
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s NeuronAppInfoContent) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppInfoContent) GoString() string {
	return s.String()
}

func (s *NeuronAppInfoContent) GetContent() *string {
	return s.Content
}

func (s *NeuronAppInfoContent) GetTitle() *string {
	return s.Title
}

func (s *NeuronAppInfoContent) SetContent(v string) *NeuronAppInfoContent {
	s.Content = &v
	return s
}

func (s *NeuronAppInfoContent) SetTitle(v string) *NeuronAppInfoContent {
	s.Title = &v
	return s
}

func (s *NeuronAppInfoContent) Validate() error {
	return dara.Validate(s)
}
