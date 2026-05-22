// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteCurrentNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetSiteCurrentNSRequest
	GetSiteId() *int64
}

type GetSiteCurrentNSRequest struct {
	// The website ID. It can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteCurrentNSRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSiteCurrentNSRequest) GoString() string {
	return s.String()
}

func (s *GetSiteCurrentNSRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteCurrentNSRequest) SetSiteId(v int64) *GetSiteCurrentNSRequest {
	s.SiteId = &v
	return s
}

func (s *GetSiteCurrentNSRequest) Validate() error {
	return dara.Validate(s)
}
