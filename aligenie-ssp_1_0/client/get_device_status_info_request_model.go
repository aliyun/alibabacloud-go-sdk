// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceStatusInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetDeviceStatusInfoRequestDeviceInfo) *GetDeviceStatusInfoRequest
	GetDeviceInfo() *GetDeviceStatusInfoRequestDeviceInfo
}

type GetDeviceStatusInfoRequest struct {
	// This parameter is required.
	DeviceInfo *GetDeviceStatusInfoRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s GetDeviceStatusInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoRequest) GetDeviceInfo() *GetDeviceStatusInfoRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetDeviceStatusInfoRequest) SetDeviceInfo(v *GetDeviceStatusInfoRequestDeviceInfo) *GetDeviceStatusInfoRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetDeviceStatusInfoRequest) Validate() error {
	return dara.Validate(s)
}

type GetDeviceStatusInfoRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceStatusInfoRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusInfoRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) SetEncodeKey(v string) *GetDeviceStatusInfoRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) SetEncodeType(v string) *GetDeviceStatusInfoRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) SetId(v string) *GetDeviceStatusInfoRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) SetIdType(v string) *GetDeviceStatusInfoRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) SetOrganizationId(v string) *GetDeviceStatusInfoRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetDeviceStatusInfoRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}
