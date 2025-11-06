// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactType(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest
	GetContactType() *string
	SetDomainName(v []*string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest
	GetDomainName() []*string
	SetLang(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest
	GetLang() *string
	SetRegistrantProfileId(v int64) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest
	GetRegistrantProfileId() *int64
	SetTransferOutProhibited(v bool) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest
	GetTransferOutProhibited() *bool
	SetUserClientIp(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest struct {
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
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
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
	// true
	TransferOutProhibited *bool `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) GetContactType() *string {
	return s.ContactType
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) GetTransferOutProhibited() *bool {
	return s.TransferOutProhibited
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetContactType(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.ContactType = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetDomainName(v []*string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetLang(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetRegistrantProfileId(v int64) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetTransferOutProhibited(v bool) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.TransferOutProhibited = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetUserClientIp(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) Validate() error {
	return dara.Validate(s)
}
