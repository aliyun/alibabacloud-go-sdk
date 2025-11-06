// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppViewStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppViewStatusRequest
	GetAppId() *string
}

type DescribeAppViewStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeAppViewStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppViewStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppViewStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppViewStatusRequest) SetAppId(v string) *DescribeAppViewStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppViewStatusRequest) Validate() error {
	return dara.Validate(s)
}
