// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSeoBypassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *UpdateSeoBypassRequest
	GetEnable() *string
	SetSiteId(v int64) *UpdateSeoBypassRequest
	GetSiteId() *int64
}

type UpdateSeoBypassRequest struct {
	// Specifies whether to enable the feature. Valid values:
	//
	// 	- **on:**
	//
	// 	- **off**
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSeoBypassRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSeoBypassRequest) GoString() string {
	return s.String()
}

func (s *UpdateSeoBypassRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateSeoBypassRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSeoBypassRequest) SetEnable(v string) *UpdateSeoBypassRequest {
	s.Enable = &v
	return s
}

func (s *UpdateSeoBypassRequest) SetSiteId(v int64) *UpdateSeoBypassRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSeoBypassRequest) Validate() error {
	return dara.Validate(s)
}
