// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIMRobotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int64) *DescribeIMRobotsRequest
	GetPage() *int64
	SetRobotIds(v string) *DescribeIMRobotsRequest
	GetRobotIds() *string
	SetRobotName(v string) *DescribeIMRobotsRequest
	GetRobotName() *string
	SetSize(v int64) *DescribeIMRobotsRequest
	GetSize() *int64
}

type DescribeIMRobotsRequest struct {
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The chatbot IDs.
	//
	// example:
	//
	// 123
	RobotIds *string `json:"RobotIds,omitempty" xml:"RobotIds,omitempty"`
	// The name of the IM chatbot.
	//
	// example:
	//
	// Chatbot name
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// The number of IM chatbots to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeIMRobotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIMRobotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsRequest) GetPage() *int64 {
	return s.Page
}

func (s *DescribeIMRobotsRequest) GetRobotIds() *string {
	return s.RobotIds
}

func (s *DescribeIMRobotsRequest) GetRobotName() *string {
	return s.RobotName
}

func (s *DescribeIMRobotsRequest) GetSize() *int64 {
	return s.Size
}

func (s *DescribeIMRobotsRequest) SetPage(v int64) *DescribeIMRobotsRequest {
	s.Page = &v
	return s
}

func (s *DescribeIMRobotsRequest) SetRobotIds(v string) *DescribeIMRobotsRequest {
	s.RobotIds = &v
	return s
}

func (s *DescribeIMRobotsRequest) SetRobotName(v string) *DescribeIMRobotsRequest {
	s.RobotName = &v
	return s
}

func (s *DescribeIMRobotsRequest) SetSize(v int64) *DescribeIMRobotsRequest {
	s.Size = &v
	return s
}

func (s *DescribeIMRobotsRequest) Validate() error {
	return dara.Validate(s)
}
