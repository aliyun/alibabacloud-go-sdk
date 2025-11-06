// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForUpdatingContactInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddTransferLock(v bool) *SaveSingleTaskForUpdatingContactInfoRequest
	GetAddTransferLock() *bool
	SetContactType(v string) *SaveSingleTaskForUpdatingContactInfoRequest
	GetContactType() *string
	SetDomainName(v string) *SaveSingleTaskForUpdatingContactInfoRequest
	GetDomainName() *string
	SetInstanceId(v string) *SaveSingleTaskForUpdatingContactInfoRequest
	GetInstanceId() *string
	SetLang(v string) *SaveSingleTaskForUpdatingContactInfoRequest
	GetLang() *string
	SetRegistrantProfileId(v int64) *SaveSingleTaskForUpdatingContactInfoRequest
	GetRegistrantProfileId() *int64
	SetUserClientIp(v string) *SaveSingleTaskForUpdatingContactInfoRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForUpdatingContactInfoRequest struct {
	// example:
	//
	// false
	AddTransferLock *bool `json:"AddTransferLock,omitempty" xml:"AddTransferLock,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// registrant
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// S123456789
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForUpdatingContactInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForUpdatingContactInfoRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) GetAddTransferLock() *bool {
	return s.AddTransferLock
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) GetContactType() *string {
	return s.ContactType
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetAddTransferLock(v bool) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.AddTransferLock = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetContactType(v string) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.ContactType = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetDomainName(v string) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetInstanceId(v string) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetLang(v string) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetRegistrantProfileId(v int64) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetUserClientIp(v string) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) Validate() error {
	return dara.Validate(s)
}
