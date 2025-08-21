// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOsType(v string) *DescribeImageInfosRequest
	GetOsType() *string
}

type DescribeImageInfosRequest struct {
	// The operating system (OS). You can specify only one OS in a request. If you do not specify a value for this parameter, images for all supported OSs are queried. Valid values:
	//
	// 	- linux
	//
	// 	- windows
	//
	// example:
	//
	// linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
}

func (s DescribeImageInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageInfosRequest) GetOsType() *string {
	return s.OsType
}

func (s *DescribeImageInfosRequest) SetOsType(v string) *DescribeImageInfosRequest {
	s.OsType = &v
	return s
}

func (s *DescribeImageInfosRequest) Validate() error {
	return dara.Validate(s)
}
