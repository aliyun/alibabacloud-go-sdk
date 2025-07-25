// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *TransferDomainRequest
	GetDomainNames() *string
	SetLang(v string) *TransferDomainRequest
	GetLang() *string
	SetRemark(v string) *TransferDomainRequest
	GetRemark() *string
	SetTargetUserId(v int64) *TransferDomainRequest
	GetTargetUserId() *int64
}

type TransferDomainRequest struct {
	// The domain names. Separate multiple domain names with commas (,). Only domain names registered with Alibaba Cloud are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1.com,test2.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The description of the domain name.
	//
	// example:
	//
	// test domain transfer
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The destination user ID. The domain names and their Domain Name System (DNS) records are transferred to the destination user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	TargetUserId *int64 `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s TransferDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferDomainRequest) GoString() string {
	return s.String()
}

func (s *TransferDomainRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *TransferDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *TransferDomainRequest) GetRemark() *string {
	return s.Remark
}

func (s *TransferDomainRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *TransferDomainRequest) SetDomainNames(v string) *TransferDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *TransferDomainRequest) SetLang(v string) *TransferDomainRequest {
	s.Lang = &v
	return s
}

func (s *TransferDomainRequest) SetRemark(v string) *TransferDomainRequest {
	s.Remark = &v
	return s
}

func (s *TransferDomainRequest) SetTargetUserId(v int64) *TransferDomainRequest {
	s.TargetUserId = &v
	return s
}

func (s *TransferDomainRequest) Validate() error {
	return dara.Validate(s)
}
