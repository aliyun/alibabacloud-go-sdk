// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsPullStreamInfoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsPullStreamInfoConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVsPullStreamInfoConfigRequest
	GetOwnerId() *int64
}

type DescribeVsPullStreamInfoConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsPullStreamInfoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamInfoConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsPullStreamInfoConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsPullStreamInfoConfigRequest) SetDomainName(v string) *DescribeVsPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigRequest) SetOwnerId(v int64) *DescribeVsPullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigRequest) Validate() error {
	return dara.Validate(s)
}
