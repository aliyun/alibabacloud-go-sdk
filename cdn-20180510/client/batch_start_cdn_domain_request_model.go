// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartCdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchStartCdnDomainRequest
	GetDomainNames() *string
	SetOwnerId(v int64) *BatchStartCdnDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchStartCdnDomainRequest
	GetSecurityToken() *string
}

type BatchStartCdnDomainRequest struct {
	// The accelerated domain names. You can specify one or more domain names. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchStartCdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStartCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStartCdnDomainRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchStartCdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchStartCdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchStartCdnDomainRequest) SetDomainNames(v string) *BatchStartCdnDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStartCdnDomainRequest) SetOwnerId(v int64) *BatchStartCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStartCdnDomainRequest) SetSecurityToken(v string) *BatchStartCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchStartCdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
