// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDeletingDSRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForDeletingDSRecordRequest
	GetDomainName() *string
	SetKeyTag(v int32) *SaveSingleTaskForDeletingDSRecordRequest
	GetKeyTag() *int32
	SetLang(v string) *SaveSingleTaskForDeletingDSRecordRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForDeletingDSRecordRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForDeletingDSRecordRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	KeyTag       *int32  `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForDeletingDSRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDeletingDSRecordRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) GetKeyTag() *int32 {
	return s.KeyTag
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) SetDomainName(v string) *SaveSingleTaskForDeletingDSRecordRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) SetKeyTag(v int32) *SaveSingleTaskForDeletingDSRecordRequest {
	s.KeyTag = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) SetLang(v string) *SaveSingleTaskForDeletingDSRecordRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) SetUserClientIp(v string) *SaveSingleTaskForDeletingDSRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) Validate() error {
	return dara.Validate(s)
}
