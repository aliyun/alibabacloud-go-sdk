// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoDRMLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrmType(v string) *VideoDRMLicenseRequest
	GetDrmType() *string
	SetLicenseRequest(v string) *VideoDRMLicenseRequest
	GetLicenseRequest() *string
}

type VideoDRMLicenseRequest struct {
	// The type of DRM encryption.
	//
	// Valid values:
	//
	// 	- fairplay
	//
	// 	- widevine
	//
	// This parameter is required.
	//
	// example:
	//
	// widevine
	DrmType *string `json:"drmType,omitempty" xml:"drmType,omitempty"`
	// The request that is initiated to obtain the license.
	//
	// example:
	//
	// CAES6B8SQgpACioSENGxDhqCLIVwwCBOyPayyWoSENGxDhqCLIVwwCBOyPayyWpI88aJmwYQARoQdRV32
	LicenseRequest *string `json:"licenseRequest,omitempty" xml:"licenseRequest,omitempty"`
}

func (s VideoDRMLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoDRMLicenseRequest) GoString() string {
	return s.String()
}

func (s *VideoDRMLicenseRequest) GetDrmType() *string {
	return s.DrmType
}

func (s *VideoDRMLicenseRequest) GetLicenseRequest() *string {
	return s.LicenseRequest
}

func (s *VideoDRMLicenseRequest) SetDrmType(v string) *VideoDRMLicenseRequest {
	s.DrmType = &v
	return s
}

func (s *VideoDRMLicenseRequest) SetLicenseRequest(v string) *VideoDRMLicenseRequest {
	s.LicenseRequest = &v
	return s
}

func (s *VideoDRMLicenseRequest) Validate() error {
	return dara.Validate(s)
}
