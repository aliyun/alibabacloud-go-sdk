// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartVodDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchStartVodDomainRequest
	GetDomainNames() *string
	SetOwnerId(v int64) *BatchStartVodDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchStartVodDomainRequest
	GetSecurityToken() *string
}

type BatchStartVodDomainRequest struct {
	// The accelerated domain name. Separate multiple domain names with commas (,).
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

func (s BatchStartVodDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStartVodDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStartVodDomainRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchStartVodDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchStartVodDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchStartVodDomainRequest) SetDomainNames(v string) *BatchStartVodDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStartVodDomainRequest) SetOwnerId(v int64) *BatchStartVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStartVodDomainRequest) SetSecurityToken(v string) *BatchStartVodDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchStartVodDomainRequest) Validate() error {
	return dara.Validate(s)
}
