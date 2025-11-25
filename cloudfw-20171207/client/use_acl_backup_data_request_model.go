// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUseAclBackupDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackUpTime(v string) *UseAclBackupDataRequest
	GetBackUpTime() *string
	SetLang(v string) *UseAclBackupDataRequest
	GetLang() *string
	SetSourceIp(v string) *UseAclBackupDataRequest
	GetSourceIp() *string
}

type UseAclBackupDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1743683400
	BackUpTime *string `json:"BackUpTime,omitempty" xml:"BackUpTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 123.113.99.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s UseAclBackupDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UseAclBackupDataRequest) GoString() string {
	return s.String()
}

func (s *UseAclBackupDataRequest) GetBackUpTime() *string {
	return s.BackUpTime
}

func (s *UseAclBackupDataRequest) GetLang() *string {
	return s.Lang
}

func (s *UseAclBackupDataRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *UseAclBackupDataRequest) SetBackUpTime(v string) *UseAclBackupDataRequest {
	s.BackUpTime = &v
	return s
}

func (s *UseAclBackupDataRequest) SetLang(v string) *UseAclBackupDataRequest {
	s.Lang = &v
	return s
}

func (s *UseAclBackupDataRequest) SetSourceIp(v string) *UseAclBackupDataRequest {
	s.SourceIp = &v
	return s
}

func (s *UseAclBackupDataRequest) Validate() error {
	return dara.Validate(s)
}
