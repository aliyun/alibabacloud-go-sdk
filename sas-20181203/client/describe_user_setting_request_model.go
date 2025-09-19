// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeUserSettingRequest
	GetSourceIp() *string
}

type DescribeUserSettingRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 58.248.87.10
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeUserSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserSettingRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeUserSettingRequest) SetSourceIp(v string) *DescribeUserSettingRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserSettingRequest) Validate() error {
	return dara.Validate(s)
}
