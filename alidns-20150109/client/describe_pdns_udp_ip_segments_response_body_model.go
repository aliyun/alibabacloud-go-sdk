// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsUdpIpSegmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpSegments(v []*DescribePdnsUdpIpSegmentsResponseBodyIpSegments) *DescribePdnsUdpIpSegmentsResponseBody
	GetIpSegments() []*DescribePdnsUdpIpSegmentsResponseBodyIpSegments
	SetPageNumber(v int64) *DescribePdnsUdpIpSegmentsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePdnsUdpIpSegmentsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribePdnsUdpIpSegmentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePdnsUdpIpSegmentsResponseBody
	GetTotalCount() *int64
	SetTotalPages(v string) *DescribePdnsUdpIpSegmentsResponseBody
	GetTotalPages() *string
}

type DescribePdnsUdpIpSegmentsResponseBody struct {
	IpSegments []*DescribePdnsUdpIpSegmentsResponseBodyIpSegments `json:"IpSegments,omitempty" xml:"IpSegments,omitempty" type:"Repeated"`
	PageNumber *int64                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPages *string                                            `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribePdnsUdpIpSegmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsUdpIpSegmentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) GetIpSegments() []*DescribePdnsUdpIpSegmentsResponseBodyIpSegments {
	return s.IpSegments
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) GetTotalPages() *string {
	return s.TotalPages
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) SetIpSegments(v []*DescribePdnsUdpIpSegmentsResponseBodyIpSegments) *DescribePdnsUdpIpSegmentsResponseBody {
	s.IpSegments = v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) SetPageNumber(v int64) *DescribePdnsUdpIpSegmentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) SetPageSize(v int64) *DescribePdnsUdpIpSegmentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) SetRequestId(v string) *DescribePdnsUdpIpSegmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) SetTotalCount(v int64) *DescribePdnsUdpIpSegmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) SetTotalPages(v string) *DescribePdnsUdpIpSegmentsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBody) Validate() error {
	if s.IpSegments != nil {
		for _, item := range s.IpSegments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePdnsUdpIpSegmentsResponseBodyIpSegments struct {
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Mask            *int64  `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SecretKey       *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s DescribePdnsUdpIpSegmentsResponseBodyIpSegments) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsUdpIpSegmentsResponseBodyIpSegments) GoString() string {
	return s.String()
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) GetCreateDate() *string {
	return s.CreateDate
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) GetId() *string {
	return s.Id
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) GetIp() *string {
	return s.Ip
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) GetMask() *int64 {
	return s.Mask
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) GetName() *string {
	return s.Name
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) GetSecretKey() *string {
	return s.SecretKey
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) GetState() *string {
	return s.State
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) SetCreateDate(v string) *DescribePdnsUdpIpSegmentsResponseBodyIpSegments {
	s.CreateDate = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) SetCreateTimestamp(v int64) *DescribePdnsUdpIpSegmentsResponseBodyIpSegments {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) SetId(v string) *DescribePdnsUdpIpSegmentsResponseBodyIpSegments {
	s.Id = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) SetIp(v string) *DescribePdnsUdpIpSegmentsResponseBodyIpSegments {
	s.Ip = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) SetMask(v int64) *DescribePdnsUdpIpSegmentsResponseBodyIpSegments {
	s.Mask = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) SetName(v string) *DescribePdnsUdpIpSegmentsResponseBodyIpSegments {
	s.Name = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) SetSecretKey(v string) *DescribePdnsUdpIpSegmentsResponseBodyIpSegments {
	s.SecretKey = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) SetState(v string) *DescribePdnsUdpIpSegmentsResponseBodyIpSegments {
	s.State = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) SetUpdateDate(v string) *DescribePdnsUdpIpSegmentsResponseBodyIpSegments {
	s.UpdateDate = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsResponseBodyIpSegments) Validate() error {
	return dara.Validate(s)
}
