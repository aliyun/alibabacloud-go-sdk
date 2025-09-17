// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindIdpListByLoginIdentifierRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableFeatures(v map[string]*string) *FindIdpListByLoginIdentifierRequest
	GetAvailableFeatures() map[string]*string
	SetClientChannel(v string) *FindIdpListByLoginIdentifierRequest
	GetClientChannel() *string
	SetClientId(v string) *FindIdpListByLoginIdentifierRequest
	GetClientId() *string
	SetClientIp(v string) *FindIdpListByLoginIdentifierRequest
	GetClientIp() *string
	SetClientOS(v string) *FindIdpListByLoginIdentifierRequest
	GetClientOS() *string
	SetClientVersion(v string) *FindIdpListByLoginIdentifierRequest
	GetClientVersion() *string
	SetLoginIdentifier(v string) *FindIdpListByLoginIdentifierRequest
	GetLoginIdentifier() *string
	SetSupportTypes(v []*string) *FindIdpListByLoginIdentifierRequest
	GetSupportTypes() []*string
	SetUuid(v string) *FindIdpListByLoginIdentifierRequest
	GetUuid() *string
}

type FindIdpListByLoginIdentifierRequest struct {
	AvailableFeatures map[string]*string `json:"AvailableFeatures,omitempty" xml:"AvailableFeatures,omitempty"`
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

func (s FindIdpListByLoginIdentifierRequest) String() string {
	return dara.Prettify(s)
}

func (s FindIdpListByLoginIdentifierRequest) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierRequest) GetAvailableFeatures() map[string]*string {
	return s.AvailableFeatures
}

func (s *FindIdpListByLoginIdentifierRequest) GetClientChannel() *string {
	return s.ClientChannel
}

func (s *FindIdpListByLoginIdentifierRequest) GetClientId() *string {
	return s.ClientId
}

func (s *FindIdpListByLoginIdentifierRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *FindIdpListByLoginIdentifierRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *FindIdpListByLoginIdentifierRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *FindIdpListByLoginIdentifierRequest) GetLoginIdentifier() *string {
	return s.LoginIdentifier
}

func (s *FindIdpListByLoginIdentifierRequest) GetSupportTypes() []*string {
	return s.SupportTypes
}

func (s *FindIdpListByLoginIdentifierRequest) GetUuid() *string {
	return s.Uuid
}

func (s *FindIdpListByLoginIdentifierRequest) SetAvailableFeatures(v map[string]*string) *FindIdpListByLoginIdentifierRequest {
	s.AvailableFeatures = v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetClientChannel(v string) *FindIdpListByLoginIdentifierRequest {
	s.ClientChannel = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetClientId(v string) *FindIdpListByLoginIdentifierRequest {
	s.ClientId = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetClientIp(v string) *FindIdpListByLoginIdentifierRequest {
	s.ClientIp = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetClientOS(v string) *FindIdpListByLoginIdentifierRequest {
	s.ClientOS = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetClientVersion(v string) *FindIdpListByLoginIdentifierRequest {
	s.ClientVersion = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetLoginIdentifier(v string) *FindIdpListByLoginIdentifierRequest {
	s.LoginIdentifier = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetSupportTypes(v []*string) *FindIdpListByLoginIdentifierRequest {
	s.SupportTypes = v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetUuid(v string) *FindIdpListByLoginIdentifierRequest {
	s.Uuid = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) Validate() error {
	return dara.Validate(s)
}
