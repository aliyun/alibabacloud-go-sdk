// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuota(v int64) *GetCertificateQuotaResponseBody
	GetQuota() *int64
	SetQuotaUsage(v int64) *GetCertificateQuotaResponseBody
	GetQuotaUsage() *int64
	SetRequestId(v string) *GetCertificateQuotaResponseBody
	GetRequestId() *string
	SetSiteCount(v int64) *GetCertificateQuotaResponseBody
	GetSiteCount() *int64
	SetSiteUsage(v []*GetCertificateQuotaResponseBodySiteUsage) *GetCertificateQuotaResponseBody
	GetSiteUsage() []*GetCertificateQuotaResponseBodySiteUsage
	SetType(v string) *GetCertificateQuotaResponseBody
	GetType() *string
}

type GetCertificateQuotaResponseBody struct {
	// Free certificate quota.
	//
	// example:
	//
	// 10
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// Usage of free certificate quota.
	//
	// example:
	//
	// 5
	QuotaUsage *int64 `json:"QuotaUsage,omitempty" xml:"QuotaUsage,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Number of sites.
	//
	// example:
	//
	// 2
	SiteCount *int64 `json:"SiteCount,omitempty" xml:"SiteCount,omitempty"`
	// List of site usage details.
	SiteUsage []*GetCertificateQuotaResponseBodySiteUsage `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty" type:"Repeated"`
	// Certificate Quota type.
	//
	// example:
	//
	// free
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCertificateQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertificateQuotaResponseBody) GetQuota() *int64 {
	return s.Quota
}

func (s *GetCertificateQuotaResponseBody) GetQuotaUsage() *int64 {
	return s.QuotaUsage
}

func (s *GetCertificateQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCertificateQuotaResponseBody) GetSiteCount() *int64 {
	return s.SiteCount
}

func (s *GetCertificateQuotaResponseBody) GetSiteUsage() []*GetCertificateQuotaResponseBodySiteUsage {
	return s.SiteUsage
}

func (s *GetCertificateQuotaResponseBody) GetType() *string {
	return s.Type
}

func (s *GetCertificateQuotaResponseBody) SetQuota(v int64) *GetCertificateQuotaResponseBody {
	s.Quota = &v
	return s
}

func (s *GetCertificateQuotaResponseBody) SetQuotaUsage(v int64) *GetCertificateQuotaResponseBody {
	s.QuotaUsage = &v
	return s
}

func (s *GetCertificateQuotaResponseBody) SetRequestId(v string) *GetCertificateQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertificateQuotaResponseBody) SetSiteCount(v int64) *GetCertificateQuotaResponseBody {
	s.SiteCount = &v
	return s
}

func (s *GetCertificateQuotaResponseBody) SetSiteUsage(v []*GetCertificateQuotaResponseBodySiteUsage) *GetCertificateQuotaResponseBody {
	s.SiteUsage = v
	return s
}

func (s *GetCertificateQuotaResponseBody) SetType(v string) *GetCertificateQuotaResponseBody {
	s.Type = &v
	return s
}

func (s *GetCertificateQuotaResponseBody) Validate() error {
	if s.SiteUsage != nil {
		for _, item := range s.SiteUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCertificateQuotaResponseBodySiteUsage struct {
	// Site ID.
	//
	// example:
	//
	// 165929521496928
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Site name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// Site usage.
	//
	// example:
	//
	// 5
	SiteUsage *int64 `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty"`
}

func (s GetCertificateQuotaResponseBodySiteUsage) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateQuotaResponseBodySiteUsage) GoString() string {
	return s.String()
}

func (s *GetCertificateQuotaResponseBodySiteUsage) GetSiteId() *string {
	return s.SiteId
}

func (s *GetCertificateQuotaResponseBodySiteUsage) GetSiteName() *string {
	return s.SiteName
}

func (s *GetCertificateQuotaResponseBodySiteUsage) GetSiteUsage() *int64 {
	return s.SiteUsage
}

func (s *GetCertificateQuotaResponseBodySiteUsage) SetSiteId(v string) *GetCertificateQuotaResponseBodySiteUsage {
	s.SiteId = &v
	return s
}

func (s *GetCertificateQuotaResponseBodySiteUsage) SetSiteName(v string) *GetCertificateQuotaResponseBodySiteUsage {
	s.SiteName = &v
	return s
}

func (s *GetCertificateQuotaResponseBodySiteUsage) SetSiteUsage(v int64) *GetCertificateQuotaResponseBodySiteUsage {
	s.SiteUsage = &v
	return s
}

func (s *GetCertificateQuotaResponseBodySiteUsage) Validate() error {
	return dara.Validate(s)
}
