// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v []*string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest
	GetDomainName() []*string
	SetLang(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest
	GetLang() *string
	SetRegistrantProfileId(v int64) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest
	GetRegistrantProfileId() *int64
	SetTransferOutProhibited(v bool) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest
	GetTransferOutProhibited() *bool
	SetUserClientIp(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest
	GetUserClientIp() *string
}

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest struct {
	// A list of domain names. If you specify multiple domain names, pass them as a **list**. Call the [QueryDomainList](https://help.aliyun.com/document_detail/69362.htm?spm=a2c4g.11186623.0.0.33f4253cSJy3m8) API to obtain a list of your domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// The language of error messages returned by the API. Valid values:
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
	// The registrant profile ID. Call the [QueryRegistrantProfiles](https://help.aliyun.com/document_detail/67701.htm?spm=a2c4g.11186623.0.0.33f420daTwRQaO) API to query the registrant profile ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// Specifies whether to enable a 60-day transfer lock on the domain name after its registrant information is updated. Valid values:
	//
	// - **false**: Do not apply the lock.
	//
	// - **true**: Apply the lock.
	//
	// Default value: **false**.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	TransferOutProhibited *bool `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	// The IP address of the user. You can set this parameter to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) GetTransferOutProhibited() *bool {
	return s.TransferOutProhibited
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) SetDomainName(v []*string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest {
	s.DomainName = v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) SetLang(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) SetRegistrantProfileId(v int64) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) SetTransferOutProhibited(v bool) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest {
	s.TransferOutProhibited = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) SetUserClientIp(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) Validate() error {
	return dara.Validate(s)
}
