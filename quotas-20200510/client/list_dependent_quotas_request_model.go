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
	// The abbreviated name of the cloud service.
	//
	// > For more information about how to obtain the abbreviated name of a cloud service, see `ProductCode` in [ListProducts](https://help.aliyun.com/document_detail/440555.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cloudsso
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// > For more information about how to obtain the quota ID of a target cloud service, see `QuotaActionCode` in [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// q_system-policy-per-access-configuration
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
