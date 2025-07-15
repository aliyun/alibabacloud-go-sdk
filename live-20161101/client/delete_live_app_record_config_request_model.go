// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAppRecordConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLiveAppRecordConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLiveAppRecordConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveAppRecordConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteLiveAppRecordConfigRequest
	GetSecurityToken() *string
	SetStreamName(v string) *DeleteLiveAppRecordConfigRequest
	GetStreamName() *string
}

type DeleteLiveAppRecordConfigRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
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
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLiveAppRecordConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAppRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAppRecordConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLiveAppRecordConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveAppRecordConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveAppRecordConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLiveAppRecordConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteLiveAppRecordConfigRequest) SetAppName(v string) *DeleteLiveAppRecordConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveAppRecordConfigRequest) SetDomainName(v string) *DeleteLiveAppRecordConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveAppRecordConfigRequest) SetOwnerId(v int64) *DeleteLiveAppRecordConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveAppRecordConfigRequest) SetSecurityToken(v string) *DeleteLiveAppRecordConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveAppRecordConfigRequest) SetStreamName(v string) *DeleteLiveAppRecordConfigRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteLiveAppRecordConfigRequest) Validate() error {
	return dara.Validate(s)
}
