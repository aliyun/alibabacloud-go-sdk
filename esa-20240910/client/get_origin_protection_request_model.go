// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetOriginProtectionRequest
	GetSiteId() *int64
}

type GetOriginProtectionRequest struct {
	// The website ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetOriginProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOriginProtectionRequest) GoString() string {
	return s.String()
}

func (s *GetOriginProtectionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginProtectionRequest) SetSiteId(v int64) *GetOriginProtectionRequest {
	s.SiteId = &v
	return s
}

func (s *GetOriginProtectionRequest) Validate() error {
	return dara.Validate(s)
}
