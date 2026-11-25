// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessOrigin(v string) *DescribeVerifyContentRequest
	GetAccessOrigin() *string
	SetDomainName(v string) *DescribeVerifyContentRequest
	GetDomainName() *string
	SetInstanceId(v string) *DescribeVerifyContentRequest
	GetInstanceId() *string
}

type DescribeVerifyContentRequest struct {
	// The access source. Valid values:
	//
	// - **share**: public cloud CNAME access
	//
	// - **asset**: custom asset access
	//
	// - **hybrid_cloud_cname**: hybrid cloud CNAME access
	//
	// - **tgw**: cloud native mode
	//
	// This parameter is required.
	//
	// example:
	//
	// share
	AccessOrigin *string `json:"AccessOrigin,omitempty" xml:"AccessOrigin,omitempty"`
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the WAF instance. You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeVerifyContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyContentRequest) GetAccessOrigin() *string {
	return s.AccessOrigin
}

func (s *DescribeVerifyContentRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVerifyContentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVerifyContentRequest) SetAccessOrigin(v string) *DescribeVerifyContentRequest {
	s.AccessOrigin = &v
	return s
}

func (s *DescribeVerifyContentRequest) SetDomainName(v string) *DescribeVerifyContentRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVerifyContentRequest) SetInstanceId(v string) *DescribeVerifyContentRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVerifyContentRequest) Validate() error {
	return dara.Validate(s)
}
