// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiSignaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiSignatures(v *DescribeApiSignaturesResponseBodyApiSignatures) *DescribeApiSignaturesResponseBody
	GetApiSignatures() *DescribeApiSignaturesResponseBodyApiSignatures
	SetPageNumber(v int32) *DescribeApiSignaturesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiSignaturesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApiSignaturesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApiSignaturesResponseBody
	GetTotalCount() *int32
}

type DescribeApiSignaturesResponseBody struct {
	// The returned signature key information. It is an array consisting of ApiSignatureItem data.
	ApiSignatures *DescribeApiSignaturesResponseBodyApiSignatures `json:"ApiSignatures,omitempty" xml:"ApiSignatures,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiSignaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiSignaturesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiSignaturesResponseBody) GetApiSignatures() *DescribeApiSignaturesResponseBodyApiSignatures {
	return s.ApiSignatures
}

func (s *DescribeApiSignaturesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiSignaturesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiSignaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiSignaturesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApiSignaturesResponseBody) SetApiSignatures(v *DescribeApiSignaturesResponseBodyApiSignatures) *DescribeApiSignaturesResponseBody {
	s.ApiSignatures = v
	return s
}

func (s *DescribeApiSignaturesResponseBody) SetPageNumber(v int32) *DescribeApiSignaturesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiSignaturesResponseBody) SetPageSize(v int32) *DescribeApiSignaturesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiSignaturesResponseBody) SetRequestId(v string) *DescribeApiSignaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiSignaturesResponseBody) SetTotalCount(v int32) *DescribeApiSignaturesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApiSignaturesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApiSignaturesResponseBodyApiSignatures struct {
	ApiSignatureItem []*DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem `json:"ApiSignatureItem,omitempty" xml:"ApiSignatureItem,omitempty" type:"Repeated"`
}

func (s DescribeApiSignaturesResponseBodyApiSignatures) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiSignaturesResponseBodyApiSignatures) GoString() string {
	return s.String()
}

func (s *DescribeApiSignaturesResponseBodyApiSignatures) GetApiSignatureItem() []*DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	return s.ApiSignatureItem
}

func (s *DescribeApiSignaturesResponseBodyApiSignatures) SetApiSignatureItem(v []*DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) *DescribeApiSignaturesResponseBodyApiSignatures {
	s.ApiSignatureItem = v
	return s
}

func (s *DescribeApiSignaturesResponseBodyApiSignatures) Validate() error {
	return dara.Validate(s)
}

type DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem struct {
	// The ID of the API.
	//
	// example:
	//
	// 46fbb52840d146f186e38e8e70fc8c90
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// testapi
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The time when the backend signature key was bound.
	//
	// example:
	//
	// 2016-07-23T08:28:48Z
	BoundTime *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	// The ID of the backend signature key.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// The name of the backend signature key.
	//
	// example:
	//
	// backendsignature
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) GoString() string {
	return s.String()
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) GetBoundTime() *string {
	return s.BoundTime
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) GetSignatureId() *string {
	return s.SignatureId
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) GetSignatureName() *string {
	return s.SignatureName
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) SetApiId(v string) *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	s.ApiId = &v
	return s
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) SetApiName(v string) *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	s.ApiName = &v
	return s
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) SetBoundTime(v string) *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	s.BoundTime = &v
	return s
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) SetSignatureId(v string) *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	s.SignatureId = &v
	return s
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) SetSignatureName(v string) *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	s.SignatureName = &v
	return s
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) Validate() error {
	return dara.Validate(s)
}
