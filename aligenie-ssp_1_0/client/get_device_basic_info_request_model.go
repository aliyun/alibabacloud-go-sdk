// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetDeviceBasicInfoRequestDeviceInfo) *GetDeviceBasicInfoRequest
	GetDeviceInfo() *GetDeviceBasicInfoRequestDeviceInfo
}

type GetDeviceBasicInfoRequest struct {
	// This parameter is required.
	DeviceInfo *GetDeviceBasicInfoRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
}

func (s GetDeviceBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoRequest) GetDeviceInfo() *GetDeviceBasicInfoRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetDeviceBasicInfoRequest) SetDeviceInfo(v *GetDeviceBasicInfoRequestDeviceInfo) *GetDeviceBasicInfoRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetDeviceBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}

type GetDeviceBasicInfoRequestDeviceInfo struct {
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

func (s GetDeviceBasicInfoRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceBasicInfoRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) SetEncodeKey(v string) *GetDeviceBasicInfoRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) SetEncodeType(v string) *GetDeviceBasicInfoRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) SetId(v string) *GetDeviceBasicInfoRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) SetIdType(v string) *GetDeviceBasicInfoRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) SetOrganizationId(v string) *GetDeviceBasicInfoRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetDeviceBasicInfoRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}
