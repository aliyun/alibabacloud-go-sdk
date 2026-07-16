// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentName(v string) *UpdateEnvironmentVersionRequest
	GetEnvironmentName() *string
	SetSiteId(v int64) *UpdateEnvironmentVersionRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *UpdateEnvironmentVersionRequest
	GetSiteVersion() *int32
}

type UpdateEnvironmentVersionRequest struct {
	// The name of the environment to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 默认环境
	EnvironmentName *string `json:"EnvironmentName,omitempty" xml:"EnvironmentName,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 33970510651****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The new site version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s UpdateEnvironmentVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentVersionRequest) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *UpdateEnvironmentVersionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateEnvironmentVersionRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *UpdateEnvironmentVersionRequest) SetEnvironmentName(v string) *UpdateEnvironmentVersionRequest {
	s.EnvironmentName = &v
	return s
}

func (s *UpdateEnvironmentVersionRequest) SetSiteId(v int64) *UpdateEnvironmentVersionRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateEnvironmentVersionRequest) SetSiteVersion(v int32) *UpdateEnvironmentVersionRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateEnvironmentVersionRequest) Validate() error {
	return dara.Validate(s)
}
