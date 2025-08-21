// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *SetDeviceSettingRequestDeviceInfo) *SetDeviceSettingRequest
	GetDeviceInfo() *SetDeviceSettingRequestDeviceInfo
	SetKey(v string) *SetDeviceSettingRequest
	GetKey() *string
	SetValue(v interface{}) *SetDeviceSettingRequest
	GetValue() interface{}
}

type SetDeviceSettingRequest struct {
	// This parameter is required.
	DeviceInfo *SetDeviceSettingRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// nightMode
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// {"enable":true}
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetDeviceSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceSettingRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingRequest) GetDeviceInfo() *SetDeviceSettingRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *SetDeviceSettingRequest) GetKey() *string {
	return s.Key
}

func (s *SetDeviceSettingRequest) GetValue() interface{} {
	return s.Value
}

func (s *SetDeviceSettingRequest) SetDeviceInfo(v *SetDeviceSettingRequestDeviceInfo) *SetDeviceSettingRequest {
	s.DeviceInfo = v
	return s
}

func (s *SetDeviceSettingRequest) SetKey(v string) *SetDeviceSettingRequest {
	s.Key = &v
	return s
}

func (s *SetDeviceSettingRequest) SetValue(v interface{}) *SetDeviceSettingRequest {
	s.Value = v
	return s
}

func (s *SetDeviceSettingRequest) Validate() error {
	return dara.Validate(s)
}

type SetDeviceSettingRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
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
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SetDeviceSettingRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceSettingRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *SetDeviceSettingRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *SetDeviceSettingRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *SetDeviceSettingRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *SetDeviceSettingRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *SetDeviceSettingRequestDeviceInfo) SetEncodeKey(v string) *SetDeviceSettingRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *SetDeviceSettingRequestDeviceInfo) SetEncodeType(v string) *SetDeviceSettingRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *SetDeviceSettingRequestDeviceInfo) SetId(v string) *SetDeviceSettingRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *SetDeviceSettingRequestDeviceInfo) SetIdType(v string) *SetDeviceSettingRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *SetDeviceSettingRequestDeviceInfo) SetOrganizationId(v string) *SetDeviceSettingRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *SetDeviceSettingRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}
