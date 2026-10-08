// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForUpdateProhibitionLockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForUpdateProhibitionLockRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForUpdateProhibitionLockRequest
	GetLang() *string
	SetStatus(v bool) *SaveSingleTaskForUpdateProhibitionLockRequest
	GetStatus() *bool
	SetUserClientIp(v string) *SaveSingleTaskForUpdateProhibitionLockRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForUpdateProhibitionLockRequest struct {
	// Domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Language of the error message returned by the API. Valid values:
	//
	// - **zh**: Chinese;
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Enabled or shutdown status. Valid values:
	//
	// - **true**: Enabled;
	//
	// - **false**: Shutdown.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// User IP address.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForUpdateProhibitionLockRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForUpdateProhibitionLockRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) GetStatus() *bool {
	return s.Status
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) SetDomainName(v string) *SaveSingleTaskForUpdateProhibitionLockRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) SetLang(v string) *SaveSingleTaskForUpdateProhibitionLockRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) SetStatus(v bool) *SaveSingleTaskForUpdateProhibitionLockRequest {
	s.Status = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) SetUserClientIp(v string) *SaveSingleTaskForUpdateProhibitionLockRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) Validate() error {
	return dara.Validate(s)
}
