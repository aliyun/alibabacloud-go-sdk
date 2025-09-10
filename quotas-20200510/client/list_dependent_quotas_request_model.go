// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDependentQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductCode(v string) *ListDependentQuotasRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *ListDependentQuotasRequest
	GetQuotaActionCode() *string
}

type ListDependentQuotasRequest struct {
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// csk
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// q_i5uzm3
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
}

func (s ListDependentQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDependentQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDependentQuotasRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListDependentQuotasRequest) SetProductCode(v string) *ListDependentQuotasRequest {
	s.ProductCode = &v
	return s
}

func (s *ListDependentQuotasRequest) SetQuotaActionCode(v string) *ListDependentQuotasRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListDependentQuotasRequest) Validate() error {
	return dara.Validate(s)
}
