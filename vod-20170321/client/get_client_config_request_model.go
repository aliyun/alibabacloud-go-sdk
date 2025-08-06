// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrand(v string) *GetClientConfigRequest
	GetBrand() *string
	SetDeviceName(v string) *GetClientConfigRequest
	GetDeviceName() *string
	SetOsName(v string) *GetClientConfigRequest
	GetOsName() *string
}

type GetClientConfigRequest struct {
	Brand      *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	OsName     *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
}

func (s GetClientConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientConfigRequest) GoString() string {
	return s.String()
}

func (s *GetClientConfigRequest) GetBrand() *string {
	return s.Brand
}

func (s *GetClientConfigRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *GetClientConfigRequest) GetOsName() *string {
	return s.OsName
}

func (s *GetClientConfigRequest) SetBrand(v string) *GetClientConfigRequest {
	s.Brand = &v
	return s
}

func (s *GetClientConfigRequest) SetDeviceName(v string) *GetClientConfigRequest {
	s.DeviceName = &v
	return s
}

func (s *GetClientConfigRequest) SetOsName(v string) *GetClientConfigRequest {
	s.OsName = &v
	return s
}

func (s *GetClientConfigRequest) Validate() error {
	return dara.Validate(s)
}
