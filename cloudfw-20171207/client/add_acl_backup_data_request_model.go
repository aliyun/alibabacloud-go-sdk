// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAclBackupDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackUpTime(v string) *AddAclBackupDataRequest
	GetBackUpTime() *string
	SetDescription(v string) *AddAclBackupDataRequest
	GetDescription() *string
	SetLang(v string) *AddAclBackupDataRequest
	GetLang() *string
	SetSourceIp(v string) *AddAclBackupDataRequest
	GetSourceIp() *string
}

type AddAclBackupDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1743683400
	BackUpTime *string `json:"BackUpTime,omitempty" xml:"BackUpTime,omitempty"`
	// example:
	//
	// bj-001
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 115.194.124.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s AddAclBackupDataRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAclBackupDataRequest) GoString() string {
	return s.String()
}

func (s *AddAclBackupDataRequest) GetBackUpTime() *string {
	return s.BackUpTime
}

func (s *AddAclBackupDataRequest) GetDescription() *string {
	return s.Description
}

func (s *AddAclBackupDataRequest) GetLang() *string {
	return s.Lang
}

func (s *AddAclBackupDataRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *AddAclBackupDataRequest) SetBackUpTime(v string) *AddAclBackupDataRequest {
	s.BackUpTime = &v
	return s
}

func (s *AddAclBackupDataRequest) SetDescription(v string) *AddAclBackupDataRequest {
	s.Description = &v
	return s
}

func (s *AddAclBackupDataRequest) SetLang(v string) *AddAclBackupDataRequest {
	s.Lang = &v
	return s
}

func (s *AddAclBackupDataRequest) SetSourceIp(v string) *AddAclBackupDataRequest {
	s.SourceIp = &v
	return s
}

func (s *AddAclBackupDataRequest) Validate() error {
	return dara.Validate(s)
}
