// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *ProductUpdateCmd
	GetAlias() *string
	SetCompanyId(v int64) *ProductUpdateCmd
	GetCompanyId() *int64
	SetDescription(v string) *ProductUpdateCmd
	GetDescription() *string
	SetId(v int64) *ProductUpdateCmd
	GetId() *int64
	SetRequestId(v string) *ProductUpdateCmd
	GetRequestId() *string
}

type ProductUpdateCmd struct {
	Alias       *string `json:"alias,omitempty" xml:"alias,omitempty"`
	CompanyId   *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RequestId   *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ProductUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s ProductUpdateCmd) GoString() string {
	return s.String()
}

func (s *ProductUpdateCmd) GetAlias() *string {
	return s.Alias
}

func (s *ProductUpdateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ProductUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *ProductUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *ProductUpdateCmd) GetRequestId() *string {
	return s.RequestId
}

func (s *ProductUpdateCmd) SetAlias(v string) *ProductUpdateCmd {
	s.Alias = &v
	return s
}

func (s *ProductUpdateCmd) SetCompanyId(v int64) *ProductUpdateCmd {
	s.CompanyId = &v
	return s
}

func (s *ProductUpdateCmd) SetDescription(v string) *ProductUpdateCmd {
	s.Description = &v
	return s
}

func (s *ProductUpdateCmd) SetId(v int64) *ProductUpdateCmd {
	s.Id = &v
	return s
}

func (s *ProductUpdateCmd) SetRequestId(v string) *ProductUpdateCmd {
	s.RequestId = &v
	return s
}

func (s *ProductUpdateCmd) Validate() error {
	return dara.Validate(s)
}
