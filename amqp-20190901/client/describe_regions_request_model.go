// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *DescribeRegionsRequest
	GetConsoleSessionId() *string
}

type DescribeRegionsRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *DescribeRegionsRequest) SetConsoleSessionId(v string) *DescribeRegionsRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
