// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucEnterprise interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *BucEnterprise
	GetAccountId() *string
	SetCodeupEnterpriseId(v string) *BucEnterprise
	GetCodeupEnterpriseId() *string
	SetCodeupNamespaceId(v int64) *BucEnterprise
	GetCodeupNamespaceId() *int64
	SetDescription(v string) *BucEnterprise
	GetDescription() *string
	SetId(v int64) *BucEnterprise
	GetId() *int64
	SetName(v string) *BucEnterprise
	GetName() *string
	SetType(v string) *BucEnterprise
	GetType() *string
}

type BucEnterprise struct {
	AccountId          *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CodeupEnterpriseId *string `json:"codeupEnterpriseId,omitempty" xml:"codeupEnterpriseId,omitempty"`
	CodeupNamespaceId  *int64  `json:"codeupNamespaceId,omitempty" xml:"codeupNamespaceId,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	Id                 *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BucEnterprise) String() string {
	return dara.Prettify(s)
}

func (s BucEnterprise) GoString() string {
	return s.String()
}

func (s *BucEnterprise) GetAccountId() *string {
	return s.AccountId
}

func (s *BucEnterprise) GetCodeupEnterpriseId() *string {
	return s.CodeupEnterpriseId
}

func (s *BucEnterprise) GetCodeupNamespaceId() *int64 {
	return s.CodeupNamespaceId
}

func (s *BucEnterprise) GetDescription() *string {
	return s.Description
}

func (s *BucEnterprise) GetId() *int64 {
	return s.Id
}

func (s *BucEnterprise) GetName() *string {
	return s.Name
}

func (s *BucEnterprise) GetType() *string {
	return s.Type
}

func (s *BucEnterprise) SetAccountId(v string) *BucEnterprise {
	s.AccountId = &v
	return s
}

func (s *BucEnterprise) SetCodeupEnterpriseId(v string) *BucEnterprise {
	s.CodeupEnterpriseId = &v
	return s
}

func (s *BucEnterprise) SetCodeupNamespaceId(v int64) *BucEnterprise {
	s.CodeupNamespaceId = &v
	return s
}

func (s *BucEnterprise) SetDescription(v string) *BucEnterprise {
	s.Description = &v
	return s
}

func (s *BucEnterprise) SetId(v int64) *BucEnterprise {
	s.Id = &v
	return s
}

func (s *BucEnterprise) SetName(v string) *BucEnterprise {
	s.Name = &v
	return s
}

func (s *BucEnterprise) SetType(v string) *BucEnterprise {
	s.Type = &v
	return s
}

func (s *BucEnterprise) Validate() error {
	return dara.Validate(s)
}
