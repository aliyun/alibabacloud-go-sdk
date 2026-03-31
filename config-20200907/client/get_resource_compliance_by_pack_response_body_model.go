// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceByPackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceComplianceByPackResponseBody
	GetRequestId() *string
	SetResourceComplianceResult(v *GetResourceComplianceByPackResponseBodyResourceComplianceResult) *GetResourceComplianceByPackResponseBody
	GetResourceComplianceResult() *GetResourceComplianceByPackResponseBodyResourceComplianceResult
}

type GetResourceComplianceByPackResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The compliance evaluation results returned.
	ResourceComplianceResult *GetResourceComplianceByPackResponseBodyResourceComplianceResult `json:"ResourceComplianceResult,omitempty" xml:"ResourceComplianceResult,omitempty" type:"Struct"`
}

func (s GetResourceComplianceByPackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceByPackResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByPackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceComplianceByPackResponseBody) GetResourceComplianceResult() *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	return s.ResourceComplianceResult
}

func (s *GetResourceComplianceByPackResponseBody) SetRequestId(v string) *GetResourceComplianceByPackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBody) SetResourceComplianceResult(v *GetResourceComplianceByPackResponseBodyResourceComplianceResult) *GetResourceComplianceByPackResponseBody {
	s.ResourceComplianceResult = v
	return s
}

func (s *GetResourceComplianceByPackResponseBody) Validate() error {
	if s.ResourceComplianceResult != nil {
		if err := s.ResourceComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceComplianceByPackResponseBodyResourceComplianceResult struct {
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-541e626622af0087****
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

func (s GetResourceComplianceByPackResponseBodyResourceComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceByPackResponseBodyResourceComplianceResult) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) GetCompliantCount() *int32 {
	return s.CompliantCount
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) GetIgnoredCount() *int32 {
	return s.IgnoredCount
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) GetInsufficientDataCount() *int32 {
	return s.InsufficientDataCount
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) GetNotApplicableCount() *int32 {
	return s.NotApplicableCount
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) SetCompliancePackId(v string) *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.CompliancePackId = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) SetCompliantCount(v int32) *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.CompliantCount = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) SetIgnoredCount(v int32) *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.IgnoredCount = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) SetInsufficientDataCount(v int32) *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.InsufficientDataCount = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) SetNonCompliantCount(v int32) *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.NonCompliantCount = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) SetNotApplicableCount(v int32) *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.NotApplicableCount = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) SetTotalCount(v int32) *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.TotalCount = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) Validate() error {
	return dara.Validate(s)
}
