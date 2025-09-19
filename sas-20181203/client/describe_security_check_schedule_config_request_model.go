// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityCheckScheduleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSecurityCheckScheduleConfigRequest
	GetLang() *string
	SetResourceOwnerId(v int64) *DescribeSecurityCheckScheduleConfigRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeSecurityCheckScheduleConfigRequest
	GetSourceIp() *string
}

type DescribeSecurityCheckScheduleConfigRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSecurityCheckScheduleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityCheckScheduleConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityCheckScheduleConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSecurityCheckScheduleConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecurityCheckScheduleConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSecurityCheckScheduleConfigRequest) SetLang(v string) *DescribeSecurityCheckScheduleConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigRequest) SetResourceOwnerId(v int64) *DescribeSecurityCheckScheduleConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigRequest) SetSourceIp(v string) *DescribeSecurityCheckScheduleConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigRequest) Validate() error {
	return dara.Validate(s)
}
