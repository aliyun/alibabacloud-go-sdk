// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteCustomLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetSiteCustomLogRequest
	GetSiteId() *int64
}

type GetSiteCustomLogRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11223***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteCustomLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSiteCustomLogRequest) GoString() string {
	return s.String()
}

func (s *GetSiteCustomLogRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteCustomLogRequest) SetSiteId(v int64) *GetSiteCustomLogRequest {
	s.SiteId = &v
	return s
}

func (s *GetSiteCustomLogRequest) Validate() error {
	return dara.Validate(s)
}
