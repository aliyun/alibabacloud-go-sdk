// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationScalingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationScalingRulesRequest
	GetAppId() *string
}

type DescribeApplicationScalingRulesRequest struct {
	// 7171a6ca-d1cd-4928-8642-7d5cfe69\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeApplicationScalingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationScalingRulesRequest) SetAppId(v string) *DescribeApplicationScalingRulesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationScalingRulesRequest) Validate() error {
	return dara.Validate(s)
}
