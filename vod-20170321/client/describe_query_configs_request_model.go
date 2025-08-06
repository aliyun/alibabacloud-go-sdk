// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribeQueryConfigsRequest
	GetType() *string
}

type DescribeQueryConfigsRequest struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeQueryConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeQueryConfigsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeQueryConfigsRequest) SetType(v string) *DescribeQueryConfigsRequest {
	s.Type = &v
	return s
}

func (s *DescribeQueryConfigsRequest) Validate() error {
	return dara.Validate(s)
}
