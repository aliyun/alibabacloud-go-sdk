// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayMetricAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribePlayMetricAuthRequest
	GetType() *string
}

type DescribePlayMetricAuthRequest struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePlayMetricAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayMetricAuthRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayMetricAuthRequest) GetType() *string {
	return s.Type
}

func (s *DescribePlayMetricAuthRequest) SetType(v string) *DescribePlayMetricAuthRequest {
	s.Type = &v
	return s
}

func (s *DescribePlayMetricAuthRequest) Validate() error {
	return dara.Validate(s)
}
