// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSignaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSignaturesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSignaturesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSignaturesResponseBody
	GetRequestId() *string
	SetSignatureInfos(v *DescribeSignaturesResponseBodySignatureInfos) *DescribeSignaturesResponseBody
	GetSignatureInfos() *DescribeSignaturesResponseBodySignatureInfos
	SetTotalCount(v int32) *DescribeSignaturesResponseBody
	GetTotalCount() *int32
}

type DescribeSignaturesResponseBody struct {
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
	// The returned signature information. It is an array consisting of SignatureInfo data.
	SignatureInfos *DescribeSignaturesResponseBodySignatureInfos `json:"SignatureInfos,omitempty" xml:"SignatureInfos,omitempty" type:"Struct"`
	// The total number of returned entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSignaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignaturesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSignaturesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSignaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSignaturesResponseBody) GetSignatureInfos() *DescribeSignaturesResponseBodySignatureInfos {
	return s.SignatureInfos
}

func (s *DescribeSignaturesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSignaturesResponseBody) SetPageNumber(v int32) *DescribeSignaturesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSignaturesResponseBody) SetPageSize(v int32) *DescribeSignaturesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSignaturesResponseBody) SetRequestId(v string) *DescribeSignaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSignaturesResponseBody) SetSignatureInfos(v *DescribeSignaturesResponseBodySignatureInfos) *DescribeSignaturesResponseBody {
	s.SignatureInfos = v
	return s
}

func (s *DescribeSignaturesResponseBody) SetTotalCount(v int32) *DescribeSignaturesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSignaturesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSignaturesResponseBodySignatureInfos struct {
	SignatureInfo []*DescribeSignaturesResponseBodySignatureInfosSignatureInfo `json:"SignatureInfo,omitempty" xml:"SignatureInfo,omitempty" type:"Repeated"`
}

func (s DescribeSignaturesResponseBodySignatureInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignaturesResponseBodySignatureInfos) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesResponseBodySignatureInfos) GetSignatureInfo() []*DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	return s.SignatureInfo
}

func (s *DescribeSignaturesResponseBodySignatureInfos) SetSignatureInfo(v []*DescribeSignaturesResponseBodySignatureInfosSignatureInfo) *DescribeSignaturesResponseBodySignatureInfos {
	s.SignatureInfo = v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeSignaturesResponseBodySignatureInfosSignatureInfo struct {
	// The creation time of the key.
	//
	// example:
	//
	// 2016-07-23T08:28:48Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The last modification time of the key.
	//
	// example:
	//
	// 2016-07-24T08:28:48Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The region where the key is located.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the backend signature key.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// The Key value of the backend signature key.
	//
	// example:
	//
	// qwertyuiop
	SignatureKey *string `json:"SignatureKey,omitempty" xml:"SignatureKey,omitempty"`
	// The name of the backend signature key.
	//
	// example:
	//
	// backendsignature
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
	// The Secret value of the backend signature key.
	//
	// example:
	//
	// asdfghjkl
	SignatureSecret *string `json:"SignatureSecret,omitempty" xml:"SignatureSecret,omitempty"`
}

func (s DescribeSignaturesResponseBodySignatureInfosSignatureInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignaturesResponseBodySignatureInfosSignatureInfo) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) GetSignatureId() *string {
	return s.SignatureId
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) GetSignatureKey() *string {
	return s.SignatureKey
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) GetSignatureName() *string {
	return s.SignatureName
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) GetSignatureSecret() *string {
	return s.SignatureSecret
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetCreatedTime(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetModifiedTime(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetRegionId(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetSignatureId(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.SignatureId = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetSignatureKey(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.SignatureKey = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetSignatureName(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.SignatureName = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetSignatureSecret(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.SignatureSecret = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) Validate() error {
	return dara.Validate(s)
}
