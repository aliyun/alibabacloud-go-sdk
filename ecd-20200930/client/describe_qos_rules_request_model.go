// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPackageId(v string) *DescribeQosRulesRequest
	GetNetworkPackageId() *string
	SetQosRuleName(v string) *DescribeQosRulesRequest
	GetQosRuleName() *string
}

type DescribeQosRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// np-cxj99qb8d34vo****
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// example:
	//
	// test
	QosRuleName *string `json:"QosRuleName,omitempty" xml:"QosRuleName,omitempty"`
}

func (s DescribeQosRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeQosRulesRequest) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *DescribeQosRulesRequest) GetQosRuleName() *string {
	return s.QosRuleName
}

func (s *DescribeQosRulesRequest) SetNetworkPackageId(v string) *DescribeQosRulesRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *DescribeQosRulesRequest) SetQosRuleName(v string) *DescribeQosRulesRequest {
	s.QosRuleName = &v
	return s
}

func (s *DescribeQosRulesRequest) Validate() error {
	return dara.Validate(s)
}
