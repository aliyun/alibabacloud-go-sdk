// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchMessageItem interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *AISearchMessageItem
	GetContent() *string
	SetParams(v interface{}) *AISearchMessageItem
	GetParams() interface{}
	SetType(v string) *AISearchMessageItem
	GetType() *string
}

type AISearchMessageItem struct {
	// example:
	//
	// "user：How to test my network speed?\\naisearch：To check your network speed, you can use..."
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// {
	//
	//  "miniapp_id": "test_miniapp",
	//
	//       "title": "testApp",
	//
	//       "version": "1.0.5",
	//
	//       "description": "description-mock",
	//
	//       "slogan": "slogan-mock",
	//
	//       "icon": "https://img.alicdn.com/test_icon.png",
	//
	//       "detail_desc": "detail-mock"
	//
	// }
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// miniapp
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AISearchMessageItem) String() string {
	return dara.Prettify(s)
}

func (s AISearchMessageItem) GoString() string {
	return s.String()
}

func (s *AISearchMessageItem) GetContent() *string {
	return s.Content
}

func (s *AISearchMessageItem) GetParams() interface{} {
	return s.Params
}

func (s *AISearchMessageItem) GetType() *string {
	return s.Type
}

func (s *AISearchMessageItem) SetContent(v string) *AISearchMessageItem {
	s.Content = &v
	return s
}

func (s *AISearchMessageItem) SetParams(v interface{}) *AISearchMessageItem {
	s.Params = v
	return s
}

func (s *AISearchMessageItem) SetType(v string) *AISearchMessageItem {
	s.Type = &v
	return s
}

func (s *AISearchMessageItem) Validate() error {
	return dara.Validate(s)
}
