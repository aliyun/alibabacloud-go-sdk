// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSensitiveDomainStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecSensitiveDomainStatisticResponseBodyData) *DescribeApisecSensitiveDomainStatisticResponseBody
	GetData() []*DescribeApisecSensitiveDomainStatisticResponseBodyData
	SetRequestId(v string) *DescribeApisecSensitiveDomainStatisticResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeApisecSensitiveDomainStatisticResponseBody
	GetTotalCount() *int64
}

type DescribeApisecSensitiveDomainStatisticResponseBody struct {
	// The response parameters.
	Data []*DescribeApisecSensitiveDomainStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 56B40D30-4960-4F19-B7D5-2B1F***6CB70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 27
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecSensitiveDomainStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSensitiveDomainStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBody) GetData() []*DescribeApisecSensitiveDomainStatisticResponseBodyData {
	return s.Data
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBody) SetData(v []*DescribeApisecSensitiveDomainStatisticResponseBodyData) *DescribeApisecSensitiveDomainStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBody) SetRequestId(v string) *DescribeApisecSensitiveDomainStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBody) SetTotalCount(v int64) *DescribeApisecSensitiveDomainStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisecSensitiveDomainStatisticResponseBodyData struct {
	// The number of APIs that are involved.
	//
	// example:
	//
	// 10
	ApiCount *int64 `json:"ApiCount,omitempty" xml:"ApiCount,omitempty"`
	// The number of sites that are involved.
	//
	// example:
	//
	// 10
	DomainCount *int64 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The code of the sensitive data.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported sensitive data types.
	//
	// example:
	//
	// 10
	SensitiveCode *string `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
	// The sensitivity level of the sensitive data.Valid values:
	//
	// 	- **S1**: low sensitivity.
	//
	// 	- **S2**: moderate sensitivity.
	//
	// 	- **S3**: high sensitivity.
	//
	// example:
	//
	// L3
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// The name of the sensitive data.
	//
	// example:
	//
	// 1002
	SensitiveName *string `json:"SensitiveName,omitempty" xml:"SensitiveName,omitempty"`
}

func (s DescribeApisecSensitiveDomainStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSensitiveDomainStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) GetApiCount() *int64 {
	return s.ApiCount
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) GetDomainCount() *int64 {
	return s.DomainCount
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) GetSensitiveCode() *string {
	return s.SensitiveCode
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) GetSensitiveName() *string {
	return s.SensitiveName
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) SetApiCount(v int64) *DescribeApisecSensitiveDomainStatisticResponseBodyData {
	s.ApiCount = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) SetDomainCount(v int64) *DescribeApisecSensitiveDomainStatisticResponseBodyData {
	s.DomainCount = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) SetSensitiveCode(v string) *DescribeApisecSensitiveDomainStatisticResponseBodyData {
	s.SensitiveCode = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) SetSensitiveLevel(v string) *DescribeApisecSensitiveDomainStatisticResponseBodyData {
	s.SensitiveLevel = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) SetSensitiveName(v string) *DescribeApisecSensitiveDomainStatisticResponseBodyData {
	s.SensitiveName = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
