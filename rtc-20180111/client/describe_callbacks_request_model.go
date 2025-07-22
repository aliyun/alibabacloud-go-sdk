// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCallbacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeCallbacksRequest
	GetAppId() *string
}

type DescribeCallbacksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 223***JQb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeCallbacksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallbacksRequest) GoString() string {
	return s.String()
}

func (s *DescribeCallbacksRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCallbacksRequest) SetAppId(v string) *DescribeCallbacksRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCallbacksRequest) Validate() error {
	return dara.Validate(s)
}
