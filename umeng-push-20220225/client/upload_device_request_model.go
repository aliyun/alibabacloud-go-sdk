// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceTokens(v string) *UploadDeviceRequest
	GetDeviceTokens() *string
}

type UploadDeviceRequest struct {
	// example:
	//
	// device_token_1\\ndevice_token_2\\ndevice_token_3\\n...
	//
	// alias1\\nalias2\\nalias3\\n...
	DeviceTokens *string `json:"DeviceTokens,omitempty" xml:"DeviceTokens,omitempty"`
}

func (s UploadDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDeviceRequest) GoString() string {
	return s.String()
}

func (s *UploadDeviceRequest) GetDeviceTokens() *string {
	return s.DeviceTokens
}

func (s *UploadDeviceRequest) SetDeviceTokens(v string) *UploadDeviceRequest {
	s.DeviceTokens = &v
	return s
}

func (s *UploadDeviceRequest) Validate() error {
	return dara.Validate(s)
}
