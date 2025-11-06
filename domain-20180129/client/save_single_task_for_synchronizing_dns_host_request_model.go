// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForSynchronizingDnsHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SaveSingleTaskForSynchronizingDnsHostRequest
	GetInstanceId() *string
	SetLang(v string) *SaveSingleTaskForSynchronizingDnsHostRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForSynchronizingDnsHostRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForSynchronizingDnsHostRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ST2017120814571100001303
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForSynchronizingDnsHostRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDnsHostRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDnsHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveSingleTaskForSynchronizingDnsHostRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForSynchronizingDnsHostRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForSynchronizingDnsHostRequest) SetInstanceId(v string) *SaveSingleTaskForSynchronizingDnsHostRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostRequest) SetLang(v string) *SaveSingleTaskForSynchronizingDnsHostRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostRequest) SetUserClientIp(v string) *SaveSingleTaskForSynchronizingDnsHostRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostRequest) Validate() error {
	return dara.Validate(s)
}
