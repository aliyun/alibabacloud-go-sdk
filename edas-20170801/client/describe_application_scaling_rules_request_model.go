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
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// example:
	//
	// 78194c76-3dca-418e-a263-cccd1ab4****
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
