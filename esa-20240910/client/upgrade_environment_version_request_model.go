// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeEnvironmentVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentName(v string) *UpgradeEnvironmentVersionRequest
	GetEnvironmentName() *string
	SetSiteId(v int64) *UpgradeEnvironmentVersionRequest
	GetSiteId() *int64
}

type UpgradeEnvironmentVersionRequest struct {
	// The environment name. The version of this environment is upgraded and deployed to the environment with the next priority level.
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
	// 1245678****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpgradeEnvironmentVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEnvironmentVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeEnvironmentVersionRequest) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *UpgradeEnvironmentVersionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpgradeEnvironmentVersionRequest) SetEnvironmentName(v string) *UpgradeEnvironmentVersionRequest {
	s.EnvironmentName = &v
	return s
}

func (s *UpgradeEnvironmentVersionRequest) SetSiteId(v int64) *UpgradeEnvironmentVersionRequest {
	s.SiteId = &v
	return s
}

func (s *UpgradeEnvironmentVersionRequest) Validate() error {
	return dara.Validate(s)
}
