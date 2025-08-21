// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserByDeviceIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetUserByDeviceIdRequestDeviceInfo) *GetUserByDeviceIdRequest
	GetDeviceInfo() *GetUserByDeviceIdRequestDeviceInfo
}

type GetUserByDeviceIdRequest struct {
	// This parameter is required.
	DeviceInfo *GetUserByDeviceIdRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s GetUserByDeviceIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserByDeviceIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdRequest) GetDeviceInfo() *GetUserByDeviceIdRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetUserByDeviceIdRequest) SetDeviceInfo(v *GetUserByDeviceIdRequestDeviceInfo) *GetUserByDeviceIdRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetUserByDeviceIdRequest) Validate() error {
	return dara.Validate(s)
}

type GetUserByDeviceIdRequestDeviceInfo struct {
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

func (s GetUserByDeviceIdRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetUserByDeviceIdRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetUserByDeviceIdRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetUserByDeviceIdRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetUserByDeviceIdRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetUserByDeviceIdRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetUserByDeviceIdRequestDeviceInfo) SetEncodeKey(v string) *GetUserByDeviceIdRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetUserByDeviceIdRequestDeviceInfo) SetEncodeType(v string) *GetUserByDeviceIdRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetUserByDeviceIdRequestDeviceInfo) SetId(v string) *GetUserByDeviceIdRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetUserByDeviceIdRequestDeviceInfo) SetIdType(v string) *GetUserByDeviceIdRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetUserByDeviceIdRequestDeviceInfo) SetOrganizationId(v string) *GetUserByDeviceIdRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetUserByDeviceIdRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}
