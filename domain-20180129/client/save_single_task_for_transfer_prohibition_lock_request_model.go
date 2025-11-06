// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForTransferProhibitionLockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForTransferProhibitionLockRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForTransferProhibitionLockRequest
	GetLang() *string
	SetStatus(v bool) *SaveSingleTaskForTransferProhibitionLockRequest
	GetStatus() *bool
	SetUserClientIp(v string) *SaveSingleTaskForTransferProhibitionLockRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForTransferProhibitionLockRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForTransferProhibitionLockRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForTransferProhibitionLockRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) GetStatus() *bool {
	return s.Status
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) SetDomainName(v string) *SaveSingleTaskForTransferProhibitionLockRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) SetLang(v string) *SaveSingleTaskForTransferProhibitionLockRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) SetStatus(v bool) *SaveSingleTaskForTransferProhibitionLockRequest {
	s.Status = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) SetUserClientIp(v string) *SaveSingleTaskForTransferProhibitionLockRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) Validate() error {
	return dara.Validate(s)
}
