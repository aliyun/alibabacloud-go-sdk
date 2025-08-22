// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartDcdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchStartDcdnDomainRequest
	GetDomainNames() *string
	SetOwnerId(v int64) *BatchStartDcdnDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchStartDcdnDomainRequest
	GetSecurityToken() *string
}

type BatchStartDcdnDomainRequest struct {
	// The accelerated domain name. You can specify multiple accelerated domain names and separate them with commas (,).
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

func (s BatchStartDcdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStartDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStartDcdnDomainRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchStartDcdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchStartDcdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchStartDcdnDomainRequest) SetDomainNames(v string) *BatchStartDcdnDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStartDcdnDomainRequest) SetOwnerId(v int64) *BatchStartDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStartDcdnDomainRequest) SetSecurityToken(v string) *BatchStartDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchStartDcdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
