// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsStreamsNotifyUrlConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsStreamsNotifyUrlConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVsStreamsNotifyUrlConfigRequest
	GetOwnerId() *int64
}

type DescribeVsStreamsNotifyUrlConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsStreamsNotifyUrlConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsNotifyUrlConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsStreamsNotifyUrlConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsStreamsNotifyUrlConfigRequest) SetDomainName(v string) *DescribeVsStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigRequest) SetOwnerId(v int64) *DescribeVsStreamsNotifyUrlConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigRequest) Validate() error {
	return dara.Validate(s)
}
