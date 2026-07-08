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
	// The source of the domain name. Valid values:
	//
	// - **share**: The domain name is added to WAF in CNAME record mode.
	//
	// - **asset**: The domain name is added to WAF as a custom asset.
	//
	// - **hybrid_cloud_cname**: The domain name is added to WAF in hybrid cloud CNAME record mode.
	//
	// - **tgw**: The domain name is added to WAF in cloud native mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// share
	AccessOrigin *string `json:"AccessOrigin,omitempty" xml:"AccessOrigin,omitempty"`
	// The domain name that you want to query for ownership verification content.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
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
