// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSecurityGroupRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDBInstanceSecurityGroupRuleResponseBody
	GetCode() *string
	SetData(v string) *DescribeDBInstanceSecurityGroupRuleResponseBody
	GetData() *string
	SetMessage(v string) *DescribeDBInstanceSecurityGroupRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDBInstanceSecurityGroupRuleResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceSecurityGroupRuleResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the security group rule.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The information about the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6BFF4136-10B9-130B-BF8E-D3504BCE4F43
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceSecurityGroupRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSecurityGroupRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDBInstanceSecurityGroupRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeDBInstanceSecurityGroupRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDBInstanceSecurityGroupRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceSecurityGroupRuleResponseBody) SetCode(v string) *DescribeDBInstanceSecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleResponseBody) SetData(v string) *DescribeDBInstanceSecurityGroupRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleResponseBody) SetMessage(v string) *DescribeDBInstanceSecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleResponseBody) SetRequestId(v string) *DescribeDBInstanceSecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSecurityGroupRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
