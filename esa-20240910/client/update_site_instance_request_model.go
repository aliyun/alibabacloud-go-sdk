// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewInstanceId(v string) *UpdateSiteInstanceRequest
	GetNewInstanceId() *string
	SetResourceOwner(v int64) *UpdateSiteInstanceRequest
	GetResourceOwner() *int64
	SetSiteId(v int64) *UpdateSiteInstanceRequest
	GetSiteId() *int64
}

type UpdateSiteInstanceRequest struct {
	// The target instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-site-bckh96ri1eyo
	NewInstanceId *string `json:"NewInstanceId,omitempty" xml:"NewInstanceId,omitempty"`
	ResourceOwner *int64  `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The site ID. You can call [ListSites](https://help.aliyun.com/document_detail/2850189.html) to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 901109460617712
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSiteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteInstanceRequest) GetNewInstanceId() *string {
	return s.NewInstanceId
}

func (s *UpdateSiteInstanceRequest) GetResourceOwner() *int64 {
	return s.ResourceOwner
}

func (s *UpdateSiteInstanceRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSiteInstanceRequest) SetNewInstanceId(v string) *UpdateSiteInstanceRequest {
	s.NewInstanceId = &v
	return s
}

func (s *UpdateSiteInstanceRequest) SetResourceOwner(v int64) *UpdateSiteInstanceRequest {
	s.ResourceOwner = &v
	return s
}

func (s *UpdateSiteInstanceRequest) SetSiteId(v int64) *UpdateSiteInstanceRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
