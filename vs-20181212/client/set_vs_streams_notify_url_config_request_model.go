// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVsStreamsNotifyUrlConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *SetVsStreamsNotifyUrlConfigRequest
	GetAuthKey() *string
	SetAuthType(v string) *SetVsStreamsNotifyUrlConfigRequest
	GetAuthType() *string
	SetDomainName(v string) *SetVsStreamsNotifyUrlConfigRequest
	GetDomainName() *string
	SetNotifyUrl(v string) *SetVsStreamsNotifyUrlConfigRequest
	GetNotifyUrl() *string
	SetOwnerId(v int64) *SetVsStreamsNotifyUrlConfigRequest
	GetOwnerId() *int64
}

type SetVsStreamsNotifyUrlConfigRequest struct {
	// example:
	//
	// ocs*****ace
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// example:
	//
	// type_a
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxx.com/xx
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SetVsStreamsNotifyUrlConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetVsStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *SetVsStreamsNotifyUrlConfigRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *SetVsStreamsNotifyUrlConfigRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *SetVsStreamsNotifyUrlConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetVsStreamsNotifyUrlConfigRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SetVsStreamsNotifyUrlConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetVsStreamsNotifyUrlConfigRequest) SetAuthKey(v string) *SetVsStreamsNotifyUrlConfigRequest {
	s.AuthKey = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigRequest) SetAuthType(v string) *SetVsStreamsNotifyUrlConfigRequest {
	s.AuthType = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigRequest) SetDomainName(v string) *SetVsStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigRequest) SetNotifyUrl(v string) *SetVsStreamsNotifyUrlConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigRequest) SetOwnerId(v int64) *SetVsStreamsNotifyUrlConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigRequest) Validate() error {
	return dara.Validate(s)
}
