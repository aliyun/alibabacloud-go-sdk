// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfos(v *ListDeviceBasicInfoRequestDeviceInfos) *ListDeviceBasicInfoRequest
	GetDeviceInfos() *ListDeviceBasicInfoRequestDeviceInfos
}

type ListDeviceBasicInfoRequest struct {
	// List of device identity information.
	DeviceInfos *ListDeviceBasicInfoRequestDeviceInfos `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty" type:"Struct"`
}

func (s ListDeviceBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoRequest) GetDeviceInfos() *ListDeviceBasicInfoRequestDeviceInfos {
	return s.DeviceInfos
}

func (s *ListDeviceBasicInfoRequest) SetDeviceInfos(v *ListDeviceBasicInfoRequestDeviceInfos) *ListDeviceBasicInfoRequest {
	s.DeviceInfos = v
	return s
}

func (s *ListDeviceBasicInfoRequest) Validate() error {
	if s.DeviceInfos != nil {
		if err := s.DeviceInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeviceBasicInfoRequestDeviceInfos struct {
	// Value corresponding to the encoding type. Enter the Project ID of the project where the product resides. You can view this in the Tmall Genie AI Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. Enter **PROJECT_ID*	- here.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Type of device ID:
	//
	// - OPEN_ID: Default device ID identity.
	//
	// - UNION_ID: Organization-dimension device ID identity. You must request an organization in advance on the Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// List of device identity information.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// Organization ID of the device. Required if IdType is UNION_ID.
	//
	// example:
	//
	// 1
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListDeviceBasicInfoRequestDeviceInfos) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBasicInfoRequestDeviceInfos) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) GetIdType() *string {
	return s.IdType
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) GetIds() []*string {
	return s.Ids
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) SetEncodeKey(v string) *ListDeviceBasicInfoRequestDeviceInfos {
	s.EncodeKey = &v
	return s
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) SetEncodeType(v string) *ListDeviceBasicInfoRequestDeviceInfos {
	s.EncodeType = &v
	return s
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) SetIdType(v string) *ListDeviceBasicInfoRequestDeviceInfos {
	s.IdType = &v
	return s
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) SetIds(v []*string) *ListDeviceBasicInfoRequestDeviceInfos {
	s.Ids = v
	return s
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) SetOrganizationId(v string) *ListDeviceBasicInfoRequestDeviceInfos {
	s.OrganizationId = &v
	return s
}

func (s *ListDeviceBasicInfoRequestDeviceInfos) Validate() error {
	return dara.Validate(s)
}
