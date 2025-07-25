// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransferDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainTransfers(v *DescribeTransferDomainsResponseBodyDomainTransfers) *DescribeTransferDomainsResponseBody
	GetDomainTransfers() *DescribeTransferDomainsResponseBodyDomainTransfers
	SetPageNumber(v int64) *DescribeTransferDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeTransferDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeTransferDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeTransferDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeTransferDomainsResponseBody struct {
	// The domain names that were transferred between accounts.
	DomainTransfers *DescribeTransferDomainsResponseBodyDomainTransfers `json:"DomainTransfers,omitempty" xml:"DomainTransfers,omitempty" type:"Struct"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTransferDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransferDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransferDomainsResponseBody) GetDomainTransfers() *DescribeTransferDomainsResponseBodyDomainTransfers {
	return s.DomainTransfers
}

func (s *DescribeTransferDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeTransferDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeTransferDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTransferDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeTransferDomainsResponseBody) SetDomainTransfers(v *DescribeTransferDomainsResponseBodyDomainTransfers) *DescribeTransferDomainsResponseBody {
	s.DomainTransfers = v
	return s
}

func (s *DescribeTransferDomainsResponseBody) SetPageNumber(v int64) *DescribeTransferDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTransferDomainsResponseBody) SetPageSize(v int64) *DescribeTransferDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTransferDomainsResponseBody) SetRequestId(v string) *DescribeTransferDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransferDomainsResponseBody) SetTotalCount(v int64) *DescribeTransferDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTransferDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTransferDomainsResponseBodyDomainTransfers struct {
	DomainTransfer []*DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer `json:"DomainTransfer,omitempty" xml:"DomainTransfer,omitempty" type:"Repeated"`
}

func (s DescribeTransferDomainsResponseBodyDomainTransfers) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransferDomainsResponseBodyDomainTransfers) GoString() string {
	return s.String()
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfers) GetDomainTransfer() []*DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer {
	return s.DomainTransfer
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfers) SetDomainTransfer(v []*DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) *DescribeTransferDomainsResponseBodyDomainTransfers {
	s.DomainTransfer = v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfers) Validate() error {
	return dara.Validate(s)
}

type DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer struct {
	// The time when the domain name was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-10-30T07:16Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the domain name was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1572419764000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The user ID from which the domain name was transferred.
	//
	// example:
	//
	// 2222
	FromUserId *int64 `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
	// The ID of the domain name that was transferred.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The user ID to which the domain name was transferred.
	//
	// example:
	//
	// 111111
	TargetUserId *int64 `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) GoString() string {
	return s.String()
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) GetFromUserId() *int64 {
	return s.FromUserId
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) GetId() *int64 {
	return s.Id
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) SetCreateTime(v string) *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer {
	s.CreateTime = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) SetCreateTimestamp(v int64) *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) SetDomainName(v string) *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer {
	s.DomainName = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) SetFromUserId(v int64) *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer {
	s.FromUserId = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) SetId(v int64) *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer {
	s.Id = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) SetTargetUserId(v int64) *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer {
	s.TargetUserId = &v
	return s
}

func (s *DescribeTransferDomainsResponseBodyDomainTransfersDomainTransfer) Validate() error {
	return dara.Validate(s)
}
