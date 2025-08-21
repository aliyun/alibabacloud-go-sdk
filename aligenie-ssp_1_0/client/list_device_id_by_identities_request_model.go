// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceIdByIdentitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncodeKey(v string) *ListDeviceIdByIdentitiesRequest
	GetEncodeKey() *string
	SetEncodeType(v string) *ListDeviceIdByIdentitiesRequest
	GetEncodeType() *string
	SetIdentityIds(v []*string) *ListDeviceIdByIdentitiesRequest
	GetIdentityIds() []*string
	SetIdentityType(v string) *ListDeviceIdByIdentitiesRequest
	GetIdentityType() *string
	SetProductKey(v string) *ListDeviceIdByIdentitiesRequest
	GetProductKey() *string
}

type ListDeviceIdByIdentitiesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 125****0946
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType  *string   `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	IdentityIds []*string `json:"IdentityIds,omitempty" xml:"IdentityIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// MAC
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Mm*****XnZ8
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s ListDeviceIdByIdentitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceIdByIdentitiesRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceIdByIdentitiesRequest) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListDeviceIdByIdentitiesRequest) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListDeviceIdByIdentitiesRequest) GetIdentityIds() []*string {
	return s.IdentityIds
}

func (s *ListDeviceIdByIdentitiesRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *ListDeviceIdByIdentitiesRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *ListDeviceIdByIdentitiesRequest) SetEncodeKey(v string) *ListDeviceIdByIdentitiesRequest {
	s.EncodeKey = &v
	return s
}

func (s *ListDeviceIdByIdentitiesRequest) SetEncodeType(v string) *ListDeviceIdByIdentitiesRequest {
	s.EncodeType = &v
	return s
}

func (s *ListDeviceIdByIdentitiesRequest) SetIdentityIds(v []*string) *ListDeviceIdByIdentitiesRequest {
	s.IdentityIds = v
	return s
}

func (s *ListDeviceIdByIdentitiesRequest) SetIdentityType(v string) *ListDeviceIdByIdentitiesRequest {
	s.IdentityType = &v
	return s
}

func (s *ListDeviceIdByIdentitiesRequest) SetProductKey(v string) *ListDeviceIdByIdentitiesRequest {
	s.ProductKey = &v
	return s
}

func (s *ListDeviceIdByIdentitiesRequest) Validate() error {
	return dara.Validate(s)
}
