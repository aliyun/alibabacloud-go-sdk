// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsedServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUsedServiceResponseBody
	GetRequestId() *string
	SetRecords(v []*DescribeUsedServiceResponseBodyRecords) *DescribeUsedServiceResponseBody
	GetRecords() []*DescribeUsedServiceResponseBodyRecords
}

type DescribeUsedServiceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Record details
	Records []*DescribeUsedServiceResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s DescribeUsedServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsedServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsedServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUsedServiceResponseBody) GetRecords() []*DescribeUsedServiceResponseBodyRecords {
	return s.Records
}

func (s *DescribeUsedServiceResponseBody) SetRequestId(v string) *DescribeUsedServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsedServiceResponseBody) SetRecords(v []*DescribeUsedServiceResponseBodyRecords) *DescribeUsedServiceResponseBody {
	s.Records = v
	return s
}

func (s *DescribeUsedServiceResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUsedServiceResponseBodyRecords struct {
	// English name
	//
	// example:
	//
	// account_abuse
	EnName *string `json:"enName,omitempty" xml:"enName,omitempty"`
	// Service name
	//
	// example:
	//
	// 注册风险识别服务
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Service code
	//
	// example:
	//
	// account_abuse
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
}

func (s DescribeUsedServiceResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsedServiceResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeUsedServiceResponseBodyRecords) GetEnName() *string {
	return s.EnName
}

func (s *DescribeUsedServiceResponseBodyRecords) GetName() *string {
	return s.Name
}

func (s *DescribeUsedServiceResponseBodyRecords) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeUsedServiceResponseBodyRecords) SetEnName(v string) *DescribeUsedServiceResponseBodyRecords {
	s.EnName = &v
	return s
}

func (s *DescribeUsedServiceResponseBodyRecords) SetName(v string) *DescribeUsedServiceResponseBodyRecords {
	s.Name = &v
	return s
}

func (s *DescribeUsedServiceResponseBodyRecords) SetServiceCode(v string) *DescribeUsedServiceResponseBodyRecords {
	s.ServiceCode = &v
	return s
}

func (s *DescribeUsedServiceResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
