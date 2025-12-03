// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNoticeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribeNoticeConfigRequest
	GetBizType() *string
	SetSourceIp(v string) *DescribeNoticeConfigRequest
	GetSourceIp() *string
}

type DescribeNoticeConfigRequest struct {
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 60.191.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeNoticeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNoticeConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeNoticeConfigRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeNoticeConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeNoticeConfigRequest) SetBizType(v string) *DescribeNoticeConfigRequest {
	s.BizType = &v
	return s
}

func (s *DescribeNoticeConfigRequest) SetSourceIp(v string) *DescribeNoticeConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeNoticeConfigRequest) Validate() error {
	return dara.Validate(s)
}
