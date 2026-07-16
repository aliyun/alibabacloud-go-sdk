// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentName(v string) *DeleteEnvironmentRequest
	GetEnvironmentName() *string
	SetSiteId(v int64) *DeleteEnvironmentRequest
	GetSiteId() *int64
}

type DeleteEnvironmentRequest struct {
	// The environment name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 环境1
	EnvironmentName *string `json:"EnvironmentName,omitempty" xml:"EnvironmentName,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890**
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentRequest) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *DeleteEnvironmentRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteEnvironmentRequest) SetEnvironmentName(v string) *DeleteEnvironmentRequest {
	s.EnvironmentName = &v
	return s
}

func (s *DeleteEnvironmentRequest) SetSiteId(v int64) *DeleteEnvironmentRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
