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
	// The website ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The type of the purge task. Valid values:
	//
	// 	- **file*	- (default): purges the cache by file.
	//
	// 	- **cachetag**: purges the cache by cache tag.
	//
	// 	- **directory**: purges the cache by directory.
	//
	// 	- **ignoreParams**: purges the cache by URL with specific parameters ignored.
	//
	// 	- **hostname**: purges the cache by hostname.
	//
	// 	- **purgeall**: purges all cache.
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
