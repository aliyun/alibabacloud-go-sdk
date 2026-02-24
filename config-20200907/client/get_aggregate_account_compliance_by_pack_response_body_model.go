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
	AccountComplianceResult *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult `json:"AccountComplianceResult,omitempty" xml:"AccountComplianceResult,omitempty" type:"Struct"`
	RequestId               *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountCompliances []*GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances `json:"AccountCompliances,omitempty" xml:"AccountCompliances,omitempty" type:"Repeated"`
	CompliancePackId   *string                                                                                     `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	NonCompliantCount  *int32                                                                                      `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	TotalCount         *int32                                                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AccountId      *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
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
