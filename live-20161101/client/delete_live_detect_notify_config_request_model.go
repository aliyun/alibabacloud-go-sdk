// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDetectNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteLiveDetectNotifyConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveDetectNotifyConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteLiveDetectNotifyConfigRequest
	GetSecurityToken() *string
}

type DeleteLiveDetectNotifyConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteLiveDetectNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDetectNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveDetectNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveDetectNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveDetectNotifyConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLiveDetectNotifyConfigRequest) SetDomainName(v string) *DeleteLiveDetectNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveDetectNotifyConfigRequest) SetOwnerId(v int64) *DeleteLiveDetectNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveDetectNotifyConfigRequest) SetSecurityToken(v string) *DeleteLiveDetectNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveDetectNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
