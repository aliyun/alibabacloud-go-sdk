// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedTransformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetManagedTransformRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *GetManagedTransformRequest
	GetSiteVersion() *int32
}

type GetManagedTransformRequest struct {
	// The ID of the site. You can obtain this ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site version. If version management is enabled for the site, use this parameter to specify the version whose configuration you want to retrieve. The default is 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetManagedTransformRequest) String() string {
	return dara.Prettify(s)
}

func (s GetManagedTransformRequest) GoString() string {
	return s.String()
}

func (s *GetManagedTransformRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetManagedTransformRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetManagedTransformRequest) SetSiteId(v int64) *GetManagedTransformRequest {
	s.SiteId = &v
	return s
}

func (s *GetManagedTransformRequest) SetSiteVersion(v int32) *GetManagedTransformRequest {
	s.SiteVersion = &v
	return s
}

func (s *GetManagedTransformRequest) Validate() error {
	return dara.Validate(s)
}
