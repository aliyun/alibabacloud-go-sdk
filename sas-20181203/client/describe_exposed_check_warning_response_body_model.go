// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedCheckWarningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeExposedCheckWarningResponseBody
	GetCount() *int32
	SetRequestId(v string) *DescribeExposedCheckWarningResponseBody
	GetRequestId() *string
	SetWarningList(v []*DescribeExposedCheckWarningResponseBodyWarningList) *DescribeExposedCheckWarningResponseBody
	GetWarningList() []*DescribeExposedCheckWarningResponseBodyWarningList
}

type DescribeExposedCheckWarningResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6D9CDB47-6191-4415-BE63-7E8B12CD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the baseline risk items of the exposed server.
	WarningList []*DescribeExposedCheckWarningResponseBodyWarningList `json:"WarningList,omitempty" xml:"WarningList,omitempty" type:"Repeated"`
}

func (s DescribeExposedCheckWarningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedCheckWarningResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposedCheckWarningResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeExposedCheckWarningResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExposedCheckWarningResponseBody) GetWarningList() []*DescribeExposedCheckWarningResponseBodyWarningList {
	return s.WarningList
}

func (s *DescribeExposedCheckWarningResponseBody) SetCount(v int32) *DescribeExposedCheckWarningResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeExposedCheckWarningResponseBody) SetRequestId(v string) *DescribeExposedCheckWarningResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExposedCheckWarningResponseBody) SetWarningList(v []*DescribeExposedCheckWarningResponseBodyWarningList) *DescribeExposedCheckWarningResponseBody {
	s.WarningList = v
	return s
}

func (s *DescribeExposedCheckWarningResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExposedCheckWarningResponseBodyWarningList struct {
	// The ID of the baseline.
	//
	// >  You can call the [DescribeCheckWarningSummary](https://help.aliyun.com/document_detail/116179.html) operation to query the IDs of baselines.
	//
	// example:
	//
	// 107
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// Weak password-Redis DB login weak password baseline
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	// The display name of the baseline sub type.
	//
	// example:
	//
	// Redis DB login weak password baseline
	SubTypeAlias *string `json:"SubTypeAlias,omitempty" xml:"SubTypeAlias,omitempty"`
	// The display name of the baseline type.
	//
	// example:
	//
	// Weak password
	TypeAlias *string `json:"TypeAlias,omitempty" xml:"TypeAlias,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 1d35b031-ee4e-4e53-8b53-465ab712****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExposedCheckWarningResponseBodyWarningList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedCheckWarningResponseBodyWarningList) GoString() string {
	return s.String()
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) GetRiskId() *int64 {
	return s.RiskId
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) GetRiskName() *string {
	return s.RiskName
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) GetSubTypeAlias() *string {
	return s.SubTypeAlias
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) GetTypeAlias() *string {
	return s.TypeAlias
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) SetRiskId(v int64) *DescribeExposedCheckWarningResponseBodyWarningList {
	s.RiskId = &v
	return s
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) SetRiskName(v string) *DescribeExposedCheckWarningResponseBodyWarningList {
	s.RiskName = &v
	return s
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) SetSubTypeAlias(v string) *DescribeExposedCheckWarningResponseBodyWarningList {
	s.SubTypeAlias = &v
	return s
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) SetTypeAlias(v string) *DescribeExposedCheckWarningResponseBodyWarningList {
	s.TypeAlias = &v
	return s
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) SetUuid(v string) *DescribeExposedCheckWarningResponseBodyWarningList {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedCheckWarningResponseBodyWarningList) Validate() error {
	return dara.Validate(s)
}
