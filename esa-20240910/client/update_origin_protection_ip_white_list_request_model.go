// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginProtectionIpWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *UpdateOriginProtectionIpWhiteListRequest
	GetSiteId() *int64
}

type UpdateOriginProtectionIpWhiteListRequest struct {
	// Site ID. Call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to get this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateOriginProtectionIpWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginProtectionIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateOriginProtectionIpWhiteListRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateOriginProtectionIpWhiteListRequest) SetSiteId(v int64) *UpdateOriginProtectionIpWhiteListRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateOriginProtectionIpWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
