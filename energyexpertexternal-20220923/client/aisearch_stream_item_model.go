// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchStreamItem interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *AISearchStreamItem
	GetContent() *string
	SetParams(v interface{}) *AISearchStreamItem
	GetParams() interface{}
	SetType(v string) *AISearchStreamItem
	GetType() *string
}

type AISearchStreamItem struct {
	// example:
	//
	// “network problem can be”
	Content *string     `json:"content,omitempty" xml:"content,omitempty"`
	Params  interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AISearchStreamItem) String() string {
	return dara.Prettify(s)
}

func (s AISearchStreamItem) GoString() string {
	return s.String()
}

func (s *AISearchStreamItem) GetContent() *string {
	return s.Content
}

func (s *AISearchStreamItem) GetParams() interface{} {
	return s.Params
}

func (s *AISearchStreamItem) GetType() *string {
	return s.Type
}

func (s *AISearchStreamItem) SetContent(v string) *AISearchStreamItem {
	s.Content = &v
	return s
}

func (s *AISearchStreamItem) SetParams(v interface{}) *AISearchStreamItem {
	s.Params = v
	return s
}

func (s *AISearchStreamItem) SetType(v string) *AISearchStreamItem {
	s.Type = &v
	return s
}

func (s *AISearchStreamItem) Validate() error {
	return dara.Validate(s)
}
