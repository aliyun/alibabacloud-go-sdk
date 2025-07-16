// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopCdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchStopCdnDomainRequest
	GetDomainNames() *string
	SetOwnerId(v int64) *BatchStopCdnDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchStopCdnDomainRequest
	GetSecurityToken() *string
}

type BatchStopCdnDomainRequest struct {
	// The names of the accelerated domain names. You can specify one or more domain names in each request. Separate multiple domain names with commas (,).
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

func (s BatchStopCdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStopCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStopCdnDomainRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchStopCdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchStopCdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchStopCdnDomainRequest) SetDomainNames(v string) *BatchStopCdnDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStopCdnDomainRequest) SetOwnerId(v int64) *BatchStopCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopCdnDomainRequest) SetSecurityToken(v string) *BatchStopCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchStopCdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
