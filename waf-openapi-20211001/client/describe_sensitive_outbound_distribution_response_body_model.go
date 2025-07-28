// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveOutboundDistributionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSensitiveOutboundDistributionResponseBodyData) *DescribeSensitiveOutboundDistributionResponseBody
	GetData() []*DescribeSensitiveOutboundDistributionResponseBodyData
	SetRequestId(v string) *DescribeSensitiveOutboundDistributionResponseBody
	GetRequestId() *string
}

type DescribeSensitiveOutboundDistributionResponseBody struct {
	// The traffic distribution of personal information records involved in cross-border data transfer.
	Data []*DescribeSensitiveOutboundDistributionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSensitiveOutboundDistributionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundDistributionResponseBody) GetData() []*DescribeSensitiveOutboundDistributionResponseBodyData {
	return s.Data
}

func (s *DescribeSensitiveOutboundDistributionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSensitiveOutboundDistributionResponseBody) SetData(v []*DescribeSensitiveOutboundDistributionResponseBodyData) *DescribeSensitiveOutboundDistributionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSensitiveOutboundDistributionResponseBody) SetRequestId(v string) *DescribeSensitiveOutboundDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSensitiveOutboundDistributionResponseBodyData struct {
	// The country to which the data is transferred.
	//
	// example:
	//
	// US
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The number of personal information records involved in cross-border data transfer.
	//
	// example:
	//
	// 213
	InfoOutboundCount *int64 `json:"InfoOutboundCount,omitempty" xml:"InfoOutboundCount,omitempty"`
	// The number of sensitive information records involved in cross-border data transfer.
	//
	// example:
	//
	// 144
	SensitiveOutboundCount *int64 `json:"SensitiveOutboundCount,omitempty" xml:"SensitiveOutboundCount,omitempty"`
}

func (s DescribeSensitiveOutboundDistributionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundDistributionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundDistributionResponseBodyData) GetCountry() *string {
	return s.Country
}

func (s *DescribeSensitiveOutboundDistributionResponseBodyData) GetInfoOutboundCount() *int64 {
	return s.InfoOutboundCount
}

func (s *DescribeSensitiveOutboundDistributionResponseBodyData) GetSensitiveOutboundCount() *int64 {
	return s.SensitiveOutboundCount
}

func (s *DescribeSensitiveOutboundDistributionResponseBodyData) SetCountry(v string) *DescribeSensitiveOutboundDistributionResponseBodyData {
	s.Country = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionResponseBodyData) SetInfoOutboundCount(v int64) *DescribeSensitiveOutboundDistributionResponseBodyData {
	s.InfoOutboundCount = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionResponseBodyData) SetSensitiveOutboundCount(v int64) *DescribeSensitiveOutboundDistributionResponseBodyData {
	s.SensitiveOutboundCount = &v
	return s
}

func (s *DescribeSensitiveOutboundDistributionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
