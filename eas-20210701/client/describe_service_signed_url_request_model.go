// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceSignedUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpire(v int64) *DescribeServiceSignedUrlRequest
	GetExpire() *int64
	SetInternal(v bool) *DescribeServiceSignedUrlRequest
	GetInternal() *bool
	SetType(v string) *DescribeServiceSignedUrlRequest
	GetType() *string
}

type DescribeServiceSignedUrlRequest struct {
	// The period of time for which the URL expires.
	//
	// example:
	//
	// 43200
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// Specifies whether to use the VPC connection.
	//
	// example:
	//
	// false
	Internal *bool `json:"Internal,omitempty" xml:"Internal,omitempty"`
	// The page type.
	//
	// Valid values:
	//
	// 	- webview
	//
	// 	- monitor
	//
	// example:
	//
	// webview
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeServiceSignedUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceSignedUrlRequest) GetExpire() *int64 {
	return s.Expire
}

func (s *DescribeServiceSignedUrlRequest) GetInternal() *bool {
	return s.Internal
}

func (s *DescribeServiceSignedUrlRequest) GetType() *string {
	return s.Type
}

func (s *DescribeServiceSignedUrlRequest) SetExpire(v int64) *DescribeServiceSignedUrlRequest {
	s.Expire = &v
	return s
}

func (s *DescribeServiceSignedUrlRequest) SetInternal(v bool) *DescribeServiceSignedUrlRequest {
	s.Internal = &v
	return s
}

func (s *DescribeServiceSignedUrlRequest) SetType(v string) *DescribeServiceSignedUrlRequest {
	s.Type = &v
	return s
}

func (s *DescribeServiceSignedUrlRequest) Validate() error {
	return dara.Validate(s)
}
