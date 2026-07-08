// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertWarehouseQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppTotalQuota(v int64) *GetCertWarehouseQuotaResponseBody
	GetAppTotalQuota() *int64
	SetAppUseCount(v int64) *GetCertWarehouseQuotaResponseBody
	GetAppUseCount() *int64
	SetRequestId(v string) *GetCertWarehouseQuotaResponseBody
	GetRequestId() *string
	SetTotalQuota(v int64) *GetCertWarehouseQuotaResponseBody
	GetTotalQuota() *int64
	SetUseCount(v int64) *GetCertWarehouseQuotaResponseBody
	GetUseCount() *int64
}

type GetCertWarehouseQuotaResponseBody struct {
	// The total quota for the certificate application service. This includes both complimentary and purchased quotas.
	//
	// example:
	//
	// 5000
	AppTotalQuota *int64 `json:"AppTotalQuota,omitempty" xml:"AppTotalQuota,omitempty"`
	// The used quota for the certificate application service.
	//
	// example:
	//
	// 1000
	AppUseCount *int64 `json:"AppUseCount,omitempty" xml:"AppUseCount,omitempty"`
	// The request ID. Alibaba Cloud generates this unique ID for each request. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total quota for the electronic contract signing service. This includes both complimentary and purchased quotas.
	//
	// example:
	//
	// 5000
	TotalQuota *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The used quota for the electronic contract signing service.
	//
	// example:
	//
	// 1000
	UseCount *int64 `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
}

func (s GetCertWarehouseQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCertWarehouseQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertWarehouseQuotaResponseBody) GetAppTotalQuota() *int64 {
	return s.AppTotalQuota
}

func (s *GetCertWarehouseQuotaResponseBody) GetAppUseCount() *int64 {
	return s.AppUseCount
}

func (s *GetCertWarehouseQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCertWarehouseQuotaResponseBody) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *GetCertWarehouseQuotaResponseBody) GetUseCount() *int64 {
	return s.UseCount
}

func (s *GetCertWarehouseQuotaResponseBody) SetAppTotalQuota(v int64) *GetCertWarehouseQuotaResponseBody {
	s.AppTotalQuota = &v
	return s
}

func (s *GetCertWarehouseQuotaResponseBody) SetAppUseCount(v int64) *GetCertWarehouseQuotaResponseBody {
	s.AppUseCount = &v
	return s
}

func (s *GetCertWarehouseQuotaResponseBody) SetRequestId(v string) *GetCertWarehouseQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertWarehouseQuotaResponseBody) SetTotalQuota(v int64) *GetCertWarehouseQuotaResponseBody {
	s.TotalQuota = &v
	return s
}

func (s *GetCertWarehouseQuotaResponseBody) SetUseCount(v int64) *GetCertWarehouseQuotaResponseBody {
	s.UseCount = &v
	return s
}

func (s *GetCertWarehouseQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
