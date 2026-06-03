// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdatingContactInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddTransferLock(v bool) *SaveBatchTaskForUpdatingContactInfoRequest
	GetAddTransferLock() *bool
	SetContactType(v string) *SaveBatchTaskForUpdatingContactInfoRequest
	GetContactType() *string
	SetDomainName(v []*string) *SaveBatchTaskForUpdatingContactInfoRequest
	GetDomainName() []*string
	SetLang(v string) *SaveBatchTaskForUpdatingContactInfoRequest
	GetLang() *string
	SetRegistrantProfileId(v int64) *SaveBatchTaskForUpdatingContactInfoRequest
	GetRegistrantProfileId() *int64
	SetUserClientIp(v string) *SaveBatchTaskForUpdatingContactInfoRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForUpdatingContactInfoRequest struct {
	AddTransferLock *bool `json:"AddTransferLock,omitempty" xml:"AddTransferLock,omitempty"`
	// This parameter is required.
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// This parameter is required.
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Lang       *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	RegistrantProfileId *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	UserClientIp        *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) GetAddTransferLock() *bool {
	return s.AddTransferLock
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) GetContactType() *string {
	return s.ContactType
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) SetAddTransferLock(v bool) *SaveBatchTaskForUpdatingContactInfoRequest {
	s.AddTransferLock = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) SetContactType(v string) *SaveBatchTaskForUpdatingContactInfoRequest {
	s.ContactType = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) SetDomainName(v []*string) *SaveBatchTaskForUpdatingContactInfoRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) SetLang(v string) *SaveBatchTaskForUpdatingContactInfoRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) SetRegistrantProfileId(v int64) *SaveBatchTaskForUpdatingContactInfoRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) SetUserClientIp(v string) *SaveBatchTaskForUpdatingContactInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoRequest) Validate() error {
	return dara.Validate(s)
}
