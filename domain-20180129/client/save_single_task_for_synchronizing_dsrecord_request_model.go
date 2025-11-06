// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForSynchronizingDSRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForSynchronizingDSRecordRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForSynchronizingDSRecordRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForSynchronizingDSRecordRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForSynchronizingDSRecordRequest struct {
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
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForSynchronizingDSRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDSRecordRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDSRecordRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForSynchronizingDSRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForSynchronizingDSRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForSynchronizingDSRecordRequest) SetDomainName(v string) *SaveSingleTaskForSynchronizingDSRecordRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordRequest) SetLang(v string) *SaveSingleTaskForSynchronizingDSRecordRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordRequest) SetUserClientIp(v string) *SaveSingleTaskForSynchronizingDSRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordRequest) Validate() error {
	return dara.Validate(s)
}
