// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteNameExclusiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetSiteNameExclusiveRequest
	GetSiteId() *int64
}

type GetSiteNameExclusiveRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteNameExclusiveRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSiteNameExclusiveRequest) GoString() string {
	return s.String()
}

func (s *GetSiteNameExclusiveRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteNameExclusiveRequest) SetSiteId(v int64) *GetSiteNameExclusiveRequest {
	s.SiteId = &v
	return s
}

func (s *GetSiteNameExclusiveRequest) Validate() error {
	return dara.Validate(s)
}
