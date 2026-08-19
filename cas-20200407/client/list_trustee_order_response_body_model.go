// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrusteeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTrusteeOrderResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTrusteeOrderResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTrusteeOrderResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTrusteeOrderResponseBody
	GetTotalCount() *int32
	SetTrusteeOrderList(v string) *ListTrusteeOrderResponseBody
	GetTrusteeOrderList() *string
}

type ListTrusteeOrderResponseBody struct {
	// The maximum number of records to return in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. If NextToken is empty, no more results are available.
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0068247C-A454-5FC9-93BF-C41CBB5CD19E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of orders.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The order list.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "VerifyStatus": 0,
	//
	//     "ShowRefund": false,
	//
	//     "CertificateId": 0,
	//
	//     "SourceType": "buy",
	//
	//     "IsMix": false,
	//
	//     "CertType": "DV",
	//
	//     "PartnerOrderId": null,
	//
	//     "ProductId": 53,
	//
	//     "StatusCode": "closed",
	//
	//     "KeyProtection": "UNKNOWN",
	//
	//     "BrandName": "Rapid",
	//
	//     "JobStatus": "editing",
	//
	//     "Month": 6,
	//
	//     "IsFree": false,
	//
	//     "DomainType": "ONE",
	//
	//     "IsRefunding": false,
	//
	//     "RevokeReturnCount": false,
	//
	//     "JobId": 440231,
	//
	//     "DomainCount": 1,
	//
	//     "InstanceId": "cas-ivauto-fe7kv4-15650439-renew",
	//
	//     "ProductCode": "geotrust-dv-1-starter",
	//
	//     "WildDomainCount": 0,
	//
	//     "OrderId": 15652305,
	//
	//     "Algorithm": "RSA",
	//
	//     "Year": 1,
	//
	//     "IsRenew": false,
	//
	//     "Domain": "tw.certqa.cn",
	//
	//     "AllDomain": "tw.certqa.cn",
	//
	//     "BuyDate": 1773906251000
	//
	//   }
	//
	// ]
	TrusteeOrderList *string `json:"TrusteeOrderList,omitempty" xml:"TrusteeOrderList,omitempty"`
}

func (s ListTrusteeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrusteeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrusteeOrderResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTrusteeOrderResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTrusteeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrusteeOrderResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTrusteeOrderResponseBody) GetTrusteeOrderList() *string {
	return s.TrusteeOrderList
}

func (s *ListTrusteeOrderResponseBody) SetMaxResults(v int32) *ListTrusteeOrderResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTrusteeOrderResponseBody) SetNextToken(v string) *ListTrusteeOrderResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTrusteeOrderResponseBody) SetRequestId(v string) *ListTrusteeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrusteeOrderResponseBody) SetTotalCount(v int32) *ListTrusteeOrderResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrusteeOrderResponseBody) SetTrusteeOrderList(v string) *ListTrusteeOrderResponseBody {
	s.TrusteeOrderList = &v
	return s
}

func (s *ListTrusteeOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
