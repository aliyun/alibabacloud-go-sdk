// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpHistoryConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *PdpHistoryConfig
	GetAccountId() *string
	SetAccountName(v string) *PdpHistoryConfig
	GetAccountName() *string
	SetAskId(v string) *PdpHistoryConfig
	GetAskId() *string
	SetConfigId(v int64) *PdpHistoryConfig
	GetConfigId() *int64
	SetContext(v string) *PdpHistoryConfig
	GetContext() *string
	SetEnterpriseId(v int64) *PdpHistoryConfig
	GetEnterpriseId() *int64
	SetEnv(v string) *PdpHistoryConfig
	GetEnv() *string
	SetGmtCreate(v string) *PdpHistoryConfig
	GetGmtCreate() *string
	SetId(v int64) *PdpHistoryConfig
	GetId() *int64
	SetName(v string) *PdpHistoryConfig
	GetName() *string
	SetPbcId(v int64) *PdpHistoryConfig
	GetPbcId() *int64
	SetRequestId(v string) *PdpHistoryConfig
	GetRequestId() *string
	SetServiceGroupId(v int64) *PdpHistoryConfig
	GetServiceGroupId() *int64
	SetServiceId(v int64) *PdpHistoryConfig
	GetServiceId() *int64
	SetType(v string) *PdpHistoryConfig
	GetType() *string
}

type PdpHistoryConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// account1
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// 1
	AskId *string `json:"askId,omitempty" xml:"askId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ConfigId *int64 `json:"configId,omitempty" xml:"configId,omitempty"`
	// example:
	//
	// {key:value}
	Context *string `json:"context,omitempty" xml:"context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 2022-2-22 11:11:2
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// neuron-developer-spring-config
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PbcId          *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// SERVICE_SYSTEM
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PdpHistoryConfig) String() string {
	return dara.Prettify(s)
}

func (s PdpHistoryConfig) GoString() string {
	return s.String()
}

func (s *PdpHistoryConfig) GetAccountId() *string {
	return s.AccountId
}

func (s *PdpHistoryConfig) GetAccountName() *string {
	return s.AccountName
}

func (s *PdpHistoryConfig) GetAskId() *string {
	return s.AskId
}

func (s *PdpHistoryConfig) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *PdpHistoryConfig) GetContext() *string {
	return s.Context
}

func (s *PdpHistoryConfig) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *PdpHistoryConfig) GetEnv() *string {
	return s.Env
}

func (s *PdpHistoryConfig) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PdpHistoryConfig) GetId() *int64 {
	return s.Id
}

func (s *PdpHistoryConfig) GetName() *string {
	return s.Name
}

func (s *PdpHistoryConfig) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PdpHistoryConfig) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpHistoryConfig) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *PdpHistoryConfig) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *PdpHistoryConfig) GetType() *string {
	return s.Type
}

func (s *PdpHistoryConfig) SetAccountId(v string) *PdpHistoryConfig {
	s.AccountId = &v
	return s
}

func (s *PdpHistoryConfig) SetAccountName(v string) *PdpHistoryConfig {
	s.AccountName = &v
	return s
}

func (s *PdpHistoryConfig) SetAskId(v string) *PdpHistoryConfig {
	s.AskId = &v
	return s
}

func (s *PdpHistoryConfig) SetConfigId(v int64) *PdpHistoryConfig {
	s.ConfigId = &v
	return s
}

func (s *PdpHistoryConfig) SetContext(v string) *PdpHistoryConfig {
	s.Context = &v
	return s
}

func (s *PdpHistoryConfig) SetEnterpriseId(v int64) *PdpHistoryConfig {
	s.EnterpriseId = &v
	return s
}

func (s *PdpHistoryConfig) SetEnv(v string) *PdpHistoryConfig {
	s.Env = &v
	return s
}

func (s *PdpHistoryConfig) SetGmtCreate(v string) *PdpHistoryConfig {
	s.GmtCreate = &v
	return s
}

func (s *PdpHistoryConfig) SetId(v int64) *PdpHistoryConfig {
	s.Id = &v
	return s
}

func (s *PdpHistoryConfig) SetName(v string) *PdpHistoryConfig {
	s.Name = &v
	return s
}

func (s *PdpHistoryConfig) SetPbcId(v int64) *PdpHistoryConfig {
	s.PbcId = &v
	return s
}

func (s *PdpHistoryConfig) SetRequestId(v string) *PdpHistoryConfig {
	s.RequestId = &v
	return s
}

func (s *PdpHistoryConfig) SetServiceGroupId(v int64) *PdpHistoryConfig {
	s.ServiceGroupId = &v
	return s
}

func (s *PdpHistoryConfig) SetServiceId(v int64) *PdpHistoryConfig {
	s.ServiceId = &v
	return s
}

func (s *PdpHistoryConfig) SetType(v string) *PdpHistoryConfig {
	s.Type = &v
	return s
}

func (s *PdpHistoryConfig) Validate() error {
	return dara.Validate(s)
}
