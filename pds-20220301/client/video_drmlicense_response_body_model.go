// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoDRMLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *VideoDRMLicenseResponseBody
	GetData() *string
	SetDeviceInfo(v string) *VideoDRMLicenseResponseBody
	GetDeviceInfo() *string
	SetStates(v string) *VideoDRMLicenseResponseBody
	GetStates() *string
}

type VideoDRMLicenseResponseBody struct {
	// The returned DRM license.
	//
	// example:
	//
	// cb9swCy8P50H9KePsxET3jZ1tm41bDs9HTsxbWnsjf3bsf6QGdiS4kZPhDaskimbNyAfNjmhQRmWFt3AhwNF3
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The information about the device from which the DRM request was initiated.
	//
	// example:
	//
	// ""
	DeviceInfo *string `json:"device_info,omitempty" xml:"device_info,omitempty"`
	// The request state returned by the DRM server.
	//
	// example:
	//
	// 0
	States *string `json:"states,omitempty" xml:"states,omitempty"`
}

func (s VideoDRMLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VideoDRMLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *VideoDRMLicenseResponseBody) GetData() *string {
	return s.Data
}

func (s *VideoDRMLicenseResponseBody) GetDeviceInfo() *string {
	return s.DeviceInfo
}

func (s *VideoDRMLicenseResponseBody) GetStates() *string {
	return s.States
}

func (s *VideoDRMLicenseResponseBody) SetData(v string) *VideoDRMLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *VideoDRMLicenseResponseBody) SetDeviceInfo(v string) *VideoDRMLicenseResponseBody {
	s.DeviceInfo = &v
	return s
}

func (s *VideoDRMLicenseResponseBody) SetStates(v string) *VideoDRMLicenseResponseBody {
	s.States = &v
	return s
}

func (s *VideoDRMLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
