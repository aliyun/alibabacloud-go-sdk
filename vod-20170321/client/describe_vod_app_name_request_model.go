// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodAppNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribeVodAppNameRequest
	GetType() *string
}

type DescribeVodAppNameRequest struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVodAppNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAppNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodAppNameRequest) GetType() *string {
	return s.Type
}

func (s *DescribeVodAppNameRequest) SetType(v string) *DescribeVodAppNameRequest {
	s.Type = &v
	return s
}

func (s *DescribeVodAppNameRequest) Validate() error {
	return dara.Validate(s)
}
