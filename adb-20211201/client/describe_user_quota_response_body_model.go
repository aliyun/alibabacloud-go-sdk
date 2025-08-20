// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticACU(v string) *DescribeUserQuotaResponseBody
	GetElasticACU() *string
	SetRequestId(v string) *DescribeUserQuotaResponseBody
	GetRequestId() *string
	SetReserverdCompteACU(v string) *DescribeUserQuotaResponseBody
	GetReserverdCompteACU() *string
	SetReserverdStorageACU(v string) *DescribeUserQuotaResponseBody
	GetReserverdStorageACU() *string
	SetResourceGroupCount(v string) *DescribeUserQuotaResponseBody
	GetResourceGroupCount() *string
}

type DescribeUserQuotaResponseBody struct {
	// The available elastic AnalyticDB compute units (ACUs).
	//
	// example:
	//
	// 512ACU
	ElasticACU *string `json:"ElasticACU,omitempty" xml:"ElasticACU,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0322C7FB-4584-5D2A-BF7F-F9036E940C35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The available reserved computing resources.
	//
	// example:
	//
	// 48ACU
	ReserverdCompteACU *string `json:"ReserverdCompteACU,omitempty" xml:"ReserverdCompteACU,omitempty"`
	// The available reserved storage resources.
	//
	// example:
	//
	// 24ACU
	ReserverdStorageACU *string `json:"ReserverdStorageACU,omitempty" xml:"ReserverdStorageACU,omitempty"`
	// The number of available resource groups.
	//
	// example:
	//
	// 10
	ResourceGroupCount *string `json:"ResourceGroupCount,omitempty" xml:"ResourceGroupCount,omitempty"`
}

func (s DescribeUserQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaResponseBody) GetElasticACU() *string {
	return s.ElasticACU
}

func (s *DescribeUserQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserQuotaResponseBody) GetReserverdCompteACU() *string {
	return s.ReserverdCompteACU
}

func (s *DescribeUserQuotaResponseBody) GetReserverdStorageACU() *string {
	return s.ReserverdStorageACU
}

func (s *DescribeUserQuotaResponseBody) GetResourceGroupCount() *string {
	return s.ResourceGroupCount
}

func (s *DescribeUserQuotaResponseBody) SetElasticACU(v string) *DescribeUserQuotaResponseBody {
	s.ElasticACU = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetRequestId(v string) *DescribeUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetReserverdCompteACU(v string) *DescribeUserQuotaResponseBody {
	s.ReserverdCompteACU = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetReserverdStorageACU(v string) *DescribeUserQuotaResponseBody {
	s.ReserverdStorageACU = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetResourceGroupCount(v string) *DescribeUserQuotaResponseBody {
	s.ResourceGroupCount = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
