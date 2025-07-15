// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDetectNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDetectNotifyConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveDetectNotifyConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeLiveDetectNotifyConfigRequest
	GetSecurityToken() *string
}

type DescribeLiveDetectNotifyConfigRequest struct {
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

func (s DescribeLiveDetectNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDetectNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDetectNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDetectNotifyConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveDetectNotifyConfigRequest) SetDomainName(v string) *DescribeLiveDetectNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDetectNotifyConfigRequest) SetOwnerId(v int64) *DescribeLiveDetectNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDetectNotifyConfigRequest) SetSecurityToken(v string) *DescribeLiveDetectNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveDetectNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
