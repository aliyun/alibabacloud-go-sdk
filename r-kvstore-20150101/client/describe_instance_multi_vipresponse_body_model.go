// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMultiVIPResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMasterDns(v string) *DescribeInstanceMultiVIPResponseBody
	GetMasterDns() *string
	SetMasterDnsRecord(v []*string) *DescribeInstanceMultiVIPResponseBody
	GetMasterDnsRecord() []*string
	SetMaxQuota(v int64) *DescribeInstanceMultiVIPResponseBody
	GetMaxQuota() *int64
	SetMultiVIPList(v []*DescribeInstanceMultiVIPResponseBodyMultiVIPList) *DescribeInstanceMultiVIPResponseBody
	GetMultiVIPList() []*DescribeInstanceMultiVIPResponseBodyMultiVIPList
	SetRequestId(v string) *DescribeInstanceMultiVIPResponseBody
	GetRequestId() *string
}

type DescribeInstanceMultiVIPResponseBody struct {
	// The IP address of the primary DNS server.
	//
	// example:
	//
	// r-8vb30e8n0m4nvu7tff.redis.zhangbei.rds.aliyuncs.com
	MasterDns *string `json:"MasterDns,omitempty" xml:"MasterDns,omitempty"`
	// The list of IP addresses corresponding to all LBs of the current instance.
	MasterDnsRecord []*string `json:"MasterDnsRecord,omitempty" xml:"MasterDnsRecord,omitempty" type:"Repeated"`
	// The maximum number of VIPs that can be created.
	//
	// example:
	//
	// 2
	MaxQuota *int64 `json:"MaxQuota,omitempty" xml:"MaxQuota,omitempty"`
	// The list of multiple LB VIPs.
	MultiVIPList []*DescribeInstanceMultiVIPResponseBodyMultiVIPList `json:"MultiVIPList,omitempty" xml:"MultiVIPList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ABAF95F6-35C1-4177-AF3A-70969EBD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceMultiVIPResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMultiVIPResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMultiVIPResponseBody) GetMasterDns() *string {
	return s.MasterDns
}

func (s *DescribeInstanceMultiVIPResponseBody) GetMasterDnsRecord() []*string {
	return s.MasterDnsRecord
}

func (s *DescribeInstanceMultiVIPResponseBody) GetMaxQuota() *int64 {
	return s.MaxQuota
}

func (s *DescribeInstanceMultiVIPResponseBody) GetMultiVIPList() []*DescribeInstanceMultiVIPResponseBodyMultiVIPList {
	return s.MultiVIPList
}

func (s *DescribeInstanceMultiVIPResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceMultiVIPResponseBody) SetMasterDns(v string) *DescribeInstanceMultiVIPResponseBody {
	s.MasterDns = &v
	return s
}

func (s *DescribeInstanceMultiVIPResponseBody) SetMasterDnsRecord(v []*string) *DescribeInstanceMultiVIPResponseBody {
	s.MasterDnsRecord = v
	return s
}

func (s *DescribeInstanceMultiVIPResponseBody) SetMaxQuota(v int64) *DescribeInstanceMultiVIPResponseBody {
	s.MaxQuota = &v
	return s
}

func (s *DescribeInstanceMultiVIPResponseBody) SetMultiVIPList(v []*DescribeInstanceMultiVIPResponseBodyMultiVIPList) *DescribeInstanceMultiVIPResponseBody {
	s.MultiVIPList = v
	return s
}

func (s *DescribeInstanceMultiVIPResponseBody) SetRequestId(v string) *DescribeInstanceMultiVIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceMultiVIPResponseBody) Validate() error {
	if s.MultiVIPList != nil {
		for _, item := range s.MultiVIPList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceMultiVIPResponseBodyMultiVIPList struct {
	// The endpoint of the instance.
	//
	// example:
	//
	// r-bp1p4pzsr2rtubcvns-conn1.redis.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
}

func (s DescribeInstanceMultiVIPResponseBodyMultiVIPList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMultiVIPResponseBodyMultiVIPList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMultiVIPResponseBodyMultiVIPList) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeInstanceMultiVIPResponseBodyMultiVIPList) SetConnectionString(v string) *DescribeInstanceMultiVIPResponseBodyMultiVIPList {
	s.ConnectionString = &v
	return s
}

func (s *DescribeInstanceMultiVIPResponseBodyMultiVIPList) Validate() error {
	return dara.Validate(s)
}
