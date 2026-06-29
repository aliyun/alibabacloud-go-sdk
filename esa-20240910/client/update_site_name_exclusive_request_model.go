// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteNameExclusiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *UpdateSiteNameExclusiveRequest
	GetEnable() *string
	SetSiteId(v int64) *UpdateSiteNameExclusiveRequest
	GetSiteId() *int64
}

type UpdateSiteNameExclusiveRequest struct {
	// The feature switch. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSiteNameExclusiveRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteNameExclusiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteNameExclusiveRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateSiteNameExclusiveRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSiteNameExclusiveRequest) SetEnable(v string) *UpdateSiteNameExclusiveRequest {
	s.Enable = &v
	return s
}

func (s *UpdateSiteNameExclusiveRequest) SetSiteId(v int64) *UpdateSiteNameExclusiveRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteNameExclusiveRequest) Validate() error {
	return dara.Validate(s)
}
