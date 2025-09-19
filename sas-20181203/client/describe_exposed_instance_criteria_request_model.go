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
	// The ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the account ID.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The value of the search condition. Fuzzy match is supported.
	//
	// >  You can specify the name, ID, public IP address, private IP address, component, port, or IP address of an exposed asset.
	//
	// example:
	//
	// id
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
