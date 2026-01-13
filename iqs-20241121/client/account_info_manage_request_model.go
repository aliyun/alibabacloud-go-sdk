// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountInfoManageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *AccountInfoManageRequest
	GetAccountId() *string
	SetName(v string) *AccountInfoManageRequest
	GetName() *string
	SetQuarkKey(v string) *AccountInfoManageRequest
	GetQuarkKey() *string
	SetQuarkName(v string) *AccountInfoManageRequest
	GetQuarkName() *string
	SetTestQps(v int32) *AccountInfoManageRequest
	GetTestQps() *int32
	SetTestQueryPerDay(v int32) *AccountInfoManageRequest
	GetTestQueryPerDay() *int32
}

type AccountInfoManageRequest struct {
	AccountId       *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	QuarkKey        *string `json:"quarkKey,omitempty" xml:"quarkKey,omitempty"`
	QuarkName       *string `json:"quarkName,omitempty" xml:"quarkName,omitempty"`
	TestQps         *int32  `json:"testQps,omitempty" xml:"testQps,omitempty"`
	TestQueryPerDay *int32  `json:"testQueryPerDay,omitempty" xml:"testQueryPerDay,omitempty"`
}

func (s AccountInfoManageRequest) String() string {
	return dara.Prettify(s)
}

func (s AccountInfoManageRequest) GoString() string {
	return s.String()
}

func (s *AccountInfoManageRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *AccountInfoManageRequest) GetName() *string {
	return s.Name
}

func (s *AccountInfoManageRequest) GetQuarkKey() *string {
	return s.QuarkKey
}

func (s *AccountInfoManageRequest) GetQuarkName() *string {
	return s.QuarkName
}

func (s *AccountInfoManageRequest) GetTestQps() *int32 {
	return s.TestQps
}

func (s *AccountInfoManageRequest) GetTestQueryPerDay() *int32 {
	return s.TestQueryPerDay
}

func (s *AccountInfoManageRequest) SetAccountId(v string) *AccountInfoManageRequest {
	s.AccountId = &v
	return s
}

func (s *AccountInfoManageRequest) SetName(v string) *AccountInfoManageRequest {
	s.Name = &v
	return s
}

func (s *AccountInfoManageRequest) SetQuarkKey(v string) *AccountInfoManageRequest {
	s.QuarkKey = &v
	return s
}

func (s *AccountInfoManageRequest) SetQuarkName(v string) *AccountInfoManageRequest {
	s.QuarkName = &v
	return s
}

func (s *AccountInfoManageRequest) SetTestQps(v int32) *AccountInfoManageRequest {
	s.TestQps = &v
	return s
}

func (s *AccountInfoManageRequest) SetTestQueryPerDay(v int32) *AccountInfoManageRequest {
	s.TestQueryPerDay = &v
	return s
}

func (s *AccountInfoManageRequest) Validate() error {
	return dara.Validate(s)
}
