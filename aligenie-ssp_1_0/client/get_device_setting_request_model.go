// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetDeviceSettingRequestDeviceInfo) *GetDeviceSettingRequest
	GetDeviceInfo() *GetDeviceSettingRequestDeviceInfo
	SetKeys(v []*string) *GetDeviceSettingRequest
	GetKeys() []*string
}

type GetDeviceSettingRequest struct {
	// List of device identity information.
	DeviceInfo *GetDeviceSettingRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Set of specified keys for device Settings,
	//
	// Do Not Disturb mode: nightMode
	//
	// This parameter is required.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s GetDeviceSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceSettingRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingRequest) GetDeviceInfo() *GetDeviceSettingRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetDeviceSettingRequest) GetKeys() []*string {
	return s.Keys
}

func (s *GetDeviceSettingRequest) SetDeviceInfo(v *GetDeviceSettingRequestDeviceInfo) *GetDeviceSettingRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetDeviceSettingRequest) SetKeys(v []*string) *GetDeviceSettingRequest {
	s.Keys = v
	return s
}

func (s *GetDeviceSettingRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeviceSettingRequestDeviceInfo struct {
	// The value corresponding to the encoding type. Enter the Project ID of the project that contains the ProductKey of this product in the Tmall Genie AI Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. Set this parameter to PROJECT_ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID. Configure this parameter as deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of Device ID.
	//
	// OPEN_ID: The default device identity.
	//
	// UNION_ID: The device identity at the organization dimension. You must request an organization in advance on the Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required if IdType is UNION_ID.
	//
	// example:
	//
	// 122
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceSettingRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceSettingRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetDeviceSettingRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetDeviceSettingRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetDeviceSettingRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetDeviceSettingRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetDeviceSettingRequestDeviceInfo) SetEncodeKey(v string) *GetDeviceSettingRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceSettingRequestDeviceInfo) SetEncodeType(v string) *GetDeviceSettingRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceSettingRequestDeviceInfo) SetId(v string) *GetDeviceSettingRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetDeviceSettingRequestDeviceInfo) SetIdType(v string) *GetDeviceSettingRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetDeviceSettingRequestDeviceInfo) SetOrganizationId(v string) *GetDeviceSettingRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetDeviceSettingRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}
