// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateAccountComplianceByPackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountComplianceResult(v *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) *GetAggregateAccountComplianceByPackResponseBody
	GetAccountComplianceResult() *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult
	SetRequestId(v string) *GetAggregateAccountComplianceByPackResponseBody
	GetRequestId() *string
}

type GetAggregateAccountComplianceByPackResponseBody struct {
	// The compliance evaluation results of member accounts for which the compliance package takes effect in an account group.
	AccountComplianceResult *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult `json:"AccountComplianceResult,omitempty" xml:"AccountComplianceResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateAccountComplianceByPackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateAccountComplianceByPackResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateAccountComplianceByPackResponseBody) GetAccountComplianceResult() *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult {
	return s.AccountComplianceResult
}

func (s *GetAggregateAccountComplianceByPackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateAccountComplianceByPackResponseBody) SetAccountComplianceResult(v *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) *GetAggregateAccountComplianceByPackResponseBody {
	s.AccountComplianceResult = v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBody) SetRequestId(v string) *GetAggregateAccountComplianceByPackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBody) Validate() error {
	if s.AccountComplianceResult != nil {
		if err := s.AccountComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult struct {
	// The compliance evaluation result of member accounts.
	AccountCompliances []*GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances `json:"AccountCompliances,omitempty" xml:"AccountCompliances,omitempty" type:"Repeated"`
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-541e626622af0087****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The number of non-compliant member accounts.
	//
	// example:
	//
	// 0
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	// The total number of member accounts.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) GetAccountCompliances() []*GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances {
	return s.AccountCompliances
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) GetNonCompliantCount() *int32 {
	return s.NonCompliantCount
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) SetAccountCompliances(v []*GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult {
	s.AccountCompliances = v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) SetCompliancePackId(v string) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) SetNonCompliantCount(v int32) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult {
	s.NonCompliantCount = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) SetTotalCount(v int32) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult {
	s.TotalCount = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) Validate() error {
	if s.AccountCompliances != nil {
		for _, item := range s.AccountCompliances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances struct {
	// The ID of the member account in the account group.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member account in the account group.
	//
	// example:
	//
	// Alice
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The compliance evaluation result. Valid values:
	//
	// 	- COMPLIANT: The resource was evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resource was evaluated as incompliant.
	//
	// 	- NOT_APPLICABLE: The rule did not apply to your resource.
	//
	// 	- INSUFFICIENT_DATA: No resource data was available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
}

func (s GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) GoString() string {
	return s.String()
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) GetAccountName() *string {
	return s.AccountName
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) SetAccountId(v int64) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances {
	s.AccountId = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) SetAccountName(v string) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances {
	s.AccountName = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) SetComplianceType(v string) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) Validate() error {
	return dara.Validate(s)
}
