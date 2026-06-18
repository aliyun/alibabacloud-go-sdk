// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateManagedTransformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddClientGeolocationHeader(v string) *UpdateManagedTransformRequest
	GetAddClientGeolocationHeader() *string
	SetAddRealClientIpHeader(v string) *UpdateManagedTransformRequest
	GetAddRealClientIpHeader() *string
	SetRealClientIpHeaderName(v string) *UpdateManagedTransformRequest
	GetRealClientIpHeaderName() *string
	SetSiteId(v int64) *UpdateManagedTransformRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *UpdateManagedTransformRequest
	GetSiteVersion() *int32
}

type UpdateManagedTransformRequest struct {
	// Specifies whether to add a header that contains visitor geolocation information. Valid values:
	//
	// - `on`: Add the header.
	//
	// - `off`: Do not add the header.
	//
	// example:
	//
	// on
	AddClientGeolocationHeader *string `json:"AddClientGeolocationHeader,omitempty" xml:"AddClientGeolocationHeader,omitempty"`
	// Specifies whether to add the `ali-real-client-ip` header, which contains the real client IP. Valid values:
	//
	// - `on`: Add the header.
	//
	// - `off`: Do not add the header.
	//
	// example:
	//
	// on
	AddRealClientIpHeader *string `json:"AddRealClientIpHeader,omitempty" xml:"AddRealClientIpHeader,omitempty"`
	// The name of the header that contains the real client IP. The name must start with a letter and can contain letters, digits (0-9), and hyphens (-).
	//
	// example:
	//
	// test-header
	RealClientIpHeaderName *string `json:"RealClientIpHeaderName,omitempty" xml:"RealClientIpHeaderName,omitempty"`
	// The ID of the site. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version of the site. For sites with version management enabled, this parameter specifies the version to which the configuration applies. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s UpdateManagedTransformRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedTransformRequest) GoString() string {
	return s.String()
}

func (s *UpdateManagedTransformRequest) GetAddClientGeolocationHeader() *string {
	return s.AddClientGeolocationHeader
}

func (s *UpdateManagedTransformRequest) GetAddRealClientIpHeader() *string {
	return s.AddRealClientIpHeader
}

func (s *UpdateManagedTransformRequest) GetRealClientIpHeaderName() *string {
	return s.RealClientIpHeaderName
}

func (s *UpdateManagedTransformRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateManagedTransformRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *UpdateManagedTransformRequest) SetAddClientGeolocationHeader(v string) *UpdateManagedTransformRequest {
	s.AddClientGeolocationHeader = &v
	return s
}

func (s *UpdateManagedTransformRequest) SetAddRealClientIpHeader(v string) *UpdateManagedTransformRequest {
	s.AddRealClientIpHeader = &v
	return s
}

func (s *UpdateManagedTransformRequest) SetRealClientIpHeaderName(v string) *UpdateManagedTransformRequest {
	s.RealClientIpHeaderName = &v
	return s
}

func (s *UpdateManagedTransformRequest) SetSiteId(v int64) *UpdateManagedTransformRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateManagedTransformRequest) SetSiteVersion(v int32) *UpdateManagedTransformRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateManagedTransformRequest) Validate() error {
	return dara.Validate(s)
}
