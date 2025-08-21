// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceIdByIdentitiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncodeKey(v string) *ListDeviceIdByIdentitiesShrinkRequest
	GetEncodeKey() *string
	SetEncodeType(v string) *ListDeviceIdByIdentitiesShrinkRequest
	GetEncodeType() *string
	SetIdentityIdsShrink(v string) *ListDeviceIdByIdentitiesShrinkRequest
	GetIdentityIdsShrink() *string
	SetIdentityType(v string) *ListDeviceIdByIdentitiesShrinkRequest
	GetIdentityType() *string
	SetProductKey(v string) *ListDeviceIdByIdentitiesShrinkRequest
	GetProductKey() *string
}

type ListDeviceIdByIdentitiesShrinkRequest struct {
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
	EncodeType        *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	IdentityIdsShrink *string `json:"IdentityIds,omitempty" xml:"IdentityIds,omitempty"`
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

func (s ListDeviceIdByIdentitiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceIdByIdentitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) GetIdentityIdsShrink() *string {
	return s.IdentityIdsShrink
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) SetEncodeKey(v string) *ListDeviceIdByIdentitiesShrinkRequest {
	s.EncodeKey = &v
	return s
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) SetEncodeType(v string) *ListDeviceIdByIdentitiesShrinkRequest {
	s.EncodeType = &v
	return s
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) SetIdentityIdsShrink(v string) *ListDeviceIdByIdentitiesShrinkRequest {
	s.IdentityIdsShrink = &v
	return s
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) SetIdentityType(v string) *ListDeviceIdByIdentitiesShrinkRequest {
	s.IdentityType = &v
	return s
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) SetProductKey(v string) *ListDeviceIdByIdentitiesShrinkRequest {
	s.ProductKey = &v
	return s
}

func (s *ListDeviceIdByIdentitiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
