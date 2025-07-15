// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveRecordNotifyConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveRecordNotifyConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeLiveRecordNotifyConfigRequest
	GetSecurityToken() *string
}

type DescribeLiveRecordNotifyConfigRequest struct {
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

func (s DescribeLiveRecordNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveRecordNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveRecordNotifyConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveRecordNotifyConfigRequest) SetDomainName(v string) *DescribeLiveRecordNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigRequest) SetOwnerId(v int64) *DescribeLiveRecordNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigRequest) SetSecurityToken(v string) *DescribeLiveRecordNotifyConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
