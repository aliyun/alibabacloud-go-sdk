// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForTransferProhibitionLockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v []*string) *SaveBatchTaskForTransferProhibitionLockRequest
	GetDomainName() []*string
	SetLang(v string) *SaveBatchTaskForTransferProhibitionLockRequest
	GetLang() *string
	SetStatus(v bool) *SaveBatchTaskForTransferProhibitionLockRequest
	GetStatus() *bool
	SetUserClientIp(v string) *SaveBatchTaskForTransferProhibitionLockRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForTransferProhibitionLockRequest struct {
	// The domain names for which you want to enable or disable the transfer prohibition lock.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// The language of the error message that is returned if the request fails. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to enable or disable the transfer prohibition lock. Valid values:
	//
	// - **true**: Enable the transfer prohibition lock.
	//
	// - **false**: Disable the transfer prohibition lock.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// The client IP address.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForTransferProhibitionLockRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForTransferProhibitionLockRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) GetStatus() *bool {
	return s.Status
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) SetDomainName(v []*string) *SaveBatchTaskForTransferProhibitionLockRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) SetLang(v string) *SaveBatchTaskForTransferProhibitionLockRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) SetStatus(v bool) *SaveBatchTaskForTransferProhibitionLockRequest {
	s.Status = &v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) SetUserClientIp(v string) *SaveBatchTaskForTransferProhibitionLockRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) Validate() error {
	return dara.Validate(s)
}
