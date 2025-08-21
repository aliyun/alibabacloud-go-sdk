// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetDeviceTagRequestDeviceInfo) *GetDeviceTagRequest
	GetDeviceInfo() *GetDeviceTagRequestDeviceInfo
}

type GetDeviceTagRequest struct {
	// This parameter is required.
	DeviceInfo *GetDeviceTagRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s GetDeviceTagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceTagRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceTagRequest) GetDeviceInfo() *GetDeviceTagRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetDeviceTagRequest) SetDeviceInfo(v *GetDeviceTagRequestDeviceInfo) *GetDeviceTagRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetDeviceTagRequest) Validate() error {
	return dara.Validate(s)
}

type GetDeviceTagRequestDeviceInfo struct {
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
	// 1
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceTagRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceTagRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceTagRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetDeviceTagRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetDeviceTagRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetDeviceTagRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetDeviceTagRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetDeviceTagRequestDeviceInfo) SetEncodeKey(v string) *GetDeviceTagRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceTagRequestDeviceInfo) SetEncodeType(v string) *GetDeviceTagRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceTagRequestDeviceInfo) SetId(v string) *GetDeviceTagRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetDeviceTagRequestDeviceInfo) SetIdType(v string) *GetDeviceTagRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetDeviceTagRequestDeviceInfo) SetOrganizationId(v string) *GetDeviceTagRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetDeviceTagRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}
