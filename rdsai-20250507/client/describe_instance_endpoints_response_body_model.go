// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *DescribeInstanceEndpointsResponseBody
	GetBranchName() *string
	SetDBInstanceEndpoints(v []*DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) *DescribeInstanceEndpointsResponseBody
	GetDBInstanceEndpoints() []*DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints
	SetInstanceEndpoints(v []*DescribeInstanceEndpointsResponseBodyInstanceEndpoints) *DescribeInstanceEndpointsResponseBody
	GetInstanceEndpoints() []*DescribeInstanceEndpointsResponseBodyInstanceEndpoints
	SetInstanceName(v string) *DescribeInstanceEndpointsResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *DescribeInstanceEndpointsResponseBody
	GetRequestId() *string
}

type DescribeInstanceEndpointsResponseBody struct {
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The list of endpoint information of the database instance.
	DBInstanceEndpoints []*DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints `json:"DBInstanceEndpoints,omitempty" xml:"DBInstanceEndpoints,omitempty" type:"Repeated"`
	// The list of endpoint information of the AI application instance.
	InstanceEndpoints []*DescribeInstanceEndpointsResponseBodyInstanceEndpoints `json:"InstanceEndpoints,omitempty" xml:"InstanceEndpoints,omitempty" type:"Repeated"`
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32DEFB4A-861F-5D1D-ADD5-918E4FD7AB8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceEndpointsResponseBody) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeInstanceEndpointsResponseBody) GetDBInstanceEndpoints() []*DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints {
	return s.DBInstanceEndpoints
}

func (s *DescribeInstanceEndpointsResponseBody) GetInstanceEndpoints() []*DescribeInstanceEndpointsResponseBodyInstanceEndpoints {
	return s.InstanceEndpoints
}

func (s *DescribeInstanceEndpointsResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceEndpointsResponseBody) SetBranchName(v string) *DescribeInstanceEndpointsResponseBody {
	s.BranchName = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBody) SetDBInstanceEndpoints(v []*DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) *DescribeInstanceEndpointsResponseBody {
	s.DBInstanceEndpoints = v
	return s
}

func (s *DescribeInstanceEndpointsResponseBody) SetInstanceEndpoints(v []*DescribeInstanceEndpointsResponseBodyInstanceEndpoints) *DescribeInstanceEndpointsResponseBody {
	s.InstanceEndpoints = v
	return s
}

func (s *DescribeInstanceEndpointsResponseBody) SetInstanceName(v string) *DescribeInstanceEndpointsResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBody) SetRequestId(v string) *DescribeInstanceEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBody) Validate() error {
	if s.DBInstanceEndpoints != nil {
		for _, item := range s.DBInstanceEndpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InstanceEndpoints != nil {
		for _, item := range s.InstanceEndpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints struct {
	// The endpoint.
	//
	// example:
	//
	// pgm-xxxx.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The network type. Valid values:
	//
	// - **public**: Internet.
	//
	// - **vpc**: private network.
	//
	// example:
	//
	// vpc
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The connection port.
	//
	// example:
	//
	// 5432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) GetIpType() *string {
	return s.IpType
}

func (s *DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) GetPort() *string {
	return s.Port
}

func (s *DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) SetConnectionString(v string) *DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints {
	s.ConnectionString = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) SetIpType(v string) *DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints {
	s.IpType = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) SetPort(v string) *DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints {
	s.Port = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBodyDBInstanceEndpoints) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceEndpointsResponseBodyInstanceEndpoints struct {
	// The endpoint.
	//
	// example:
	//
	// 8.152.XXX.XXX:8000
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The domain name.
	//
	// example:
	//
	// xxx.apsaradb.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 8.152.XXX.XXX
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The network type. Valid values:
	//
	// - **public**: Internet.
	//
	// - **vpc**: private network.
	//
	// example:
	//
	// public
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The connection port.
	//
	// example:
	//
	// 8000
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeInstanceEndpointsResponseBodyInstanceEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceEndpointsResponseBodyInstanceEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) GetDomain() *string {
	return s.Domain
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) GetIP() *string {
	return s.IP
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) GetIpType() *string {
	return s.IpType
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) GetPort() *string {
	return s.Port
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) SetConnectionString(v string) *DescribeInstanceEndpointsResponseBodyInstanceEndpoints {
	s.ConnectionString = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) SetDomain(v string) *DescribeInstanceEndpointsResponseBodyInstanceEndpoints {
	s.Domain = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) SetIP(v string) *DescribeInstanceEndpointsResponseBodyInstanceEndpoints {
	s.IP = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) SetIpType(v string) *DescribeInstanceEndpointsResponseBodyInstanceEndpoints {
	s.IpType = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) SetPort(v string) *DescribeInstanceEndpointsResponseBodyInstanceEndpoints {
	s.Port = &v
	return s
}

func (s *DescribeInstanceEndpointsResponseBodyInstanceEndpoints) Validate() error {
	return dara.Validate(s)
}
