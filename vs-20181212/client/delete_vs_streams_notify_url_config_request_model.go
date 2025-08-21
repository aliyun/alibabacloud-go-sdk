// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVsStreamsNotifyUrlConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteVsStreamsNotifyUrlConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteVsStreamsNotifyUrlConfigRequest
	GetOwnerId() *int64
}

type DeleteVsStreamsNotifyUrlConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteVsStreamsNotifyUrlConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVsStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteVsStreamsNotifyUrlConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteVsStreamsNotifyUrlConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVsStreamsNotifyUrlConfigRequest) SetDomainName(v string) *DeleteVsStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteVsStreamsNotifyUrlConfigRequest) SetOwnerId(v int64) *DeleteVsStreamsNotifyUrlConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVsStreamsNotifyUrlConfigRequest) Validate() error {
	return dara.Validate(s)
}
