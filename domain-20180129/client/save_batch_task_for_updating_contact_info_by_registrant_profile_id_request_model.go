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
	// The contact type to modify. Valid values:
	//
	// - **registrant**: The domain name\\"s registrant.
	//
	// - **admin**: The administrative contact for the domain name.
	//
	// - **billing**: The billing contact.
	//
	// - **tech**: The technical contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// registrant
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// An array of domain names to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// The language of the error message that is returned if the request fails. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the registrant profile. This ID is automatically generated when you create a registrant profile. You can find registrant profile IDs by calling the [QueryRegistrantProfiles](https://help.aliyun.com/document_detail/67701.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// Specifies whether to enable the transfer lock. This parameter is valid only when **ContactType*	- is set to **registrant**. If enabled, this feature prevents the domain name from being transferred for 60 days after the registrant information is modified.
	//
	// - **true**: Enables the lock, which prevents the domain name from being transferred out.
	//
	// - **false**: Disables the lock, which allows the domain name to be transferred out.
	//
	// Default value: **false**.
	//
	// example:
	//
	// true
	TransferOutProhibited *bool `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	// The IP address of the client. You can set this parameter to **127.0.0.1**.
	//
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
