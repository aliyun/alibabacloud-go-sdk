// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentThemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetThemeId(v string) *DescribeDataAgentThemeRequest
	GetThemeId() *string
}

type DescribeDataAgentThemeRequest struct {
	// The business ID of the theme.
	//
	// example:
	//
	// 0f8b2c1d************9a3e5f7b1c2d
	ThemeId *string `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
}

func (s DescribeDataAgentThemeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentThemeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentThemeRequest) GetThemeId() *string {
	return s.ThemeId
}

func (s *DescribeDataAgentThemeRequest) SetThemeId(v string) *DescribeDataAgentThemeRequest {
	s.ThemeId = &v
	return s
}

func (s *DescribeDataAgentThemeRequest) Validate() error {
	return dara.Validate(s)
}
