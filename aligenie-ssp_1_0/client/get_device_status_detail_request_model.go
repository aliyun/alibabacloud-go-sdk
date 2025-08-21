// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceStatusDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetDeviceStatusDetailRequestDeviceInfo) *GetDeviceStatusDetailRequest
	GetDeviceInfo() *GetDeviceStatusDetailRequestDeviceInfo
	SetKeys(v []*string) *GetDeviceStatusDetailRequest
	GetKeys() []*string
}

type GetDeviceStatusDetailRequest struct {
	// This parameter is required.
	DeviceInfo *GetDeviceStatusDetailRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s GetDeviceStatusDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailRequest) GetDeviceInfo() *GetDeviceStatusDetailRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetDeviceStatusDetailRequest) GetKeys() []*string {
	return s.Keys
}

func (s *GetDeviceStatusDetailRequest) SetDeviceInfo(v *GetDeviceStatusDetailRequestDeviceInfo) *GetDeviceStatusDetailRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetDeviceStatusDetailRequest) SetKeys(v []*string) *GetDeviceStatusDetailRequest {
	s.Keys = v
	return s
}

func (s *GetDeviceStatusDetailRequest) Validate() error {
	return dara.Validate(s)
}

type GetDeviceStatusDetailRequestDeviceInfo struct {
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
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetDeviceStatusDetailRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusDetailRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) SetEncodeKey(v string) *GetDeviceStatusDetailRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) SetEncodeType(v string) *GetDeviceStatusDetailRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) SetId(v string) *GetDeviceStatusDetailRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) SetIdType(v string) *GetDeviceStatusDetailRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) SetOrganizationId(v string) *GetDeviceStatusDetailRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetDeviceStatusDetailRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}
