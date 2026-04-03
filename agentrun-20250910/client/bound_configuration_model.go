// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBoundConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetBoundTools(v []*BoundTool) *BoundConfiguration
	GetBoundTools() []*BoundTool
}

type BoundConfiguration struct {
	// 绑定的工具配置列表，每个配置项定义一个工具与特定 HTTP 路径和方法的绑定关系
	BoundTools []*BoundTool `json:"boundTools" xml:"boundTools" type:"Repeated"`
}

func (s BoundConfiguration) String() string {
	return dara.Prettify(s)
}

func (s BoundConfiguration) GoString() string {
	return s.String()
}

func (s *BoundConfiguration) GetBoundTools() []*BoundTool {
	return s.BoundTools
}

func (s *BoundConfiguration) SetBoundTools(v []*BoundTool) *BoundConfiguration {
	s.BoundTools = v
	return s
}

func (s *BoundConfiguration) Validate() error {
	if s.BoundTools != nil {
		for _, item := range s.BoundTools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
