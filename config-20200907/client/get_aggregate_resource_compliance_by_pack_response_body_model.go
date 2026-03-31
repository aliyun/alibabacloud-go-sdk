// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceByPackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAggregateResourceComplianceByPackResponseBody
	GetRequestId() *string
	SetResourceComplianceResult(v *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) *GetAggregateResourceComplianceByPackResponseBody
	GetResourceComplianceResult() *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult
}

type GetAggregateResourceComplianceByPackResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The compliance evaluation results returned.
	ResourceComplianceResult *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult `json:"ResourceComplianceResult,omitempty" xml:"ResourceComplianceResult,omitempty" type:"Struct"`
}

func (s GetAggregateResourceComplianceByPackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceByPackResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByPackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateResourceComplianceByPackResponseBody) GetResourceComplianceResult() *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	return s.ResourceComplianceResult
}

func (s *GetAggregateResourceComplianceByPackResponseBody) SetRequestId(v string) *GetAggregateResourceComplianceByPackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBody) SetResourceComplianceResult(v *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) *GetAggregateResourceComplianceByPackResponseBody {
	s.ResourceComplianceResult = v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBody) Validate() error {
	if s.ResourceComplianceResult != nil {
		if err := s.ResourceComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult struct {
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId      *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliantCount        *int32  `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	IgnoredCount          *int32  `json:"IgnoredCount,omitempty" xml:"IgnoredCount,omitempty"`
	InsufficientDataCount *int32  `json:"InsufficientDataCount,omitempty" xml:"InsufficientDataCount,omitempty"`
	// The number of non-compliant resources.
	//
	// example:
	//
	// 7
	NonCompliantCount  *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	NotApplicableCount *int32 `json:"NotApplicableCount,omitempty" xml:"NotApplicableCount,omitempty"`
	// The total number of resources.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) GetCompliantCount() *int32 {
	return s.CompliantCount
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) GetIgnoredCount() *int32 {
	return s.IgnoredCount
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) GetInsufficientDataCount() *int32 {
	return s.InsufficientDataCount
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) GetNotApplicableCount() *int32 {
	return s.NotApplicableCount
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) SetCompliancePackId(v string) *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) SetCompliantCount(v int32) *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.CompliantCount = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) SetIgnoredCount(v int32) *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.IgnoredCount = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) SetInsufficientDataCount(v int32) *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.InsufficientDataCount = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) SetNonCompliantCount(v int32) *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.NonCompliantCount = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) SetNotApplicableCount(v int32) *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.NotApplicableCount = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) SetTotalCount(v int32) *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.TotalCount = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) Validate() error {
	return dara.Validate(s)
}
