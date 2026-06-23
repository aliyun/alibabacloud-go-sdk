// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPurgeQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetPurgeQuotaRequest
	GetSiteId() *int64
	SetType(v string) *GetPurgeQuotaRequest
	GetType() *string
}

type GetPurgeQuotaRequest struct {
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The refresh task type. Valid values:
	//
	// - **file*	- (default): file refresh.
	//
	// - **cachetag**: cache tag refresh.
	//
	// - **directory**: directory refresh.
	//
	// - **ignoreParams**: ignore-parameters refresh.
	//
	// - **hostname**: hostname refresh.
	//
	// - **purgeall**: purge all cache under the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPurgeQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPurgeQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetPurgeQuotaRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetPurgeQuotaRequest) GetType() *string {
	return s.Type
}

func (s *GetPurgeQuotaRequest) SetSiteId(v int64) *GetPurgeQuotaRequest {
	s.SiteId = &v
	return s
}

func (s *GetPurgeQuotaRequest) SetType(v string) *GetPurgeQuotaRequest {
	s.Type = &v
	return s
}

func (s *GetPurgeQuotaRequest) Validate() error {
	return dara.Validate(s)
}
