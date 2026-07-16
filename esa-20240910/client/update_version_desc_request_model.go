// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVersionDescRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateVersionDescRequest
	GetDescription() *string
	SetSiteId(v int64) *UpdateVersionDescRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *UpdateVersionDescRequest
	GetSiteVersion() *int32
}

type UpdateVersionDescRequest struct {
	// The description of the version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 更新版本。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 134567****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number to be updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s UpdateVersionDescRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVersionDescRequest) GoString() string {
	return s.String()
}

func (s *UpdateVersionDescRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateVersionDescRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateVersionDescRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *UpdateVersionDescRequest) SetDescription(v string) *UpdateVersionDescRequest {
	s.Description = &v
	return s
}

func (s *UpdateVersionDescRequest) SetSiteId(v int64) *UpdateVersionDescRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateVersionDescRequest) SetSiteVersion(v int32) *UpdateVersionDescRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateVersionDescRequest) Validate() error {
	return dara.Validate(s)
}
