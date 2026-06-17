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
	// The time of the backup, specified as a UNIX timestamp (the number of seconds that have elapsed since 00:00:00 UTC on January 1, 1970).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1743683400
	BackUpTime *string `json:"BackUpTime,omitempty" xml:"BackUpTime,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address.
	//
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
