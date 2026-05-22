// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteWafSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *GetSiteWafSettingsRequest
	GetPath() *string
	SetSiteId(v int64) *GetSiteWafSettingsRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *GetSiteWafSettingsRequest
	GetSiteVersion() *int32
}

type GetSiteWafSettingsRequest struct {
	// Specifies the configuration path. If not provided, all configurations will be retrieved.
	//
	// example:
	//
	// bot_management
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Site version.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetSiteWafSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSiteWafSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetSiteWafSettingsRequest) GetPath() *string {
	return s.Path
}

func (s *GetSiteWafSettingsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteWafSettingsRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetSiteWafSettingsRequest) SetPath(v string) *GetSiteWafSettingsRequest {
	s.Path = &v
	return s
}

func (s *GetSiteWafSettingsRequest) SetSiteId(v int64) *GetSiteWafSettingsRequest {
	s.SiteId = &v
	return s
}

func (s *GetSiteWafSettingsRequest) SetSiteVersion(v int32) *GetSiteWafSettingsRequest {
	s.SiteVersion = &v
	return s
}

func (s *GetSiteWafSettingsRequest) Validate() error {
	return dara.Validate(s)
}
