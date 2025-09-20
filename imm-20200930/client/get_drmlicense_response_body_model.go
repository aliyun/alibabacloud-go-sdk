// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDRMLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v string) *GetDRMLicenseResponseBody
	GetDeviceInfo() *string
	SetLicense(v string) *GetDRMLicenseResponseBody
	GetLicense() *string
	SetRequestId(v string) *GetDRMLicenseResponseBody
	GetRequestId() *string
	SetStates(v int64) *GetDRMLicenseResponseBody
	GetStates() *int64
}

type GetDRMLicenseResponseBody struct {
	// example:
	//
	// IEEE1284DeviceID
	DeviceInfo *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// example:
	//
	// AESzB8SQgpACioSEJ3yqiFwruAOUgIvlCx*****
	License *string `json:"License,omitempty" xml:"License,omitempty"`
	// example:
	//
	// 896ABAD1-C452-4BED-B5E0-302955F*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	States *int64 `json:"States,omitempty" xml:"States,omitempty"`
}

func (s GetDRMLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDRMLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseResponseBody) GetDeviceInfo() *string {
	return s.DeviceInfo
}

func (s *GetDRMLicenseResponseBody) GetLicense() *string {
	return s.License
}

func (s *GetDRMLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDRMLicenseResponseBody) GetStates() *int64 {
	return s.States
}

func (s *GetDRMLicenseResponseBody) SetDeviceInfo(v string) *GetDRMLicenseResponseBody {
	s.DeviceInfo = &v
	return s
}

func (s *GetDRMLicenseResponseBody) SetLicense(v string) *GetDRMLicenseResponseBody {
	s.License = &v
	return s
}

func (s *GetDRMLicenseResponseBody) SetRequestId(v string) *GetDRMLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDRMLicenseResponseBody) SetStates(v int64) *GetDRMLicenseResponseBody {
	s.States = &v
	return s
}

func (s *GetDRMLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
