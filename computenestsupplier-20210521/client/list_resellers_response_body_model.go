// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResellersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResellersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListResellersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResellersResponseBody
	GetRequestId() *string
	SetSupplierInformation(v []*ListResellersResponseBodySupplierInformation) *ListResellersResponseBody
	GetSupplierInformation() []*ListResellersResponseBodySupplierInformation
	SetTotalCount(v int32) *ListResellersResponseBody
	GetTotalCount() *int32
}

type ListResellersResponseBody struct {
	// The number of entries returned per page. The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A361BA9E-xxxx-xxxx-xxxx-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of results.
	SupplierInformation []*ListResellersResponseBodySupplierInformation `json:"SupplierInformation,omitempty" xml:"SupplierInformation,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResellersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResellersResponseBody) GoString() string {
	return s.String()
}

func (s *ListResellersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResellersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResellersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResellersResponseBody) GetSupplierInformation() []*ListResellersResponseBodySupplierInformation {
	return s.SupplierInformation
}

func (s *ListResellersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResellersResponseBody) SetMaxResults(v int32) *ListResellersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResellersResponseBody) SetNextToken(v string) *ListResellersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResellersResponseBody) SetRequestId(v string) *ListResellersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResellersResponseBody) SetSupplierInformation(v []*ListResellersResponseBodySupplierInformation) *ListResellersResponseBody {
	s.SupplierInformation = v
	return s
}

func (s *ListResellersResponseBody) SetTotalCount(v int32) *ListResellersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResellersResponseBody) Validate() error {
	if s.SupplierInformation != nil {
		for _, item := range s.SupplierInformation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResellersResponseBodySupplierInformation struct {
	// The description of the reseller.
	//
	// example:
	//
	// It is a XXXX company.
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The icon of the reseller.
	//
	// example:
	//
	// https://service-info-public.oss-cn-hangzhou.aliyuncs.com/31978070/service-image/d5c3b585-ff6b-4e4e-8885-xxxx.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The name of the reseller.
	//
	// example:
	//
	// Distributora
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The UID of the reseller.
	//
	// example:
	//
	// 152xxxxxxxxxxx
	SupplierUid *int64 `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The URL of the reseller\\"s official website.
	//
	// example:
	//
	// http://www.aliyun.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
}

func (s ListResellersResponseBodySupplierInformation) String() string {
	return dara.Prettify(s)
}

func (s ListResellersResponseBodySupplierInformation) GoString() string {
	return s.String()
}

func (s *ListResellersResponseBodySupplierInformation) GetSupplierDesc() *string {
	return s.SupplierDesc
}

func (s *ListResellersResponseBodySupplierInformation) GetSupplierLogo() *string {
	return s.SupplierLogo
}

func (s *ListResellersResponseBodySupplierInformation) GetSupplierName() *string {
	return s.SupplierName
}

func (s *ListResellersResponseBodySupplierInformation) GetSupplierUid() *int64 {
	return s.SupplierUid
}

func (s *ListResellersResponseBodySupplierInformation) GetSupplierUrl() *string {
	return s.SupplierUrl
}

func (s *ListResellersResponseBodySupplierInformation) SetSupplierDesc(v string) *ListResellersResponseBodySupplierInformation {
	s.SupplierDesc = &v
	return s
}

func (s *ListResellersResponseBodySupplierInformation) SetSupplierLogo(v string) *ListResellersResponseBodySupplierInformation {
	s.SupplierLogo = &v
	return s
}

func (s *ListResellersResponseBodySupplierInformation) SetSupplierName(v string) *ListResellersResponseBodySupplierInformation {
	s.SupplierName = &v
	return s
}

func (s *ListResellersResponseBodySupplierInformation) SetSupplierUid(v int64) *ListResellersResponseBodySupplierInformation {
	s.SupplierUid = &v
	return s
}

func (s *ListResellersResponseBodySupplierInformation) SetSupplierUrl(v string) *ListResellersResponseBodySupplierInformation {
	s.SupplierUrl = &v
	return s
}

func (s *ListResellersResponseBodySupplierInformation) Validate() error {
	return dara.Validate(s)
}
