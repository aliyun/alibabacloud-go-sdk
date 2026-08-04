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
	// List of user identifier information.
	//
	// This parameter is required.
	DeviceInfo *SetDeviceSettingRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// The collection of keys specified for device settings:
	//
	// Do Not Disturb mode: nightMode
	//
	// This parameter is required.
	//
	// example:
	//
	// nightMode
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Attribute Value
	//
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
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetDeviceSettingRequestDeviceInfo struct {
	// Value corresponding to the encoding type. Enter the Project ID of the project where the product resides. You can view it in the Tmall Genie AI platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. Enter PROJECT_ID here.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID. Enter the value of deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of Device ID:
	//
	// OPEN_ID: The default device identity.
	//
	// UNION_ID: The organization-dimension device identity. You must request an organization in advance on the Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required if IdType is UNION_ID.
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
