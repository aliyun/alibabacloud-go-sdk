// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAnalyzerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWith(v string) *DescribeUserAnalyzerRequest
	GetWith() *string
}

type DescribeUserAnalyzerRequest struct {
	// The Associated information,output properties based on hierarchy.
	//
	// 	- **all**: Outputs associated app information
	//
	// example:
	//
	// all
	With *string `json:"with,omitempty" xml:"with,omitempty"`
}

func (s DescribeUserAnalyzerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAnalyzerRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAnalyzerRequest) GetWith() *string {
	return s.With
}

func (s *DescribeUserAnalyzerRequest) SetWith(v string) *DescribeUserAnalyzerRequest {
	s.With = &v
	return s
}

func (s *DescribeUserAnalyzerRequest) Validate() error {
	return dara.Validate(s)
}
