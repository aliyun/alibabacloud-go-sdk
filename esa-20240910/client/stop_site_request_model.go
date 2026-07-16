// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *StopSiteRequest
	GetSiteId() *int64
}

type StopSiteRequest struct {
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34003500310****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s StopSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s StopSiteRequest) GoString() string {
	return s.String()
}

func (s *StopSiteRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *StopSiteRequest) SetSiteId(v int64) *StopSiteRequest {
	s.SiteId = &v
	return s
}

func (s *StopSiteRequest) Validate() error {
	return dara.Validate(s)
}
