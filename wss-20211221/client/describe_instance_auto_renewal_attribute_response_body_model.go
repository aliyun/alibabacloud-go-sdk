// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewalAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetCount() *int32
	SetInstanceAttribute(v []*DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetInstanceAttribute() []*DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute
	SetMaxResults(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetNextToken() *string
	SetPageNo(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetRequestId() *string
}

type DescribeInstanceAutoRenewalAttributeResponseBody struct {
	// The total number of entries that meet the query conditions.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of instances and their auto-renewal status.
	InstanceAttribute []*DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute `json:"InstanceAttribute,omitempty" xml:"InstanceAttribute,omitempty" type:"Repeated"`
	// The maximum number of entries to return per query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The mark for paging. Leave this parameter empty for the first query. For subsequent paging, pass the NextToken value returned in the previous response as is. Do not construct or decode the value.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6k=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number for paging, starting from 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 72481C12-69AB-5CE1-8A35-A8EFA921****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetInstanceAttribute() []*DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute {
	return s.InstanceAttribute
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetCount(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetInstanceAttribute(v []*DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.InstanceAttribute = v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetMaxResults(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetNextToken(v string) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetPageNo(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetPageSize(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) Validate() error {
	if s.InstanceAttribute != nil {
		for _, item := range s.InstanceAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 123456789012****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The auto-renewal status. Valid values: true (auto-renewal enabled) and false (auto-renewal disabled).
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The instance ID, which is used to identify the purchased instance.
	//
	// example:
	//
	// cp-0bw2ew49****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) SetAliUid(v int64) *DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute {
	s.AliUid = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) SetAutoRenewal(v bool) *DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyInstanceAttribute) Validate() error {
	return dara.Validate(s)
}
