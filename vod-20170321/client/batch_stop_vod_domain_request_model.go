// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopVodDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchStopVodDomainRequest
	GetDomainNames() *string
	SetOwnerId(v int64) *BatchStopVodDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchStopVodDomainRequest
	GetSecurityToken() *string
}

type BatchStopVodDomainRequest struct {
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

func (s BatchStopVodDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStopVodDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStopVodDomainRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchStopVodDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchStopVodDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchStopVodDomainRequest) SetDomainNames(v string) *BatchStopVodDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStopVodDomainRequest) SetOwnerId(v int64) *BatchStopVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopVodDomainRequest) SetSecurityToken(v string) *BatchStopVodDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchStopVodDomainRequest) Validate() error {
	return dara.Validate(s)
}
