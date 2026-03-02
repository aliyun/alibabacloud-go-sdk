// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpResource interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *PdpResource
	GetAddress() *string
	SetCompanyId(v int64) *PdpResource
	GetCompanyId() *int64
	SetConfiguration(v string) *PdpResource
	GetConfiguration() *string
	SetCredentials(v string) *PdpResource
	GetCredentials() *string
	SetDescription(v string) *PdpResource
	GetDescription() *string
	SetEnv(v string) *PdpResource
	GetEnv() *string
	SetId(v int64) *PdpResource
	GetId() *int64
	SetName(v string) *PdpResource
	GetName() *string
	SetNamespace(v string) *PdpResource
	GetNamespace() *string
	SetProductId(v int64) *PdpResource
	GetProductId() *int64
	SetRequestId(v string) *PdpResource
	GetRequestId() *string
	SetType(v string) *PdpResource
	GetType() *string
	SetUseScope(v string) *PdpResource
	GetUseScope() *string
	SetUseType(v string) *PdpResource
	GetUseType() *string
}

type PdpResource struct {
	// This parameter is required.
	//
	// example:
	//
	// http://
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CompanyId     *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Configuration *string `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {ak:xxxx,sk:xxx}
	Credentials *string `json:"credentials,omitempty" xml:"credentials,omitempty"`
	// example:
	//
	// 用于统一管理资源
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Redis
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Redis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Inner
	UseScope *string `json:"useScope,omitempty" xml:"useScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Share
	UseType *string `json:"useType,omitempty" xml:"useType,omitempty"`
}

func (s PdpResource) String() string {
	return dara.Prettify(s)
}

func (s PdpResource) GoString() string {
	return s.String()
}

func (s *PdpResource) GetAddress() *string {
	return s.Address
}

func (s *PdpResource) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PdpResource) GetConfiguration() *string {
	return s.Configuration
}

func (s *PdpResource) GetCredentials() *string {
	return s.Credentials
}

func (s *PdpResource) GetDescription() *string {
	return s.Description
}

func (s *PdpResource) GetEnv() *string {
	return s.Env
}

func (s *PdpResource) GetId() *int64 {
	return s.Id
}

func (s *PdpResource) GetName() *string {
	return s.Name
}

func (s *PdpResource) GetNamespace() *string {
	return s.Namespace
}

func (s *PdpResource) GetProductId() *int64 {
	return s.ProductId
}

func (s *PdpResource) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpResource) GetType() *string {
	return s.Type
}

func (s *PdpResource) GetUseScope() *string {
	return s.UseScope
}

func (s *PdpResource) GetUseType() *string {
	return s.UseType
}

func (s *PdpResource) SetAddress(v string) *PdpResource {
	s.Address = &v
	return s
}

func (s *PdpResource) SetCompanyId(v int64) *PdpResource {
	s.CompanyId = &v
	return s
}

func (s *PdpResource) SetConfiguration(v string) *PdpResource {
	s.Configuration = &v
	return s
}

func (s *PdpResource) SetCredentials(v string) *PdpResource {
	s.Credentials = &v
	return s
}

func (s *PdpResource) SetDescription(v string) *PdpResource {
	s.Description = &v
	return s
}

func (s *PdpResource) SetEnv(v string) *PdpResource {
	s.Env = &v
	return s
}

func (s *PdpResource) SetId(v int64) *PdpResource {
	s.Id = &v
	return s
}

func (s *PdpResource) SetName(v string) *PdpResource {
	s.Name = &v
	return s
}

func (s *PdpResource) SetNamespace(v string) *PdpResource {
	s.Namespace = &v
	return s
}

func (s *PdpResource) SetProductId(v int64) *PdpResource {
	s.ProductId = &v
	return s
}

func (s *PdpResource) SetRequestId(v string) *PdpResource {
	s.RequestId = &v
	return s
}

func (s *PdpResource) SetType(v string) *PdpResource {
	s.Type = &v
	return s
}

func (s *PdpResource) SetUseScope(v string) *PdpResource {
	s.UseScope = &v
	return s
}

func (s *PdpResource) SetUseType(v string) *PdpResource {
	s.UseType = &v
	return s
}

func (s *PdpResource) Validate() error {
	return dara.Validate(s)
}
