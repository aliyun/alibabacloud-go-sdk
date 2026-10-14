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
	// The operating system type. Only a single type can be specified. If this parameter is not specified, all types are queried.
	//
	// Valid values:
	//
	// - Linux
	//
	// - windows
	//
	// In practice, passing linux or Linux may return an empty list, whereas omitting this parameter can return images with an OSType of linux. Verify the actual results before relying on this filter.
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
