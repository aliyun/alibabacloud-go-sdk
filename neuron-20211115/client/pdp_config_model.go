// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *PdpConfig
	GetAccountId() *string
	SetAccountName(v string) *PdpConfig
	GetAccountName() *string
	SetAskId(v string) *PdpConfig
	GetAskId() *string
	SetContext(v string) *PdpConfig
	GetContext() *string
	SetEnterpriseId(v int64) *PdpConfig
	GetEnterpriseId() *int64
	SetEnv(v string) *PdpConfig
	GetEnv() *string
	SetGmtCreate(v string) *PdpConfig
	GetGmtCreate() *string
	SetGmtModified(v string) *PdpConfig
	GetGmtModified() *string
	SetId(v int64) *PdpConfig
	GetId() *int64
	SetName(v string) *PdpConfig
	GetName() *string
	SetPbcId(v int64) *PdpConfig
	GetPbcId() *int64
	SetRequestId(v string) *PdpConfig
	GetRequestId() *string
	SetServiceGroupId(v int64) *PdpConfig
	GetServiceGroupId() *int64
	SetServiceId(v int64) *PdpConfig
	GetServiceId() *int64
	SetType(v string) *PdpConfig
	GetType() *string
}

type PdpConfig struct {
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
	// {key:value}
	Context *string `json:"context,omitempty" xml:"context,omitempty"`
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
	// example:
	//
	// 2022-2-22 11:11:2
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// neuron-developer-spring-config
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PbcId     *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SERVICE_SYSTEM
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PdpConfig) String() string {
	return dara.Prettify(s)
}

func (s PdpConfig) GoString() string {
	return s.String()
}

func (s *PdpConfig) GetAccountId() *string {
	return s.AccountId
}

func (s *PdpConfig) GetAccountName() *string {
	return s.AccountName
}

func (s *PdpConfig) GetAskId() *string {
	return s.AskId
}

func (s *PdpConfig) GetContext() *string {
	return s.Context
}

func (s *PdpConfig) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *PdpConfig) GetEnv() *string {
	return s.Env
}

func (s *PdpConfig) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PdpConfig) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PdpConfig) GetId() *int64 {
	return s.Id
}

func (s *PdpConfig) GetName() *string {
	return s.Name
}

func (s *PdpConfig) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PdpConfig) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpConfig) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *PdpConfig) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *PdpConfig) GetType() *string {
	return s.Type
}

func (s *PdpConfig) SetAccountId(v string) *PdpConfig {
	s.AccountId = &v
	return s
}

func (s *PdpConfig) SetAccountName(v string) *PdpConfig {
	s.AccountName = &v
	return s
}

func (s *PdpConfig) SetAskId(v string) *PdpConfig {
	s.AskId = &v
	return s
}

func (s *PdpConfig) SetContext(v string) *PdpConfig {
	s.Context = &v
	return s
}

func (s *PdpConfig) SetEnterpriseId(v int64) *PdpConfig {
	s.EnterpriseId = &v
	return s
}

func (s *PdpConfig) SetEnv(v string) *PdpConfig {
	s.Env = &v
	return s
}

func (s *PdpConfig) SetGmtCreate(v string) *PdpConfig {
	s.GmtCreate = &v
	return s
}

func (s *PdpConfig) SetGmtModified(v string) *PdpConfig {
	s.GmtModified = &v
	return s
}

func (s *PdpConfig) SetId(v int64) *PdpConfig {
	s.Id = &v
	return s
}

func (s *PdpConfig) SetName(v string) *PdpConfig {
	s.Name = &v
	return s
}

func (s *PdpConfig) SetPbcId(v int64) *PdpConfig {
	s.PbcId = &v
	return s
}

func (s *PdpConfig) SetRequestId(v string) *PdpConfig {
	s.RequestId = &v
	return s
}

func (s *PdpConfig) SetServiceGroupId(v int64) *PdpConfig {
	s.ServiceGroupId = &v
	return s
}

func (s *PdpConfig) SetServiceId(v int64) *PdpConfig {
	s.ServiceId = &v
	return s
}

func (s *PdpConfig) SetType(v string) *PdpConfig {
	s.Type = &v
	return s
}

func (s *PdpConfig) Validate() error {
	return dara.Validate(s)
}
