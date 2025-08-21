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
	return dara.Validate(s)
}

type ListDeviceBasicInfoRequestDeviceInfos struct {
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
	// OPEN_ID
	IdType *string   `json:"IdType,omitempty" xml:"IdType,omitempty"`
	Ids    []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
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
