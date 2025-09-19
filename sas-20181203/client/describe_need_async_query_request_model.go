// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNeedAsyncQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribeNeedAsyncQueryRequest
	GetType() *string
}

type DescribeNeedAsyncQueryRequest struct {
	// The type of the query. Valid values:
	//
	// 	- **suspicious**: alerts
	//
	// This parameter is required.
	//
	// example:
	//
	// suspicious
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNeedAsyncQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNeedAsyncQueryRequest) GoString() string {
	return s.String()
}

func (s *DescribeNeedAsyncQueryRequest) GetType() *string {
	return s.Type
}

func (s *DescribeNeedAsyncQueryRequest) SetType(v string) *DescribeNeedAsyncQueryRequest {
	s.Type = &v
	return s
}

func (s *DescribeNeedAsyncQueryRequest) Validate() error {
	return dara.Validate(s)
}
