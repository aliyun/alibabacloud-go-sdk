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
	// Specifies whether to include the header that indicates the geographical location of a client in an origin request. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	AddClientGeolocationHeader *string `json:"AddClientGeolocationHeader,omitempty" xml:"AddClientGeolocationHeader,omitempty"`
	// Specifies whether to include the "ali-real-client-ip" header that indicates the client\\"s real IP address in an origin request. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	AddRealClientIpHeader  *string `json:"AddRealClientIpHeader,omitempty" xml:"AddRealClientIpHeader,omitempty"`
	RealClientIpHeaderName *string `json:"RealClientIpHeaderName,omitempty" xml:"RealClientIpHeaderName,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the website. You can use this parameter to specify a version of your website to apply the feature settings. By default, version 0 is used.
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
