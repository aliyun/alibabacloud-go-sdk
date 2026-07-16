// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackEnvironmentVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentName(v string) *RollbackEnvironmentVersionRequest
	GetEnvironmentName() *string
	SetSiteId(v int64) *RollbackEnvironmentVersionRequest
	GetSiteId() *int64
}

type RollbackEnvironmentVersionRequest struct {
	// The environment name.
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
	// 33862229675****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s RollbackEnvironmentVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackEnvironmentVersionRequest) GoString() string {
	return s.String()
}

func (s *RollbackEnvironmentVersionRequest) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *RollbackEnvironmentVersionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *RollbackEnvironmentVersionRequest) SetEnvironmentName(v string) *RollbackEnvironmentVersionRequest {
	s.EnvironmentName = &v
	return s
}

func (s *RollbackEnvironmentVersionRequest) SetSiteId(v int64) *RollbackEnvironmentVersionRequest {
	s.SiteId = &v
	return s
}

func (s *RollbackEnvironmentVersionRequest) Validate() error {
	return dara.Validate(s)
}
