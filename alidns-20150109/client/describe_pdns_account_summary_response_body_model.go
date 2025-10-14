// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsAccountSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribePdnsAccountSummaryResponseBodyData) *DescribePdnsAccountSummaryResponseBody
	GetData() *DescribePdnsAccountSummaryResponseBodyData
	SetRequestId(v string) *DescribePdnsAccountSummaryResponseBody
	GetRequestId() *string
}

type DescribePdnsAccountSummaryResponseBody struct {
	Data      *DescribePdnsAccountSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePdnsAccountSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAccountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsAccountSummaryResponseBody) GetData() *DescribePdnsAccountSummaryResponseBodyData {
	return s.Data
}

func (s *DescribePdnsAccountSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsAccountSummaryResponseBody) SetData(v *DescribePdnsAccountSummaryResponseBodyData) *DescribePdnsAccountSummaryResponseBody {
	s.Data = v
	return s
}

func (s *DescribePdnsAccountSummaryResponseBody) SetRequestId(v string) *DescribePdnsAccountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsAccountSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePdnsAccountSummaryResponseBodyData struct {
	DomainCount    *int64 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	HttpCount      *int64 `json:"HttpCount,omitempty" xml:"HttpCount,omitempty"`
	HttpsCount     *int64 `json:"HttpsCount,omitempty" xml:"HttpsCount,omitempty"`
	SubDomainCount *int64 `json:"SubDomainCount,omitempty" xml:"SubDomainCount,omitempty"`
	ThreatCount    *int64 `json:"ThreatCount,omitempty" xml:"ThreatCount,omitempty"`
	TotalCount     *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserId         *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribePdnsAccountSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsAccountSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePdnsAccountSummaryResponseBodyData) GetDomainCount() *int64 {
	return s.DomainCount
}

func (s *DescribePdnsAccountSummaryResponseBodyData) GetHttpCount() *int64 {
	return s.HttpCount
}

func (s *DescribePdnsAccountSummaryResponseBodyData) GetHttpsCount() *int64 {
	return s.HttpsCount
}

func (s *DescribePdnsAccountSummaryResponseBodyData) GetSubDomainCount() *int64 {
	return s.SubDomainCount
}

func (s *DescribePdnsAccountSummaryResponseBodyData) GetThreatCount() *int64 {
	return s.ThreatCount
}

func (s *DescribePdnsAccountSummaryResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePdnsAccountSummaryResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribePdnsAccountSummaryResponseBodyData) SetDomainCount(v int64) *DescribePdnsAccountSummaryResponseBodyData {
	s.DomainCount = &v
	return s
}

func (s *DescribePdnsAccountSummaryResponseBodyData) SetHttpCount(v int64) *DescribePdnsAccountSummaryResponseBodyData {
	s.HttpCount = &v
	return s
}

func (s *DescribePdnsAccountSummaryResponseBodyData) SetHttpsCount(v int64) *DescribePdnsAccountSummaryResponseBodyData {
	s.HttpsCount = &v
	return s
}

func (s *DescribePdnsAccountSummaryResponseBodyData) SetSubDomainCount(v int64) *DescribePdnsAccountSummaryResponseBodyData {
	s.SubDomainCount = &v
	return s
}

func (s *DescribePdnsAccountSummaryResponseBodyData) SetThreatCount(v int64) *DescribePdnsAccountSummaryResponseBodyData {
	s.ThreatCount = &v
	return s
}

func (s *DescribePdnsAccountSummaryResponseBodyData) SetTotalCount(v int64) *DescribePdnsAccountSummaryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribePdnsAccountSummaryResponseBodyData) SetUserId(v int64) *DescribePdnsAccountSummaryResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribePdnsAccountSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
