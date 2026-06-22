// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedInstanceCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceDirectoryAccountId(v string) *DescribeExposedInstanceCriteriaRequest
	GetResourceDirectoryAccountId() *string
	SetValue(v string) *DescribeExposedInstanceCriteriaRequest
	GetValue() *string
}

type DescribeExposedInstanceCriteriaRequest struct {
	// The ID of the member account in the resource directory.
	//
	// >Call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The value of the query condition. Fuzzy match is supported.
	//
	// > This parameter supports queries by asset name, asset ID, public IP address of the asset, private IP address of the asset, exposed component, exposed port, or exposed IP address.
	//
	// example:
	//
	// testInstanceName
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeExposedInstanceCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceCriteriaRequest) GetResourceDirectoryAccountId() *string {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeExposedInstanceCriteriaRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeExposedInstanceCriteriaRequest) SetResourceDirectoryAccountId(v string) *DescribeExposedInstanceCriteriaRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeExposedInstanceCriteriaRequest) SetValue(v string) *DescribeExposedInstanceCriteriaRequest {
	s.Value = &v
	return s
}

func (s *DescribeExposedInstanceCriteriaRequest) Validate() error {
	return dara.Validate(s)
}
