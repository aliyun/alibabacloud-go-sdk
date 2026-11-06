// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndBindAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIconUrl(v string) *CreateAndBindAppRequest
	GetAppIconUrl() *string
	SetAppName(v string) *CreateAndBindAppRequest
	GetAppName() *string
	SetImageId(v string) *CreateAndBindAppRequest
	GetImageId() *string
	SetOriginAppType(v string) *CreateAndBindAppRequest
	GetOriginAppType() *string
	SetPort(v string) *CreateAndBindAppRequest
	GetPort() *string
	SetUrlSuffix(v string) *CreateAndBindAppRequest
	GetUrlSuffix() *string
	SetVersionId(v string) *CreateAndBindAppRequest
	GetVersionId() *string
}

type CreateAndBindAppRequest struct {
	// The URL of the application icon. Specify an accessible icon URL when a custom icon needs to be displayed.
	//
	// example:
	//
	// https://example.com/apps/business/icon.png
	AppIconUrl *string `json:"AppIconUrl,omitempty" xml:"AppIconUrl,omitempty"`
	// The name of the application to be created, which is used to identify the application in the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business management platform
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the image to which the application is bound. The image must be a custom image or community image that can be managed by the current account.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-bp1example****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The application form. If omitted, the default value is CLIENT_BASE. Valid values:
	//
	// - **CLIENT_BASE**: client application
	//
	// - **WEB_BASE**: browser application
	//
	// - **CLOUD_NATIVE_BASE**: cloud-native application
	//
	// - **PORT_SERVICE_BASE**: port service application
	//
	// - **AGENT_SERVICE_BASE**: agent application
	//
	// example:
	//
	// PORT_SERVICE_BASE
	OriginAppType *string `json:"OriginAppType,omitempty" xml:"OriginAppType,omitempty"`
	// The service port of the application. For applications that require port-based access, specify the actual port number that the application listens on.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The path suffix of the application access URL, which describes the access entry point together with the application service port.
	//
	// example:
	//
	// /console
	UrlSuffix *string `json:"UrlSuffix,omitempty" xml:"UrlSuffix,omitempty"`
	// The ID of the image version to which the application is bound. If omitted, the current active version of the image specified by ImageId is used. If specified, the version must belong to the image.
	//
	// example:
	//
	// iv-bp1example****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateAndBindAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndBindAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAndBindAppRequest) GetAppIconUrl() *string {
	return s.AppIconUrl
}

func (s *CreateAndBindAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAndBindAppRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateAndBindAppRequest) GetOriginAppType() *string {
	return s.OriginAppType
}

func (s *CreateAndBindAppRequest) GetPort() *string {
	return s.Port
}

func (s *CreateAndBindAppRequest) GetUrlSuffix() *string {
	return s.UrlSuffix
}

func (s *CreateAndBindAppRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateAndBindAppRequest) SetAppIconUrl(v string) *CreateAndBindAppRequest {
	s.AppIconUrl = &v
	return s
}

func (s *CreateAndBindAppRequest) SetAppName(v string) *CreateAndBindAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAndBindAppRequest) SetImageId(v string) *CreateAndBindAppRequest {
	s.ImageId = &v
	return s
}

func (s *CreateAndBindAppRequest) SetOriginAppType(v string) *CreateAndBindAppRequest {
	s.OriginAppType = &v
	return s
}

func (s *CreateAndBindAppRequest) SetPort(v string) *CreateAndBindAppRequest {
	s.Port = &v
	return s
}

func (s *CreateAndBindAppRequest) SetUrlSuffix(v string) *CreateAndBindAppRequest {
	s.UrlSuffix = &v
	return s
}

func (s *CreateAndBindAppRequest) SetVersionId(v string) *CreateAndBindAppRequest {
	s.VersionId = &v
	return s
}

func (s *CreateAndBindAppRequest) Validate() error {
	return dara.Validate(s)
}
