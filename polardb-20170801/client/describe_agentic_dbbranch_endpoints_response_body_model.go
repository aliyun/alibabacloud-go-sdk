// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAgenticDBBranchEndpointsResponseBodyItems) *DescribeAgenticDBBranchEndpointsResponseBody
	GetItems() []*DescribeAgenticDBBranchEndpointsResponseBodyItems
	SetRequestId(v string) *DescribeAgenticDBBranchEndpointsResponseBody
	GetRequestId() *string
}

type DescribeAgenticDBBranchEndpointsResponseBody struct {
	Items []*DescribeAgenticDBBranchEndpointsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// E5F6A7B8-C9D0-1234-EFAB-567890123EFA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAgenticDBBranchEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchEndpointsResponseBody) GetItems() []*DescribeAgenticDBBranchEndpointsResponseBodyItems {
	return s.Items
}

func (s *DescribeAgenticDBBranchEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBBranchEndpointsResponseBody) SetItems(v []*DescribeAgenticDBBranchEndpointsResponseBodyItems) *DescribeAgenticDBBranchEndpointsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponseBody) SetRequestId(v string) *DescribeAgenticDBBranchEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAgenticDBBranchEndpointsResponseBodyItems struct {
	// example:
	//
	// cloud_admin
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// example:
	//
	// 10.0.1.100
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// postgresql://cloud_admin:******@10.0.1.100:5432/neondb
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// neondb
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// example:
	//
	// ep-3m4n5o6p7q8r
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// example:
	//
	// ReadWrite
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 5432
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeAgenticDBBranchEndpointsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchEndpointsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) GetAccount() *string {
	return s.Account
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) GetAddress() *string {
	return s.Address
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) GetDatabase() *string {
	return s.Database
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) GetPassword() *string {
	return s.Password
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) GetPort() *int32 {
	return s.Port
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) SetAccount(v string) *DescribeAgenticDBBranchEndpointsResponseBodyItems {
	s.Account = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) SetAddress(v string) *DescribeAgenticDBBranchEndpointsResponseBodyItems {
	s.Address = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) SetConnectionString(v string) *DescribeAgenticDBBranchEndpointsResponseBodyItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) SetDatabase(v string) *DescribeAgenticDBBranchEndpointsResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) SetEndpointId(v string) *DescribeAgenticDBBranchEndpointsResponseBodyItems {
	s.EndpointId = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) SetEndpointType(v string) *DescribeAgenticDBBranchEndpointsResponseBodyItems {
	s.EndpointType = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) SetPassword(v string) *DescribeAgenticDBBranchEndpointsResponseBodyItems {
	s.Password = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) SetPort(v int32) *DescribeAgenticDBBranchEndpointsResponseBodyItems {
	s.Port = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
