// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteLiveRecordNotifyConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveRecordNotifyConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteLiveRecordNotifyConfigRequest
	GetSecurityToken() *string
}

type DeleteLiveRecordNotifyConfigRequest struct {
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

func (s DeleteLiveRecordNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveRecordNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveRecordNotifyConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLiveRecordNotifyConfigRequest) SetDomainName(v string) *DeleteLiveRecordNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveRecordNotifyConfigRequest) SetOwnerId(v int64) *DeleteLiveRecordNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveRecordNotifyConfigRequest) SetSecurityToken(v string) *DeleteLiveRecordNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveRecordNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
