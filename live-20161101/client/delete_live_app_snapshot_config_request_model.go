// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAppSnapshotConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLiveAppSnapshotConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLiveAppSnapshotConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveAppSnapshotConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteLiveAppSnapshotConfigRequest
	GetSecurityToken() *string
}

type DeleteLiveAppSnapshotConfigRequest struct {
	// The name of the application to which the live stream belongs. The value of this parameter must be the same as the application name in the ingest URL. Otherwise, the configuration does not take effect. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
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
}

func (s DeleteLiveAppSnapshotConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAppSnapshotConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAppSnapshotConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLiveAppSnapshotConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveAppSnapshotConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveAppSnapshotConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLiveAppSnapshotConfigRequest) SetAppName(v string) *DeleteLiveAppSnapshotConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveAppSnapshotConfigRequest) SetDomainName(v string) *DeleteLiveAppSnapshotConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveAppSnapshotConfigRequest) SetOwnerId(v int64) *DeleteLiveAppSnapshotConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveAppSnapshotConfigRequest) SetSecurityToken(v string) *DeleteLiveAppSnapshotConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveAppSnapshotConfigRequest) Validate() error {
	return dara.Validate(s)
}
