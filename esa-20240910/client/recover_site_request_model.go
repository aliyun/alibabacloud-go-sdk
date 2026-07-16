// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *RecoverSiteRequest
	GetSiteId() *int64
}

type RecoverSiteRequest struct {
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s RecoverSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverSiteRequest) GoString() string {
	return s.String()
}

func (s *RecoverSiteRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *RecoverSiteRequest) SetSiteId(v int64) *RecoverSiteRequest {
	s.SiteId = &v
	return s
}

func (s *RecoverSiteRequest) Validate() error {
	return dara.Validate(s)
}
