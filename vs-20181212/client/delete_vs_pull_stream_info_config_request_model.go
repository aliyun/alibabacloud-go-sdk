// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVsPullStreamInfoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteVsPullStreamInfoConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteVsPullStreamInfoConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteVsPullStreamInfoConfigRequest
	GetOwnerId() *int64
	SetStreamName(v string) *DeleteVsPullStreamInfoConfigRequest
	GetStreamName() *string
}

type DeleteVsPullStreamInfoConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteVsPullStreamInfoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVsPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteVsPullStreamInfoConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteVsPullStreamInfoConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteVsPullStreamInfoConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVsPullStreamInfoConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteVsPullStreamInfoConfigRequest) SetAppName(v string) *DeleteVsPullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteVsPullStreamInfoConfigRequest) SetDomainName(v string) *DeleteVsPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteVsPullStreamInfoConfigRequest) SetOwnerId(v int64) *DeleteVsPullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVsPullStreamInfoConfigRequest) SetStreamName(v string) *DeleteVsPullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteVsPullStreamInfoConfigRequest) Validate() error {
	return dara.Validate(s)
}
