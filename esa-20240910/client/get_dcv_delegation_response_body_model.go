// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcvDelegationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDelegationDomain(v string) *GetDcvDelegationResponseBody
	GetDelegationDomain() *string
	SetRequestId(v string) *GetDcvDelegationResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetDcvDelegationResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetDcvDelegationResponseBody
	GetSiteName() *string
}

type GetDcvDelegationResponseBody struct {
	// The DCV managed domain name.
	//
	// example:
	//
	// 123456****.dcv.aliyun-esa.com
	DelegationDomain *string `json:"DelegationDomain,omitempty" xml:"DelegationDomain,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The site ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s GetDcvDelegationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDcvDelegationResponseBody) GoString() string {
	return s.String()
}

func (s *GetDcvDelegationResponseBody) GetDelegationDomain() *string {
	return s.DelegationDomain
}

func (s *GetDcvDelegationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDcvDelegationResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetDcvDelegationResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetDcvDelegationResponseBody) SetDelegationDomain(v string) *GetDcvDelegationResponseBody {
	s.DelegationDomain = &v
	return s
}

func (s *GetDcvDelegationResponseBody) SetRequestId(v string) *GetDcvDelegationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDcvDelegationResponseBody) SetSiteId(v int64) *GetDcvDelegationResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetDcvDelegationResponseBody) SetSiteName(v string) *GetDcvDelegationResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetDcvDelegationResponseBody) Validate() error {
	return dara.Validate(s)
}
