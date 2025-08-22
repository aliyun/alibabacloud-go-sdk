// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopDcdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchStopDcdnDomainRequest
	GetDomainNames() *string
	SetOwnerId(v int64) *BatchStopDcdnDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchStopDcdnDomainRequest
	GetSecurityToken() *string
}

type BatchStopDcdnDomainRequest struct {
	// The accelerated domain names. If you need to specify multiple accelerated domain names, separate domain names with commas (,).
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

func (s BatchStopDcdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStopDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStopDcdnDomainRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchStopDcdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchStopDcdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchStopDcdnDomainRequest) SetDomainNames(v string) *BatchStopDcdnDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStopDcdnDomainRequest) SetOwnerId(v int64) *BatchStopDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopDcdnDomainRequest) SetSecurityToken(v string) *BatchStopDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchStopDcdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
