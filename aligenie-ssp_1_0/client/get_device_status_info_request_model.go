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
	// List of device identity information.
	//
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
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeviceStatusInfoRequestDeviceInfo struct {
	// The value corresponding to the encoding type. Enter the Project ID of the project to which the product belongs. You can view it in the Tmall Genie AI Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. Set this parameter to **PROJECT_ID**.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device identifier. Specify the value of deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The Type of the Device ID:
	//
	// - OPEN_ID: The default device ID identity.
	//
	// - UNION_ID: The organization-dimension device ID identity. You must request an organization in the Open Platform in advance.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID of the device. This parameter is required if IdType is set to UNION_ID.
	//
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
