// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulListPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeVulListPageResponseBodyData) *DescribeVulListPageResponseBody
	GetData() []*DescribeVulListPageResponseBodyData
	SetRequestId(v string) *DescribeVulListPageResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVulListPageResponseBody
	GetTotalCount() *int32
}

type DescribeVulListPageResponseBody struct {
	// The response parameters.
	Data []*DescribeVulListPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4347E985-6E64-467B-96EC-30D4EA9E32FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVulListPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulListPageResponseBody) GetData() []*DescribeVulListPageResponseBodyData {
	return s.Data
}

func (s *DescribeVulListPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulListPageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVulListPageResponseBody) SetData(v []*DescribeVulListPageResponseBodyData) *DescribeVulListPageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVulListPageResponseBody) SetRequestId(v string) *DescribeVulListPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulListPageResponseBody) SetTotalCount(v int32) *DescribeVulListPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulListPageResponseBody) Validate() error {
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

type DescribeVulListPageResponseBodyData struct {
	// The CVE ID.
	//
	// example:
	//
	// CVE-2022-42836
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The Server Guard extended field.
	//
	// example:
	//
	// {\\"relatedType\\":[{\\"type\\":\\"sys\\"}]}
	ExtAegis *string `json:"ExtAegis,omitempty" xml:"ExtAegis,omitempty"`
	// The primary key ID in the database.
	//
	// example:
	//
	// 40586
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether version comparison is supported. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// example:
	//
	// 1
	IsAegis *int32 `json:"IsAegis,omitempty" xml:"IsAegis,omitempty"`
	// Indicates whether proof-of-concept (PoC) verification is supported. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// example:
	//
	// 0
	IsSas *int32 `json:"IsSas,omitempty" xml:"IsSas,omitempty"`
	// The ID.
	//
	// example:
	//
	// AVD-2018-8218
	OtherId *string `json:"OtherId,omitempty" xml:"OtherId,omitempty"`
	// The release time in UTC (ISO 8601 format), for example, 2022-12-13T08:00Z.
	//
	// example:
	//
	// 2022-12-13T08:00Z
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// Windows Terminal remote code execute vulnerability
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeVulListPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVulListPageResponseBodyData) GetCveId() *string {
	return s.CveId
}

func (s *DescribeVulListPageResponseBodyData) GetExtAegis() *string {
	return s.ExtAegis
}

func (s *DescribeVulListPageResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DescribeVulListPageResponseBodyData) GetIsAegis() *int32 {
	return s.IsAegis
}

func (s *DescribeVulListPageResponseBodyData) GetIsSas() *int32 {
	return s.IsSas
}

func (s *DescribeVulListPageResponseBodyData) GetOtherId() *string {
	return s.OtherId
}

func (s *DescribeVulListPageResponseBodyData) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *DescribeVulListPageResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *DescribeVulListPageResponseBodyData) SetCveId(v string) *DescribeVulListPageResponseBodyData {
	s.CveId = &v
	return s
}

func (s *DescribeVulListPageResponseBodyData) SetExtAegis(v string) *DescribeVulListPageResponseBodyData {
	s.ExtAegis = &v
	return s
}

func (s *DescribeVulListPageResponseBodyData) SetId(v int64) *DescribeVulListPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeVulListPageResponseBodyData) SetIsAegis(v int32) *DescribeVulListPageResponseBodyData {
	s.IsAegis = &v
	return s
}

func (s *DescribeVulListPageResponseBodyData) SetIsSas(v int32) *DescribeVulListPageResponseBodyData {
	s.IsSas = &v
	return s
}

func (s *DescribeVulListPageResponseBodyData) SetOtherId(v string) *DescribeVulListPageResponseBodyData {
	s.OtherId = &v
	return s
}

func (s *DescribeVulListPageResponseBodyData) SetReleaseTime(v int64) *DescribeVulListPageResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeVulListPageResponseBodyData) SetTitle(v string) *DescribeVulListPageResponseBodyData {
	s.Title = &v
	return s
}

func (s *DescribeVulListPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
