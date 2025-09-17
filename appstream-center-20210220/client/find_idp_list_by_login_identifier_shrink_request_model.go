// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindIdpListByLoginIdentifierShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableFeaturesShrink(v string) *FindIdpListByLoginIdentifierShrinkRequest
	GetAvailableFeaturesShrink() *string
	SetClientChannel(v string) *FindIdpListByLoginIdentifierShrinkRequest
	GetClientChannel() *string
	SetClientId(v string) *FindIdpListByLoginIdentifierShrinkRequest
	GetClientId() *string
	SetClientIp(v string) *FindIdpListByLoginIdentifierShrinkRequest
	GetClientIp() *string
	SetClientOS(v string) *FindIdpListByLoginIdentifierShrinkRequest
	GetClientOS() *string
	SetClientVersion(v string) *FindIdpListByLoginIdentifierShrinkRequest
	GetClientVersion() *string
	SetLoginIdentifier(v string) *FindIdpListByLoginIdentifierShrinkRequest
	GetLoginIdentifier() *string
	SetSupportTypes(v []*string) *FindIdpListByLoginIdentifierShrinkRequest
	GetSupportTypes() []*string
	SetUuid(v string) *FindIdpListByLoginIdentifierShrinkRequest
	GetUuid() *string
}

type FindIdpListByLoginIdentifierShrinkRequest struct {
	AvailableFeaturesShrink *string `json:"AvailableFeatures,omitempty" xml:"AvailableFeatures,omitempty"`
	// example:
	//
	// pc
	ClientChannel *string `json:"ClientChannel,omitempty" xml:"ClientChannel,omitempty"`
	// example:
	//
	// 370b56f8-2812-4b6c-bfa6-2560791c****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Enterprise\\" 10.0 (Build 14393)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 2.0.1-D-20211008.101607
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Aliyun123***
	LoginIdentifier *string   `json:"LoginIdentifier,omitempty" xml:"LoginIdentifier,omitempty"`
	SupportTypes    []*string `json:"SupportTypes,omitempty" xml:"SupportTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 2943802884B27030B6759F9132B2****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s FindIdpListByLoginIdentifierShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FindIdpListByLoginIdentifierShrinkRequest) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) GetAvailableFeaturesShrink() *string {
	return s.AvailableFeaturesShrink
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) GetClientChannel() *string {
	return s.ClientChannel
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) GetClientId() *string {
	return s.ClientId
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) GetLoginIdentifier() *string {
	return s.LoginIdentifier
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) GetSupportTypes() []*string {
	return s.SupportTypes
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetAvailableFeaturesShrink(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.AvailableFeaturesShrink = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetClientChannel(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.ClientChannel = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetClientId(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetClientIp(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.ClientIp = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetClientOS(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.ClientOS = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetClientVersion(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetLoginIdentifier(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.LoginIdentifier = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetSupportTypes(v []*string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.SupportTypes = v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetUuid(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.Uuid = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) Validate() error {
	return dara.Validate(s)
}
