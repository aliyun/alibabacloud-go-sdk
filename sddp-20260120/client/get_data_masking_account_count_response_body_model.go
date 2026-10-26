// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataMaskingAccountCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountCount(v *GetDataMaskingAccountCountResponseBodyAccountCount) *GetDataMaskingAccountCountResponseBody
	GetAccountCount() *GetDataMaskingAccountCountResponseBodyAccountCount
	SetRequestId(v string) *GetDataMaskingAccountCountResponseBody
	GetRequestId() *string
}

type GetDataMaskingAccountCountResponseBody struct {
	// The statistics on the number of account permissions for column encryption.
	AccountCount *GetDataMaskingAccountCountResponseBodyAccountCount `json:"AccountCount,omitempty" xml:"AccountCount,omitempty" type:"Struct"`
	// The request ID, which can be used to locate the API call.
	//
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataMaskingAccountCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataMaskingAccountCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataMaskingAccountCountResponseBody) GetAccountCount() *GetDataMaskingAccountCountResponseBodyAccountCount {
	return s.AccountCount
}

func (s *GetDataMaskingAccountCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataMaskingAccountCountResponseBody) SetAccountCount(v *GetDataMaskingAccountCountResponseBodyAccountCount) *GetDataMaskingAccountCountResponseBody {
	s.AccountCount = v
	return s
}

func (s *GetDataMaskingAccountCountResponseBody) SetRequestId(v string) *GetDataMaskingAccountCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataMaskingAccountCountResponseBody) Validate() error {
	if s.AccountCount != nil {
		if err := s.AccountCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataMaskingAccountCountResponseBodyAccountCount struct {
	// The number of accounts with plaintext access (fullAccess).
	//
	// example:
	//
	// 2
	FullAccessCount *int64 `json:"FullAccessCount,omitempty" xml:"FullAccessCount,omitempty"`
	// The number of accounts prohibited from decryption (noneAccess).
	//
	// example:
	//
	// 1
	NoneAccessCount *int64 `json:"NoneAccessCount,omitempty" xml:"NoneAccessCount,omitempty"`
	// The number of accounts with restricted access (restrictedAccess).
	//
	// example:
	//
	// 3
	RestrictedAccessCount *int64 `json:"RestrictedAccessCount,omitempty" xml:"RestrictedAccessCount,omitempty"`
	// The total number of database accounts.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDataMaskingAccountCountResponseBodyAccountCount) String() string {
	return dara.Prettify(s)
}

func (s GetDataMaskingAccountCountResponseBodyAccountCount) GoString() string {
	return s.String()
}

func (s *GetDataMaskingAccountCountResponseBodyAccountCount) GetFullAccessCount() *int64 {
	return s.FullAccessCount
}

func (s *GetDataMaskingAccountCountResponseBodyAccountCount) GetNoneAccessCount() *int64 {
	return s.NoneAccessCount
}

func (s *GetDataMaskingAccountCountResponseBodyAccountCount) GetRestrictedAccessCount() *int64 {
	return s.RestrictedAccessCount
}

func (s *GetDataMaskingAccountCountResponseBodyAccountCount) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetDataMaskingAccountCountResponseBodyAccountCount) SetFullAccessCount(v int64) *GetDataMaskingAccountCountResponseBodyAccountCount {
	s.FullAccessCount = &v
	return s
}

func (s *GetDataMaskingAccountCountResponseBodyAccountCount) SetNoneAccessCount(v int64) *GetDataMaskingAccountCountResponseBodyAccountCount {
	s.NoneAccessCount = &v
	return s
}

func (s *GetDataMaskingAccountCountResponseBodyAccountCount) SetRestrictedAccessCount(v int64) *GetDataMaskingAccountCountResponseBodyAccountCount {
	s.RestrictedAccessCount = &v
	return s
}

func (s *GetDataMaskingAccountCountResponseBodyAccountCount) SetTotalCount(v int64) *GetDataMaskingAccountCountResponseBodyAccountCount {
	s.TotalCount = &v
	return s
}

func (s *GetDataMaskingAccountCountResponseBodyAccountCount) Validate() error {
	return dara.Validate(s)
}
