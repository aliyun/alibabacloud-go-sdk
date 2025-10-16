// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclBackupDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackUpTime(v string) *DeleteAclBackupDataRequest
	GetBackUpTime() *string
	SetLang(v string) *DeleteAclBackupDataRequest
	GetLang() *string
	SetSourceIp(v string) *DeleteAclBackupDataRequest
	GetSourceIp() *string
}

type DeleteAclBackupDataRequest struct {
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
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteAclBackupDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclBackupDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclBackupDataRequest) GetBackUpTime() *string {
	return s.BackUpTime
}

func (s *DeleteAclBackupDataRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteAclBackupDataRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteAclBackupDataRequest) SetBackUpTime(v string) *DeleteAclBackupDataRequest {
	s.BackUpTime = &v
	return s
}

func (s *DeleteAclBackupDataRequest) SetLang(v string) *DeleteAclBackupDataRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAclBackupDataRequest) SetSourceIp(v string) *DeleteAclBackupDataRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAclBackupDataRequest) Validate() error {
	return dara.Validate(s)
}
