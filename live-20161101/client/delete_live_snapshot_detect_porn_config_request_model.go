// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotDetectPornConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLiveSnapshotDetectPornConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLiveSnapshotDetectPornConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveSnapshotDetectPornConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteLiveSnapshotDetectPornConfigRequest
	GetSecurityToken() *string
}

type DeleteLiveSnapshotDetectPornConfigRequest struct {
	// The name of the application to which the live stream belongs. You can call the [DescribeLiveSnapshotDetectPornConfig](https://help.aliyun.com/document_detail/2847918.html) operation to query the application name.
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

func (s DeleteLiveSnapshotDetectPornConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotDetectPornConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) SetAppName(v string) *DeleteLiveSnapshotDetectPornConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) SetDomainName(v string) *DeleteLiveSnapshotDetectPornConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) SetOwnerId(v int64) *DeleteLiveSnapshotDetectPornConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) SetSecurityToken(v string) *DeleteLiveSnapshotDetectPornConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveSnapshotDetectPornConfigRequest) Validate() error {
	return dara.Validate(s)
}
