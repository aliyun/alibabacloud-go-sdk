// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilterConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribeFilterConfigsRequest
	GetType() *string
}

type DescribeFilterConfigsRequest struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFilterConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilterConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFilterConfigsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeFilterConfigsRequest) SetType(v string) *DescribeFilterConfigsRequest {
	s.Type = &v
	return s
}

func (s *DescribeFilterConfigsRequest) Validate() error {
	return dara.Validate(s)
}
