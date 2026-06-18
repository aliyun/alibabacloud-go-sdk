// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateMemoryNodeRequest
	GetContent() *string
}

type UpdateMemoryNodeRequest struct {
	// The memory fragment content.
	//
	// This parameter is required.
	//
	// example:
	//
	// 用户喜欢吃西红柿炒鸡蛋
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s UpdateMemoryNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryNodeRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateMemoryNodeRequest) SetContent(v string) *UpdateMemoryNodeRequest {
	s.Content = &v
	return s
}

func (s *UpdateMemoryNodeRequest) Validate() error {
	return dara.Validate(s)
}
